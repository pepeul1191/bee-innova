<div class="container-fluid h-100">
  <div class="row h-100">
    <!-- Lado izquierdo: Fondo de color sólido -->
    <div class="col-md-6 d-none d-md-block bg-secondary"></div>

    <!-- Lado derecho: Formulario -->
    <div class="col-md-6 d-flex align-items-center justify-content-center">
      <div class="w-50">
        <h2 class="mb-4">Iniciar Sesión</h2>
        <!-- Formulario -->
        <form action="/sign-in" method="post" class="needs-validation" id="sign-in-form">
           <input type="hidden" name="_csrf" value="{{.CsrfToken}}">

          <!-- Mensaje de error (inicialmente oculto, manejado por JS si es necesario) -->
          <div class="alert alert-danger <% unless flash[:alert] %> d-none <% end %> text-center" role="alert" id="message">
            <% if flash[:alert] %>
              <%= flash[:alert] %>
            <% end %>
          </div>

          <div class="mb-3">
            <label for="email" class="form-label">Correo Electrónico</label>
            <input type="email" class="form-control" id="email" name="email" placeholder="Ingresa su correo de la universidad" required>
            <div class="invalid-feedback">Por favor, ingresa un correo válido.</div>
          </div>

          <div class="mb-3">
            <label for="password" class="form-label">Contraseña</label>
            <input type="password" class="form-control" id="password" name="password" placeholder="Contraseña" required>
            <div class="invalid-feedback">Por favor, ingresa tu contraseña.</div>
          </div>

          <div class="d-grid gap-2 mb-3">
            <button type="submit" class="btn btn-primary">Iniciar Sesión</button>
            <a href="/auth/google_oauth2" class="btn btn-outline-secondary btn-google" id="google-sign-in">
              <i class="fa fa-google"></i> Iniciar Sesión con Google
            </a>
          </div>

          <div class="text-center">
            <a href="/sign-up" class="d-block mb-2">Crear cuenta</a>
            <a href="/reset-password" class="d-block">Recuperar contraseña</a>
            <hr>
            <a href="/" class="d-block">Regresar a Innova ULima</a>
          </div>
        </form>
      </div>
    </div>
  </div>
</div>
