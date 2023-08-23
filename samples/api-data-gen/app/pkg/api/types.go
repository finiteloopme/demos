package api

import "github.com/getkin/kin-openapi/openapi3"

type ApiSpec struct {
	Api   *openapi3.T `json:"api"`
	Uri   string      `json:"uri"`
	Paths []Path      `json:"paths"`
}

type Path struct {
	Path       string      `json:"path"`
	Operations []Operation `json:"operations"`
}

type OperationType int

const (
	GET OperationType = iota
	POST
	PUT
	DELETE
	PATCH
)

func (restOperation OperationType) String() string {
	switch restOperation {
	case GET:
		return "GET"
	case POST:
		return "POST"
	case PUT:
		return "PUT"
	case DELETE:
		return "DELETE"
	case PATCH:
		return "PATCH"
	}
	return "Unknown"
}

type Operation struct {
	Type        OperationType `json:"type"`
	OperationID string        `json:"operationId"`
}
