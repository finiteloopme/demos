package gen

import (
	"log"
	"math/rand"
	"os"
	"testing"

	"github.com/finiteloopme/demos/samples/api-data-gen/app/pkg/api"
)

func loadGen(t *testing.T) *ApiGen {
	uri := "https://raw.githubusercontent.com/OAI/OpenAPI-Specification/main/examples/v3.0/petstore-expanded.json"
	apiSpec := &api.ApiSpec{}
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
		t.Fatalf("Number of operations for %s should be greater than 0.  Encountered: %v", apiSpec.Paths[randomIdx].Path, len(apiSpec.Paths[randomIdx].Operations))

	}
	defer func() {
		log.SetOutput(os.Stderr)
	}()

	apiGen := &ApiGen{}
	err = apiGen.LoadApiSpec(apiSpec)

	if err != nil {
		t.Fatalf("Enexpected error: %s", err)
	}
	return apiGen
}

func TestLoadGenDetails(t *testing.T) {
	_ = loadGen(t)
}
