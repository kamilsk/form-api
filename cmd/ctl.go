package cmd

import (
	"context"
	"io"
	"io/ioutil"
	"os"
	"reflect"
	"strings"
	"time"

	pb "github.com/kamilsk/form-api/pkg/server/grpc"

	"github.com/kamilsk/form-api/pkg/config"
	"github.com/kamilsk/go-kit/pkg/fn"
	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"gopkg.in/yaml.v2"
	)

type schema struct {
	Kind    string                 `yaml:"kind"`
	Payload map[string]interface{} `yaml:"payload"`
}

type factory map[*cobra.Command]map[string]func() interface{}

func (f factory) new(cmd *cobra.Command) (interface{}, error) {
	data, err := f.data(cmd.Flag("file").Value.String())
	if err != nil {
		return nil, err
	}
	builder, ok := f[cmd][data.Kind]
	if !ok {
		return nil, errors.Errorf("unknown payload type %q", data.Kind)
	}
	entity := builder()
	if err = mapstructure.Decode(data.Payload, &entity); err != nil {
		return nil, errors.Wrapf(err, "trying to decode payload to %#v", entity)
	}
	return entity, nil
}

func (f factory) data(file string) (schema, error) {
	var (
		err error
		out schema
		raw []byte
		src io.Reader = os.Stdin
	)
	if file != "" {
		if src, err = os.Open(file); err != nil {
			return out, errors.Wrapf(err, "trying to open file %q", file)
		}
	} else {
		file = "/dev/stdin"
	}
	if raw, err = ioutil.ReadAll(src); err != nil {
		return out, errors.Wrapf(err, "trying to read file %q", file)
	}
	err = yaml.Unmarshal(raw, &out)
	return out, errors.Wrapf(err, "trying to decode file %q as YAML", file)
}

var entities factory

func call(cnf config.GRPCConfig, entity interface{}) (interface{}, error) {
	conn, err := grpc.Dial(cnf.Interface, grpc.WithInsecure())
	if err != nil {
		return nil, errors.Wrapf(err, "trying to connect to the gRPC server %q", cnf.Interface)
	}
	defer conn.Close()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	switch request := entity.(type) {
	case *pb.CreateSchemaRequest:
		client := pb.NewSchemaClient(conn)
		return client.Create(ctx, request)
	case *pb.CreateTemplateRequest:
		client := pb.NewTemplateClient(conn)
		return client.Create(ctx, request)
	case *pb.ReadSchemaRequest:
		client := pb.NewSchemaClient(conn)
		return client.Read(ctx, request)
	case *pb.ReadTemplateRequest:
		client := pb.NewTemplateClient(conn)
		return client.Read(ctx, request)
	case *pb.UpdateSchemaRequest:
		client := pb.NewSchemaClient(conn)
		return client.Update(ctx, request)
	case *pb.UpdateTemplateRequest:
		client := pb.NewTemplateClient(conn)
		return client.Update(ctx, request)
	case *pb.DeleteSchemaRequest:
		client := pb.NewSchemaClient(conn)
		return client.Delete(ctx, request)
	case *pb.DeleteTemplateRequest:
		client := pb.NewTemplateClient(conn)
		return client.Delete(ctx, request)
	default:
		return nil, errors.Errorf("unknown type %T", request)
	}
}

func convert(entity interface{}) map[string]interface{} {
	t := reflect.ValueOf(entity).Elem().Type()
	m := make(map[string]interface{}, t.NumField())
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		v, ok := f.Tag.Lookup("json")
		if !ok {
			continue
		}
		p := strings.Split(v, ",")
		if p[0] == "-" {
			continue
		}
		switch f.Type.String() {
		case "[]uint8":
			m[p[0]] = "binary"
		default:
			m[p[0]] = f.Type.String()
		}
	}
	return m
}

var (
	controlCmd = &cobra.Command{
		Use:   "ctl",
		Short: "Communicate with Forma server via gRPC",
	}
	createCmd = &cobra.Command{
		Use:   "create",
		Short: "Create some kind",
		RunE: func(cmd *cobra.Command, args []string) error {
			entity, err := entities.new(cmd)
			if err != nil {
				return err
			}
			response, err := call(cnf.Union.GRPCConfig, entity)
			if err != nil {
				return err
			}
			cmd.Printf("`ctl %s` was called, in:%#v out:%#v\n", cmd.Use, entity, response)
			return nil
		},
	}
	readCmd = &cobra.Command{
		Use:   "read",
		Short: "Read some kind",
		RunE: func(cmd *cobra.Command, args []string) error {
			entity, err := entities.new(cmd)
			if err != nil {
				return err
			}
			response, err := call(cnf.Union.GRPCConfig, entity)
			if err != nil {
				return err
			}
			cmd.Printf("`ctl %s` was called, in:%#v out:%#v\n", cmd.Use, entity, response)
			return nil
		},
	}
	updateCmd = &cobra.Command{
		Use:   "update",
		Short: "Update some kind",
		RunE: func(cmd *cobra.Command, args []string) error {
			entity, err := entities.new(cmd)
			if err != nil {
				return err
			}
			response, err := call(cnf.Union.GRPCConfig, entity)
			if err != nil {
				return err
			}
			cmd.Printf("`ctl %s` was called, in:%#v out:%#v\n", cmd.Use, entity, response)
			return nil
		},
	}
	deleteCmd = &cobra.Command{
		Use:   "delete",
		Short: "Delete some kind",
		RunE: func(cmd *cobra.Command, args []string) error {
			entity, err := entities.new(cmd)
			if err != nil {
				return err
			}
			response, err := call(cnf.Union.GRPCConfig, entity)
			if err != nil {
				return err
			}
			cmd.Printf("`ctl %s` was called, in:%#v out:%#v\n", cmd.Use, entity, response)
			return nil
		},
	}
	schemaCmd = &cobra.Command{
		Use:   "schema",
		Short: "Print schema of another control command",
		RunE: func(cmd *cobra.Command, args []string) error {
			var (
				target   *cobra.Command
				builders map[string]func() interface{}
				found    bool
			)
			use := cmd.Flag("for").Value.String()
			for target, builders = range entities {
				if strings.EqualFold(target.Use, use) {
					found = true
					break
				}
			}
			if !found {
				return errors.Errorf("unknown control command %q", use)
			}
			for kind, builder := range builders {
				yaml.NewEncoder(cmd.OutOrStdout()).Encode(schema{
					Kind:    kind,
					Payload: convert(builder()),
				})
				cmd.Println()
			}
			return nil
		},
	}
)

func init() {
	v := viper.New()
	fn.Must(
		func() error { return v.BindEnv("forma_token") },
		func() error { return v.BindEnv("grpc_host") },
		func() error {
			v.SetDefault("forma_token", defaults["forma_token"])
			v.SetDefault("grpc_host", defaults["grpc_host"])
			return nil
		},
		func() error {
			file := ""
			flags := controlCmd.PersistentFlags()
			flags.StringVarP(&file, "file", "f", file, "entity source (default is stdin)")
			flags.StringVarP(&cnf.Union.GRPCConfig.Interface,
				"grpc-host", "", v.GetString("grpc_host"), "gRPC server host")
			flags.DurationVarP(&cnf.Union.GRPCConfig.Timeout,
				"timeout", "t", time.Second, "connection timeout")
			flags.StringVarP((*string)(&cnf.Union.GRPCConfig.Token),
				"token", "", v.GetString("forma_token"), "user access token")
			schemaCmd.Flags().String("for", "", "which command: create, read, update or delete")
			schemaCmd.MarkFlagRequired("for")
			return nil
		},
		func() error {
			entities = factory{
				createCmd: {
					"Schema":   func() interface{} { return &pb.CreateSchemaRequest{} },
					"Template": func() interface{} { return &pb.CreateTemplateRequest{} },
				},
				readCmd: {
					"Schema":   func() interface{} { return &pb.DeleteSchemaRequest{} },
					"Template": func() interface{} { return &pb.DeleteTemplateRequest{} },
				},
				updateCmd: {
					"Schema":   func() interface{} { return &pb.UpdateSchemaRequest{} },
					"Template": func() interface{} { return &pb.UpdateTemplateRequest{} },
				},
				deleteCmd: {
					"Schema":   func() interface{} { return &pb.DeleteSchemaRequest{} },
					"Template": func() interface{} { return &pb.DeleteTemplateRequest{} },
				},
			}
			return nil
		},
	)
	controlCmd.AddCommand(createCmd, readCmd, updateCmd, deleteCmd)
}
