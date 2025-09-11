    <h3 class="mb-4">
      <i class="fa fa-list me-2"></i>Gestión de Roles
    </h3>

    

    <div class="card mb-4">
      <div class="card-header">
        <h6 class="mb-0">
          <i class="fa fa-filter me-2"></i>
          Filtros de Búsqueda
        </h6>
      </div>
      <div class="card-body">
        <form method="GET" action="/management/roles/">
          <input type="hidden" name="per_page" value="10">
          <div class="row mb-3 align-items-end">
            <div class="col-md-3">
              <label for="nombre" class="form-label">Nombre</label>
              <div class="input-group">
                <input type="text" id="nombre" name="name" class="form-control" placeholder="Buscar..." value="">
              </div>
            </div>
            <div class="col-md-3">
              <div class="input-group">
                <button type="submit" class="btn btn-primary">
                  <i class="fa fa-search"></i> Buscar
                </button>
                <a href="/management/roles/" class="btn btn-secondary" style="margin-left: 10px;">
                  <i class="fa fa-refresh"></i> Limpiar
                </a>
              </div>
            </div>
          </div>
        </form>        
      </div>
    </div>
    
    <div class="card mb-4">
      <div class="card-header d-flex justify-content-between align-items-center">
        <h6 class="mb-0">
          <i class="fa fa-list me-2"></i>
          Listado de Roles
        </h6>
      </div>
      <div class="card-body">
        <div class="d-flex justify-content-between align-items-center mb-3">
          <div class="d-flex align-items-center me-3">
            <label for="rows-per-page" class="form-label mb-0 me-2">Filas por página:</label>
            <form method="GET" id="per_page_form">
              <input type="hidden" name="name" value="">
              <select name="per_page" class="form-select" style="width: 120px;" onchange="this.form.submit()">
                <option value="5">5</option>
                <option value="10" selected="">10</option>
                <option value="25">25</option>
                <option value="50">50</option>
                <option value="100">100</option>
              </select>
            </form>
          </div>
          <div class="d-flex gap-2">
            <a href="/management/roles/add" class="btn btn-primary">
              <i class="fa fa-plus"></i> Agregar Rol
            </a>
          </div>
        </div>

        <div class="table-responsive">
          <table class="table table-striped table-hover">
            <thead>
              <tr>
                <th>Nombre</th>
                <th class="text-end">Acciones</th>
              </tr>
            </thead>
            <tbody>
              
                <tr>
                  <td colspan="2" class="text-center">No se encontraron roles.</td>
                </tr>
              
            </tbody>
            <tfoot>
              <tr>
                <td colspan="2">
                  <div class="d-flex justify-content-between align-items-center mt-3">
                    <div class="text-left">
                      Página 1 de 0 - Mostrando registros 1 - 0 de un total de 0
                    </div>

                    <nav aria-label="Page navigation">
                      <ul class="pagination mb-0">
                        <li class="page-item disabled">
                          <a class="page-link" href="/management/roles/?page=1&amp;per_page=10&amp;name=" aria-label="First">
                            <i class="fa fa-angle-double-left"></i> Primero
                          </a>
                        </li>

                        <li class="page-item disabled">
                          <a class="page-link" href="/management/roles/?page=0&amp;per_page=10&amp;name=" aria-label="Previous">
                            <i class="fa fa-angle-left"></i> Anterior
                          </a>
                        </li>

                        <li class="page-item ">
                          <a class="page-link" href="/management/roles/?page=2&amp;per_page=10&amp;name=" aria-label="Next">
                            Siguiente <i class="fa fa-angle-right"></i>
                          </a>
                        </li>

                        <li class="page-item ">
                          <a class="page-link" href="/management/roles/?page=0&amp;per_page=10&amp;name=" aria-label="Last">
                            Último <i class="fa fa-angle-double-right"></i>
                          </a>
                        </li>
                      </ul>
                    </nav>
                  </div>
                </td>
              </tr>
            </tfoot>
          </table>
        </div>
      </div>
    </div>
