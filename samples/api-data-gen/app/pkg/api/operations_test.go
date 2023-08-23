package api

import (
	"math/rand"
	"testing"
)

func loadAllOperations(t *testing.T) *ApiSpec {
	apiSpec := loadPaths(t)
	apiSpec.LoadAllOperations()

	return apiSpec
}

func TestLoadAllOperations(t *testing.T) {
	_ = loadAllOperations(t)
}

func TestPrintAllOperations(t *testing.T) {
	apiSpec := loadAllOperations(t)
	apiSpec.PrintAllOperations()
}

func TestLoadOperations(t *testing.T) {
	apiSpec := loadAllOperations(t)
	randomPath := rand.Intn(len(apiSpec.Paths))
	apiSpec.LoadOperations(apiSpec.Paths[randomPath].Path)
}

func TestPrintOperations(t *testing.T) {
	apiSpec := loadAllOperations(t)
	randomPath := rand.Intn(len(apiSpec.Paths))
	apiSpec.PrintOperations(apiSpec.Paths[randomPath].Path)
}
