<!-- views/partials/_navbar.tpl -->
<nav class="navbar navbar-expand navbar-dark bg-dark fixed-top">
    <div class="container-fluid">
        <!-- Botón sidebar -->
        <button class="btn btn-link text-white" id="sidebarToggle">
            <i class="fa fa-bars"></i>
        </button>

        <!-- Elementos derecha -->
        <ul class="navbar-nav ms-auto">
            <!-- Notificaciones -->
            <li class="nav-item dropdown">
                <a class="nav-link dropdown-toggle" href="#" id="notificationsDropdown" role="button" data-bs-toggle="dropdown" aria-expanded="false">
                    <i class="fa fa-bell"></i>
                    <span class="badge bg-danger">3</span>
                </a>
                <ul class="dropdown-menu dropdown-menu-end" aria-labelledby="notificationsDropdown">
                    <li><a class="dropdown-item" href="#">Notificación 1</a></li>
                    <li><a class="dropdown-item" href="#">Notificación 2</a></li>
                    <li><hr class="dropdown-divider"></li>
                    <li><a class="dropdown-item" href="#">Ver todas</a></li>
                </ul>
            </li>
            
            <!-- Usuario (se actualizará con JavaScript) -->
            <li class="nav-item dropdown">
                <a class="nav-link dropdown-toggle" href="#" id="userDropdown" role="button" data-bs-toggle="dropdown" aria-expanded="false">
                    <i class="fa fa-user"></i> 
                    <span class="user-name" id="userName">Cargando...</span>
                </a>
                <ul class="dropdown-menu dropdown-menu-end" aria-labelledby="userDropdown">
                    <li><a class="dropdown-item" href="#"><i class="fa fa-user me-2"></i> Perfil</a></li>
                    <li><hr class="dropdown-divider"></li>
                    <li>
                        <div id="authSection">
                            <!-- Se cargará dinámicamente con JavaScript -->
                            <a class="dropdown-item" href="/sign-in">
                                <i class="fa fa-sign-in me-2"></i> Iniciar sesión
                            </a>
                        </div>
                    </li>
                </ul>
            </li>
        </ul>
    </div>
</nav>
