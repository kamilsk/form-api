// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package domain

import (
	json "encoding/json"

	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson34644b0eDecodeGoOctolabOrgEcosystemFormaInternalDomain(in *jlexer.Lexer, out *Input) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "id":
			out.ID = string(in.String())
		case "name":
			out.Name = string(in.String())
		case "type":
			out.Type = string(in.String())
		case "title":
			out.Title = string(in.String())
		case "placeholder":
			out.Placeholder = string(in.String())
		case "value":
			out.Value = string(in.String())
		case "minlength":
			out.MinLength = int(in.Int())
		case "maxlength":
			out.MaxLength = int(in.Int())
		case "required":
			out.Required = bool(in.Bool())
		case "strict":
			out.Strict = bool(in.Bool())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson34644b0eEncodeGoOctolabOrgEcosystemFormaInternalDomain(out *jwriter.Writer, in Input) {
	out.RawByte('{')
	first := true
	_ = first
	if in.ID != "" {
		const prefix string = ",\"id\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.ID))
	}
	{
		const prefix string = ",\"name\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"type\":"
		out.RawString(prefix)
		out.String(string(in.Type))
	}
	if in.Title != "" {
		const prefix string = ",\"title\":"
		out.RawString(prefix)
		out.String(string(in.Title))
	}
	if in.Placeholder != "" {
		const prefix string = ",\"placeholder\":"
		out.RawString(prefix)
		out.String(string(in.Placeholder))
	}
	if in.Value != "" {
		const prefix string = ",\"value\":"
		out.RawString(prefix)
		out.String(string(in.Value))
	}
	if in.MinLength != 0 {
		const prefix string = ",\"minlength\":"
		out.RawString(prefix)
		out.Int(int(in.MinLength))
	}
	if in.MaxLength != 0 {
		const prefix string = ",\"maxlength\":"
		out.RawString(prefix)
		out.Int(int(in.MaxLength))
	}
	if in.Required {
		const prefix string = ",\"required\":"
		out.RawString(prefix)
		out.Bool(bool(in.Required))
	}
	if in.Strict {
		const prefix string = ",\"strict\":"
		out.RawString(prefix)
		out.Bool(bool(in.Strict))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Input) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson34644b0eEncodeGoOctolabOrgEcosystemFormaInternalDomain(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Input) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson34644b0eEncodeGoOctolabOrgEcosystemFormaInternalDomain(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Input) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson34644b0eDecodeGoOctolabOrgEcosystemFormaInternalDomain(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Input) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson34644b0eDecodeGoOctolabOrgEcosystemFormaInternalDomain(l, v)
}
