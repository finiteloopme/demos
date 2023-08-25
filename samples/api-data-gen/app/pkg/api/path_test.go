package api

import (
	"math/rand"
	"testing"
)

func loadPaths(t *testing.T) *ApiSpec {

	apiSpec := loadSpec(t)
	apiSpec.LoadPaths()
	if len(apiSpec.Paths) < 1 {
		t.Fatalf("Expected the extracted paths to be more than 1.  Received: %v", len(apiSpec.Paths))
	}
	return apiSpec
}
func TestLoadPaths(t *testing.T) {
	_ = loadPaths(t)
}

func TestGetPath(t *testing.T) {
	apiSpec := loadPaths(t)
	randomIdx := rand.Intn(len(apiSpec.Paths))
	_, err := apiSpec.GetPath(apiSpec.Paths[randomIdx].Path)
	if err != nil {
		t.Fatalf("Unexpected error: %s", err)
	}
}

func TestPrintPaths(t *testing.T) {
	apiSpec := loadPaths(t)
	apiSpec.PrintPaths()
}

func TestPrintPath(t *testing.T) {
	apiSpec := loadPaths(t)
	randomIdx := rand.Intn(len(apiSpec.Paths))
	apiSpec.PrintPath(apiSpec.Paths[randomIdx].Path)
}
