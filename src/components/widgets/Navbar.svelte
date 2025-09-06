<!-- src/components/Navbar.svelte -->
<script>
  export let username = '';
  export let userImage = '';
</script>

<nav class="navbar navbar-light bg-light top-navbar">
  <div class="container-fluid">
    <div class="d-flex align-items-center ms-auto">
      <div class="dropdown user-menu">
        <button class="btn btn-link dropdown-toggle d-flex align-items-center text-white text-decoration-none user-name"
                type="button" id="userMenu" data-bs-toggle="dropdown" aria-expanded="false">
          <span class="me-2 d-none d-md-inline">{username}</span>
          <img src={userImage} alt="Foto de perfil" />
        </button>
        <ul class="dropdown-menu dropdown-menu-end" aria-labelledby="userMenu">
          <li><a class="dropdown-item" href="/session"><i class="fas fa-user me-2"></i>Ver perfil</a></li>
          <li><a class="dropdown-item" href="/setup"><i class="fas fa-cog me-2"></i>Configuración</a></li>
          <li><hr class="dropdown-divider"></li>
          <li><a class="dropdown-item text-danger" href="/sign-out"><i class="fas fa-sign-out-alt me-2"></i>Cerrar sesión</a></li>
        </ul>
      </div>
    </div>
  </div>
</nav>
