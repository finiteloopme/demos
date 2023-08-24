package api

import "github.com/getkin/kin-openapi/openapi3"

// Open API Spec
type ApiSpec struct {
	Api   *openapi3.T `json:"api"`
	Uri   string      `json:"uri"`
	Paths []*Path     `json:"paths"`
}

// REST API Endpoint
type Path struct {
	Path       string       `json:"path"`
	Operations []*Operation `json:"operations"`
}

// REST Operation
type OperationType int

const (
	GET OperationType = iota
	POST
	PUT
	DELETE
	PATCH
)

// Helper function to convert OperationType into string
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

// REST method or operation
type Operation struct {
	// Note the type is string and not OperationType
	Type           string    `json:"type"`
	OperationID    string    `json:"operationId"`
	Description    string    `json:"description"`
	Summary        string    `json:"summary"`
	RequestParams  []*Param  `json:"requests"`
	RequestHeaders []*Header `json:"req-headers"`
}

type Header struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Param struct {
	Type  string `json:"type"`
	Key   string `json:"key"`
	Value string `json:"value"`
}
