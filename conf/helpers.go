package conf

import (
	"fmt"
	"html/template"
	"strings"
)

// GenerateStylesHTML genera etiquetas <link> para hojas de estilo.
func GenerateStylesHTML(baseURL string, styles []string) template.HTML {
	if baseURL == "" {
		baseURL = "/static/"
	} else if !strings.HasSuffix(baseURL, "/") {
		baseURL += "/"
	}

	var html strings.Builder
	for _, style := range styles {
		style = strings.Trim(style, "/")
		if !strings.HasSuffix(style, ".css") {
			style += ".css"
		}
		html.WriteString(fmt.Sprintf(`<link rel="stylesheet" href="%s%s">`, baseURL, style))
	}

	return template.HTML(html.String())
}

// GenerateScriptsHTML genera etiquetas <script> para archivos JavaScript.
func GenerateScriptsHTML(baseURL string, scripts []string) template.HTML {
	if baseURL == "" {
		baseURL = "/static/"
	} else if !strings.HasSuffix(baseURL, "/") {
		baseURL += "/"
	}

	var html strings.Builder
	for _, script := range scripts {
		script = strings.Trim(script, "/")
		if !strings.HasSuffix(script, ".js") {
			script += ".js"
		}
		html.WriteString(fmt.Sprintf(`<script src="%s%s"></script>`, baseURL, script))
	}

	return template.HTML(html.String())
}
