<!-- views/partials/sidebar.tpl -->
<div class="sidebar bg-primary text-white">
    <div class="sidebar-header p-3">
        <h3><i class="fa fa-lightbulb-o" style="margin-right: 0.5rem;"></i>Bee Innova</h3>
    </div>

    <ul class="nav flex-column">
        <li class="nav-item">
            <a class="nav-link {{if eq .Navlink "management"}}active{{end}}" href="/management">
                <i class="fa fa-home me-2"></i> Inicio
            </a>
        </li>

        <li class="nav-item">
            <a class="nav-link {{if eq .Navlink "locations"}}active{{end}}" href="/management/locations">
                <i class="fa fa-map-marker me-2"></i> Locaciones
            </a>
        </li>

        <li class="nav-item">
            <a class="nav-link {{if eq .Navlink "tags"}}active{{end}}" href="/management/tags">
                <i class="fa fa-tag me-2"></i> Etiquetas
            </a>
        </li>

        <li class="nav-item">
            <a class="nav-link {{if eq .Navlink "enterprises"}}active{{end}}" href="/management/enterprises">
                <i class="fa fa-industry me-2"></i> Empresas
            </a>
        </li>

        <li class="nav-item">
            <a class="nav-link {{if eq .Navlink "employees"}}active{{end}}" href="/management/employees">
                <i class="fa fa-user me-2"></i> Empleados
            </a>
        </li>

        <li class="nav-item">
            <a class="nav-link {{if eq .Navlink "assets"}}active{{end}}" href="/management/assets">
                <i class="fa fa-cube me-2"></i> Activos
            </a>
        </li>

        <li class="nav-item">
            <a class="nav-link {{if eq .Navlink "roles"}}active{{end}}" href="/management/roles">
                <i class="fa fa-list me-2"></i> Roles de Usuarios
            </a>
        </li>
        
        <li class="nav-item">
            <a class="nav-link {{if eq .Navlink "users"}}active{{end}}" href="/management/users">
                <i class="fa fa-users me-2"></i> Usuarios
            </a>
        </li>
        
        <li class="nav-item">
            <a class="nav-link" href="#productosSubmenu" data-bs-toggle="collapse">
                <i class="fa fa-cubes me-2"></i> Productos
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
                <i class="fa fa-cog me-2"></i> Configuración
            </a>
        </li>
    </ul>
</div>
