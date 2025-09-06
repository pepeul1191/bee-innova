package common

import (
	"bee-innova/conf"
)

func GetError404StylesHelper() []string {
	// Define los activos estáticos para cada entorno
	assets := []conf.AssetGroup{
		{
			Env:   "DEV",
			Files: []string{"dist/web.min", "common/css/error"},
		},
		{
			Env:   "PRD",
			Files: []string{"dist/web.min"},
		},
	}

	return conf.GetAssetGroup(assets)
}
