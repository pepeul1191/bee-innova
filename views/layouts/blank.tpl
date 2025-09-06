<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>{{.PageTitle}} | Mi Proyecto Beego</title>
  <link rel="stylesheet" href="/static/css/style.css">
  <link rel="icon" href="/static/icon.png" type="image/png">
  <link rel="icon" href="/icon.svg" type="image/svg+xml">
  <link rel="apple-touch-icon" href="/static/icon.png">
  {{ if .Styles }}
    {{ styles .Styles }}
  {{ end }}
</head>
<body>
  <header>
    <h1>{{.PageTitle}}</h1>
  </header>
  {{.LayoutContent}}  
  <footer>
    <p>&copy; 2023 Mi Proyecto Beego</p>
  </footer>
  <script src="/static/js/main.js"></script>
  {{ if .Scripts }}
    {{ scripts .Scripts }}
  {{ end }}
</body>
</html>