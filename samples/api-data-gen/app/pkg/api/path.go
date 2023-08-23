package api

import (
	"errors"

	"github.com/finiteloopme/goutils/pkg/log"
)

func (apiSpec *ApiSpec) LoadPaths() {
	if apiSpec.Paths == nil {
		for path := range apiSpec.Api.Paths {
			restPath := &Path{
				Path: path,
			}
			apiSpec.Paths = append(apiSpec.Paths, restPath)
			// load operations
			// Making a concious decision to not load operations by default
			// And explicit call to load operations is expected
			// apiSpec.LoadOperations(restPath.Path)
		}
	}

	return
}

func (apiSpec *ApiSpec) PrintPaths() {
	if apiSpec.Paths == nil {
		apiSpec.LoadPaths()
	}
	for _, path := range apiSpec.Paths {
		log.Info(path.Path)
	}
	return
}

func (apiSpec *ApiSpec) GetPath(path string) (*Path, error) {

	if apiSpec.Paths == nil {
		apiSpec.LoadPaths()
	}

	for _, _path := range apiSpec.Paths {
		if _path.Path == path {
			return _path, nil
		}
	}

	return &Path{}, errors.New("API Spec has no path: " + path)
}

// func (apiSpec *ApiSpec) AddOrUpdate(path Path) {
// 	for _, _path := range apiSpec.Paths {
// 		if _path.Path == _path.Path {
// 			// Udpate operations on the path
// 			_path.Operations = []Operation{}
// 			for _, op := range path.Operations {
// 				_path.Operations = append(_path.Operations, op)
// 			}
// 			// return after update
// 			return
// 		}
// 	}
// 	// Add the path
// 	apiSpec.Paths = append(apiSpec.Paths, path)
// }
