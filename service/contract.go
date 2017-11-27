//go:generate echo $PWD/$GOPACKAGE/$GOFILE
//go:generate mockgen -package service -destination $PWD/service/mock_contract.go github.com/kamilsk/form-api/service DataLoader
package service

import (
	"github.com/kamilsk/form-api/data"
	"github.com/kamilsk/form-api/data/form"
)

// DataLoader defines behavior of DAO.
type DataLoader interface {
	Schema(data.UUID) (form.Schema, error)
	AddData(data.UUID, map[string][]string) (int64, error)
}