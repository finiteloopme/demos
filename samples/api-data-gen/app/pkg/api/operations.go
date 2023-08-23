package api

import "github.com/finiteloopme/goutils/pkg/log"

func (apiSpec *ApiSpec) LoadAllOperations() {
	for _, path := range apiSpec.Paths {
		apiSpec.LoadOperations(path.Path)
	}
}

func (apiSpec *ApiSpec) LoadOperations(pathName string) {
	if apiSpec.Paths == nil {
		apiSpec.LoadPaths()
	}

	path, err := apiSpec.GetPath(pathName)
	if err == nil {
		pathItem := apiSpec.Api.Paths.Find(path.Path)
		if pathItem.Get != nil {
			path.Operations = append(path.Operations, &Operation{
				Type:        GET.String(),
				OperationID: pathItem.Get.OperationID,
			})
		}
		if pathItem.Post != nil {
			path.Operations = append(path.Operations, &Operation{
				Type:        POST.String(),
				OperationID: pathItem.Get.OperationID,
			})
		}
		if pathItem.Put != nil {
			path.Operations = append(path.Operations, &Operation{
				Type:        PUT.String(),
				OperationID: pathItem.Get.OperationID,
			})
		}
		if pathItem.Delete != nil {
			path.Operations = append(path.Operations, &Operation{
				Type:        DELETE.String(),
				OperationID: pathItem.Get.OperationID,
			})
		}
		if pathItem.Patch != nil {
			path.Operations = append(path.Operations, &Operation{
				Type:        PATCH.String(),
				OperationID: pathItem.Get.OperationID,
			})
		}
	}
	// apiSpec.AddOrUpdate(path)
}

func (apiSpec *ApiSpec) PrintAllOperations() {
	for _, path := range apiSpec.Paths {
		log.Info(path.Path)
		apiSpec.LoadOperations(path.Path)
	}
}
func (apiSpec *ApiSpec) PrintOperations(pathName string) {
	path, _ := apiSpec.GetPath(pathName)
	for _, ops := range path.Operations {
		log.Info("\tMethod: " + ops.Type)
		log.Info("\tOperationID: " + ops.OperationID)
	}
}
