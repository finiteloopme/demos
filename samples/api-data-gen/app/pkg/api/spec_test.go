package api

import (
	"math/rand"
	"testing"
)

func loadSpec(t *testing.T) *ApiSpec {
	uri := "https://raw.githubusercontent.com/OAI/OpenAPI-Specification/main/examples/v3.0/petstore-expanded.json"
	apiSpec := &ApiSpec{}
	err := apiSpec.LoadSpec(uri)
	// Check API Spec loads without an error
	if err != nil {
		t.Fatalf("Unxpected error: %s", err)
	}

	return apiSpec
}

func loadSpecDetails(t *testing.T) *ApiSpec {
	uri := "https://raw.githubusercontent.com/OAI/OpenAPI-Specification/main/examples/v3.0/petstore-expanded.json"
	apiSpec := &ApiSpec{}
	err := apiSpec.LoadSpecDetails(uri)
	// Check API Spec loads without an error
	if err != nil {
		t.Fatalf("Unxpected error: %s", err)
	}
	// Check Paths
	if len(apiSpec.Paths) < 1 {
		t.Fatalf("Number of paths should greater than 0.  Encountered: %v", len(apiSpec.Paths))
	}
	randomIdx := rand.Intn(len(apiSpec.Paths))
	if len(apiSpec.Paths[randomIdx].Operations) < 1 {
		t.Fatalf("Number of operations for %s should greater than 0.  Encountered: %v", apiSpec.Paths[randomIdx].Path, len(apiSpec.Paths[randomIdx].Operations))

	}

	return apiSpec
}

func TestGetSpec(t *testing.T) {
	// uri := "https://raw.githubusercontent.com/OAI/OpenAPI-Specification/main/examples/v3.0/petstore-expanded.json"
	_ = loadSpec(t)
}

func TestLoadSpecDetails(t *testing.T) {
	_ = loadSpecDetails(t)
}
