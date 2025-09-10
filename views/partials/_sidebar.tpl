<!-- views/partials/sidebar.tpl -->
<div class="sidebar bg-primary text-white">
    <div class="sidebar-header p-3">
        <h3><i class="fa fa-lightbulb-o" style="margin-right: 0.5rem;"></i>Bee Innova</h3>
    </div>

    <ul class="nav flex-column">
        <li class="nav-item">
            <a class="nav-link {{if eq .Navlink "management"}}active{{end}}" href="/management">
                <i class="fas fa-home me-2"></i> Inicio
            </a>
        </li>

        <li class="nav-item">
            <a class="nav-link {{if eq .Navlink "locations"}}active{{end}}" href="/management/locations">
                <i class="fas fa-map me-2"></i> Locaciones
            </a>
        </li>

        <li class="nav-item">
            <a class="nav-link {{if eq .Navlink "tags"}}active{{end}}" href="/management/tags">
                <i class="fas fa-tag me-2"></i> Etiquetas
            </a>
        </li>

        <li class="nav-item">
            <a class="nav-link {{if eq .Navlink "enterprises"}}active{{end}}" href="/management/enterprises">
                <i class="fas fa-industry me-2"></i> Empresas
            </a>
        </li>

        <li class="nav-item">
            <a class="nav-link {{if eq .Navlink "employees"}}active{{end}}" href="/management/employees">
                <i class="fas fa-user me-2"></i> Empleados
            </a>
        </li>

        <li class="nav-item">
            <a class="nav-link {{if eq .Navlink "assets"}}active{{end}}" href="/management/assets">
                <i class="fas fa-cube me-2"></i> Activos
            </a>
        </li>

        <li class="nav-item">
            <a class="nav-link {{if eq .Navlink "roles"}}active{{end}}" href="/management/roles">
                <i class="fas fa-list me-2"></i> Roles de Usuarios
            </a>
        </li>
        
        <li class="nav-item">
            <a class="nav-link {{if eq .Navlink "users"}}active{{end}}" href="/management/users">
                <i class="fas fa-users me-2"></i> Usuarios
            </a>
        </li>
        
        <li class="nav-item">
            <a class="nav-link" href="#productosSubmenu" data-bs-toggle="collapse">
                <i class="fas fa-cubes me-2"></i> Productos
            </a>
            <ul id="productosSubmenu" class="collapse nav flex-column ps-4">
                <li class="nav-item">
                    <a class="nav-link {{if eq .Navlink "products"}}active{{end}}" href="/management/products">Lista</a>
                </li>
                <li class="nav-item">
                    <a class="nav-link {{if eq .Navlink "new_product"}}active{{end}}" href="/management/new">Agregar</a>
                </li>
            </ul>
        </li>
        
        <li class="nav-item">
            <a class="nav-link {{if eq .Navlink "config"}}active{{end}}" href="/management/config">
                <i class="fas fa-cog me-2"></i> Configuración
            </a>
        </li>
    </ul>
</div>