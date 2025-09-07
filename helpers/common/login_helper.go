package common

import (
	"bee-innova/conf"
)

func GetLoginStylesHelper() []string {
	// Define los activos estáticos para cada entorno
	assets := []conf.AssetGroup{
		{
			Env:   "DEV",
			Files: []string{"dist/web.min", "common/css/login"},
		},
		{
			Env:   "PRD",
			Files: []string{"dist/web.min"},
		},
	}

	return conf.GetAssetGroup(assets)
}
