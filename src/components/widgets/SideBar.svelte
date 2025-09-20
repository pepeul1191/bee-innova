<script>
  // src/components/widgets/SideBar.svelte
  import Footer from './Footer.svelte';
</script>
<style>
  .sidebar{
    padding: 0px;
    top: 58px;
  }
</style>

<nav id="sidebarMenu" class="col-md-3 col-lg-2 d-md-block bg-dark sidebar collapse">
    <div class="position-sticky pt-3">
        <ul class="nav flex-column">
            <li class="nav-item">
                <a class="nav-link active" aria-current="page" href="#">
                    Inicio
                </a>
            </li>
            <li class="nav-item">
                <a class="nav-link" href="#">
                    Usuarios
                </a>
            </li>
            <li class="nav-item">
                <a class="nav-link" href="#">
                    Configuración
                </a>
            </li>
        </ul>
    </div>
</nav>