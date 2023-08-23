package api

import "errors"

func (apiSpec *ApiSpec) GetPaths() error {
	if apiSpec.Paths == nil {
		for path := range apiSpec.Api.Paths {
			restPath := Path{
				Path: path,
			}
			apiSpec.Paths = append(apiSpec.Paths, restPath)
		}
	} else {
		return errors.New("Paths have been extracted for the current Spec")
	}

	return nil
}
