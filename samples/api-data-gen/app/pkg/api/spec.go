package api

import (
	"github.com/deepmap/oapi-codegen/pkg/util"
)

func (apiSpec *ApiSpec) LoadSpec(uri string) error {
	var err error
	if uri != apiSpec.Uri || apiSpec == nil {
		apiSpec.Uri = uri
		// Load the API spec
		apiSpec.Api, err = util.LoadSwagger(uri)
	}

	return err
}

func (apiSpec *ApiSpec) LoadSpecDetails(uri string) error {
	var err error
	if uri != apiSpec.Uri || apiSpec == nil {
		apiSpec.Uri = uri
		// Load the API spec
		apiSpec.Api, err = util.LoadSwagger(uri)
		// Load the paths or REST Api endpoints
		apiSpec.LoadPaths()
		// Load All the operations for each path
		apiSpec.LoadAllOperations()
	}
	return err
}
