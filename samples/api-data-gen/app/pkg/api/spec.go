package api

import (
	"github.com/deepmap/oapi-codegen/pkg/util"
	"github.com/finiteloopme/demos/samples/api-data-gen/app/pkg/types"
)

func (apiSpec *ApiSpec) LoadSpec(uri string) error {
	var err error
	if uri != apiSpec.Uri.String() || apiSpec == nil {
		apiSpec.Uri = types.UriType(uri)
		// Load the API spec
		apiSpec.Api, err = util.LoadSwagger(uri)
	}

	return err
}

func (apiSpec *ApiSpec) LoadSpecDetails(uri string) error {
	var err error
	if uri != apiSpec.Uri.String() || apiSpec == nil {
		apiSpec.Uri = types.UriType(uri)
		// Load the API spec
		apiSpec.Api, err = util.LoadSwagger(uri)
		// Load the paths or REST Api endpoints
		apiSpec.LoadPaths()
		// Load All the operations for each path
		apiSpec.LoadAllOperations()
	}
	return err
}
