package gen

import (
	"github.com/finiteloopme/demos/samples/api-data-gen/app/pkg/api"
)

func (apiGen *ApiGen) ConfigApiGen(apiSpec *api.ApiSpec) error {
	apiGen.Api = apiSpec.Api
	apiGen.Uri = apiSpec.Uri
	apiGen.Paths = apiSpec.Paths

	return nil
}

func ConnectToGenAI(projectId string) {

}
