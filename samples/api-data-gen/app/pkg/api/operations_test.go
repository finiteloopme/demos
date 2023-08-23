package api

import "testing"

func loadAllOperations(t *testing.T) *ApiSpec {
	apiSpec := loadPaths(t)
	apiSpec.LoadAllOperations()

	return apiSpec
}

func TestLoadAllOperations(t *testing.T) {
	_ = loadAllOperations(t)
}
