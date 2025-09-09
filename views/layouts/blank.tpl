<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>{{.PageTitle}}</title>
  <link rel="icon" href="/static/icon.png" type="image/png">
  <link rel="apple-touch-icon" href="/static/icon.png">
  {{ if .Styles }}
    {{ styles .Styles }}
  {{ end }}
</head>
<body>
  {{.LayoutContent}}  
  {{ if .Scripts }}
    {{ scripts .Scripts }}
  {{ end }}
</body>
</html>