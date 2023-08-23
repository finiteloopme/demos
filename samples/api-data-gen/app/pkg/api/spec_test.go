package api

import (
	"testing"
)

var TestUri = "https://raw.githubusercontent.com/OAI/OpenAPI-Specification/main/examples/v3.0/petstore-expanded.json"

func TestGetSpec(t *testing.T) {
	// uri := "https://raw.githubusercontent.com/OAI/OpenAPI-Specification/main/examples/v3.0/petstore-expanded.json"
	apiSpec := &ApiSpec{}
	err := apiSpec.GetSpec(TestUri)
	if err != nil {
		t.Fatalf("Unxpected error: %s", err)
	}
}
