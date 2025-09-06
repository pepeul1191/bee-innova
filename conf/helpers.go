package conf

import (
	"fmt"
	"html/template"
	"os"
	"strings"
)

type AssetGroup struct {
	Env   string
	Files []string
}

func GetAssetGroup(assetsGroup []AssetGroup) []string {
	// Obtiene el valor de la variable de entorno
	env := os.Getenv("STATIC_ENV")

	// Si la variable no está definida, usa "DEV" como valor por defecto
	if env == "" {
		env = "DEV"
	}

	// Itera sobre el slice de assets para encontrar el entorno correcto
	for _, assets := range assetsGroup {
		if assets.Env == env {
			return assets.Files
		}
	}

	// Si el entorno no coincide con ninguno, devuelve un slice vacío
	return []string{}
}

// GenerateStylesHTML genera etiquetas <link> para hojas de estilo.
func GenerateStylesHTML(styles []string) template.HTML {
	// Obtiene el valor de la variable de entorno
	staticURL := os.Getenv("STATIC_URL")
	// Si la variable no está definida, usa "DEV" como valor por defecto
	if staticURL == "" {
		staticURL = "/static/"
	}

	var html strings.Builder
	for _, style := range styles {
		style = strings.Trim(style, "/")
		if !strings.HasSuffix(style, ".css") {
			style += ".css"
		}
		html.WriteString(fmt.Sprintf(`<link rel="stylesheet" href="%s%s">`, staticURL, style))
	}

	return template.HTML(html.String())
}

// GenerateScriptsHTML genera etiquetas <script> para archivos JavaScript.
func GenerateScriptsHTML(scripts []string) template.HTML {
	// Obtiene el valor de la variable de entorno
	staticURL := os.Getenv("STATIC_URL")
	// Si la variable no está definida, usa "DEV" como valor por defecto
	if staticURL == "" {
		staticURL = "/static/"
	}

	var html strings.Builder
	for _, script := range scripts {
		script = strings.Trim(script, "/")
		if !strings.HasSuffix(script, ".js") {
			script += ".js"
		}
		html.WriteString(fmt.Sprintf(`<script src="%s%s"></script>`, staticURL, script))
	}

	return template.HTML(html.String())
}

func GetEnv(key string) string {
	return os.Getenv(key)
}
