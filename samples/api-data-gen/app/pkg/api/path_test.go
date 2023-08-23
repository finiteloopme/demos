package api

import "testing"

func TestGetPaths(t *testing.T) {
	apiSpec := &ApiSpec{}
	err := apiSpec.GetSpec(TestUri)
	if err != nil {
		t.Fatalf("Unxpected error: %s", err)
	}

	err = apiSpec.GetPaths()
	if err != nil {
		t.Fatalf("Unxpected error: %s", err)
	}
	if len(apiSpec.Paths) < 1 {
		t.Fatalf("Expected the extracted paths to be more than 1.  Received: %v", len(apiSpec.Paths))
	}
}
