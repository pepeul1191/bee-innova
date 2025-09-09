package common

import (
	"bee-innova/conf"
)

// DASHBOARD

func GetDashboardStylesHelper() []string {
	// Define los activos estáticos para cada entorno
	assets := []conf.AssetGroup{
		{
			Env:   "DEV",
			Files: []string{"dist/admin.min", "common/css/index"},
		},
		{
			Env:   "PRD",
			Files: []string{"dist/admin.min"},
		},
	}

	return conf.GetAssetGroup(assets)
}

func GetDashboardScriptsHelper() []string {
	assets := []conf.AssetGroup{
		{
			Env:   "DEV",
			Files: []string{"dist/admin"},
		},
		{
			Env:   "PRD",
			Files: []string{"dist/admin.min"},
		},
	}

	return conf.GetAssetGroup(assets)
}

// HOME

func GetIndexStylesHelper() []string {
	// Define los activos estáticos para cada entorno
	assets := []conf.AssetGroup{
		{
			Env:   "DEV",
			Files: []string{"dist/web.min", "common/css/index"},
		},
		{
			Env:   "PRD",
			Files: []string{"dist/web.min"},
		},
	}

	return conf.GetAssetGroup(assets)
}

func GetIndexScriptsHelper() []string {
	assets := []conf.AssetGroup{
		{
			Env:   "DEV",
			Files: []string{"dist/web"},
		},
		{
			Env:   "PRD",
			Files: []string{"dist/web.min"},
		},
	}

	return conf.GetAssetGroup(assets)
}

func GetAboutStylesHelper() []string {
	// Define los activos estáticos para cada entorno
	assets := []conf.AssetGroup{
		{
			Env:   "DEV",
			Files: []string{"dist/web.min", "common/css/about"},
		},
		{
			Env:   "PRD",
			Files: []string{"dist/web.min"},
		},
	}

	return conf.GetAssetGroup(assets)
}

func GetContactStylesHelper() []string {
	// Define los activos estáticos para cada entorno
	assets := []conf.AssetGroup{
		{
			Env:   "DEV",
			Files: []string{"dist/web.min", "common/css/contact"},
		},
		{
			Env:   "PRD",
			Files: []string{"dist/web.min"},
		},
	}

	return conf.GetAssetGroup(assets)
}
