<!DOCTYPE html>
<html>
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
  <body class="d-flex flex-column min-vh-100">
    <!-- Navbar -->
    <header>
      <nav class="navbar navbar-expand-lg navbar-light bg-light">
        <div class="container">
          <a class="navbar-brand" href="/">Innova ULima</a>
          <button class="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbarNav">
            <span class="navbar-toggler-icon"></span>
          </button>
          <div class="collapse navbar-collapse" id="navbarNav">
            <ul class="navbar-nav ms-auto">
              <li class="nav-item"><a class="nav-link {{ if eq .Navlink "index" }}active{{ end }}" href="/">Inicio</a></li>
              <li class="nav-item"><a class="nav-link {{ if eq .Navlink "about" }}active{{ end }}" href="/about">Sobre</a></li>
              <li class="nav-item"><a class="nav-link {{ if eq .Navlink "contact" }}active{{ end }}" href="/contact">Contacto</a></li>
              <li class="nav-item"><a class="nav-link" href="/sign-in">Ingresar</a></li>
            </ul>
          </div>
        </div>
      </nav>
    </header>
    <!-- yield -->
    {{.LayoutContent}}  
    <!-- Footer -->
    <footer class="bg-light py-4 mt-auto">
      <div class="container text-center">
        <p class="mb-0">&copy; 2025 Innova ULima. Todos los derechos reservados.</p>
        <div class="mt-2">
          <a href="/privacy" class="text-decoration-none me-3">Políticas</a>
          <a href="/terms-and-conditions" class="text-decoration-none">Términos</a>
        </div>
      </div>
    </footer>
    {{ if .Scripts }}
      {{ scripts .Scripts }}
    {{ end }}
  </body>
</html>
