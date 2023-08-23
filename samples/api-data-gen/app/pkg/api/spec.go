package api

import (
	"github.com/deepmap/oapi-codegen/pkg/util"
)

func (apiSpec *ApiSpec) GetSpec(uri string) error {
	var err error
	if uri != apiSpec.Uri || apiSpec == nil {
		apiSpec.Uri = uri
		apiSpec.Api, err = util.LoadSwagger(uri)
	}
	return err
}
