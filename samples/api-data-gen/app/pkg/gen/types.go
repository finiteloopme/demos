package gen

import (
	"github.com/finiteloopme/demos/samples/api-data-gen/app/pkg/types"
)

// Open API Spec
type ApiGen struct {
	Api   types.ApiType `json:"spec"`
	Uri   types.UriType `json:"uri"`
	Paths []*types.Path `json:"paths"`
}
