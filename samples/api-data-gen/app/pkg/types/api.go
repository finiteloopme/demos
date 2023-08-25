package types

import "github.com/getkin/kin-openapi/openapi3"

// Open API Spec
//
//	type Api struct {
//		Api   *openapi3.T `json:"api"`
//		Uri   string      `json:"uri"`
//		Paths []*Path     `json:"paths"`
//	}
type ApiType *openapi3.T

// Uri
type UriType string

func (uri *UriType) String() string {
	return string(*uri)
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
	Type        string      `json:"type"`
	OperationID string      `json:"operationId"`
	Description string      `json:"description"`
	Summary     string      `json:"summary"`
	Request     RequestType `json:"request"`
	Response    ResonseType `json:"response"`
}

type RequestType struct {
	Params  []*Param  `json:"params"`
	Headers []*Header `json:"headers"`
	Body    []byte    `json:"body"`
}

type ResonseType struct {
	Headers []*Header `json:"headers"`
	Body    []byte    `json:"body"`
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
