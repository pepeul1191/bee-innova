package common

import (
	"bee-innova/conf"
)

func GetIndexStylesHelper() []string {
	// Define los activos estáticos para cada entorno
	assets := []conf.AssetGroup{
		{
			Env:   "DEV",
			Files: []string{"bootstrap", "fontawesome"},
		},
		{
			Env:   "PRD",
			Files: []string{"bootstrap.min", "fontawesome.min"},
		},
	}

	return conf.GetAssetGroup(assets)
}

func GetIndexScriptsHelper() []string {
	assets := []conf.AssetGroup{
		{
			Env:   "DEV",
			Files: []string{"jquery", "app"},
		},
		{
			Env:   "PRD",
			Files: []string{"jquery.min", "app.min"},
		},
	}

	return conf.GetAssetGroup(assets)
}
