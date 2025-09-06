<script>
  import { onMount } from 'svelte';
  import { navigate } from 'svelte-routing';
  import IssueHeaderTab from '../../forms/IssueHeaderTab.svelte';
  import IssueAssetTab from '../../forms/IssueAssetTab.svelte';
  import InputUpload from '../../widgets/InputUpload.svelte';
  import { fetchOneById, create, update } from '../../../services/app/teacher_services.js';
  import { fetchAll } from '../../../services/app/document_types_service.js';
  import { activeLink } from '../../../helpers/store'; 

  export let title = 'Nuevo Docente';
  export let _id = '';
  let issueDetailModalInstance;
  let issueFormInstance;
  let issueDetailModal;
  let headerTabInstance;
  let assetTabInstance;
  let alertMessage = {
    text: '',
    status: '',
  };
  let issue = {};

  let documentTypes = [];

  let teacher = {
    _id: '',
    document_type_id: '',
    names: '',
    last_names: '',
    document_type_id: '',
    document_number: '',
    code: '',
    email: '',
    image_url: `user-default.png`,
  };
  
  let modalTitle;

  const listDocumentTypes = () => {
    fetchAll(URLS.CLASSROOM_SERVICE, 'jwtClassroomToken')
      .then(response => {
        console.log(response)
        documentTypes = response.data;
      })
      .catch(error => {
        console.error('Error:', error);
      });
  }

  onMount(() => {
    activeLink.set('/teachers');
    document.title = "Nueva Docente";

    listDocumentTypes();

    if (_id != ''){
      document.title = "Editar Docente";
      title = 'Editar Docente'
      fetchOneById(URLS.TICKETS_SERVICE, 'jwtTicketsToken', _id)
        .then(response => {
          issue = response.data;
          headerTabInstance.issue = issue;
          headerTabInstance.updateView();
          headerTabInstance.loadMarkDownEditor();
        })
        .catch(error => {
          console.error('Error:', error);
        });
    } else {
      headerTabInstance.loadMarkDownEditor();
    }
  });

  const afterUpload = (data, status) => {
    if(status == 'success'){
      teacher.image_url = `${data.path}`;  
      alertMessage = {
        text: 'Se ha subido correctamente la imagen del profesor',
        status: 'success',
      };  
    }else{
      console.error(data);
      alertMessage = {
        text: 'Error al subir la imagen del profesor',
        status: 'danger',
      };  
    }

    console.log(alertMessage)
  }
  
  const saveForm = () => {
    if(teacher._id == ''){
      create(URLS.CLASSROOM_SERVICE, 'jwtClassroomToken', teacher)
        .then(response => {
          console.log(response)
          teacher._id = response.data._id;
          console.log(teacher)
          alertMessage = {
            text: 'Se ha creado un nuevo profesor',
            status: 'success',
          };  
          title = 'Editar Profesor'
        })
        .catch(error => {
          console.error('Error:', error);
          alertMessage = {
            text: 'Error al crear el profesor',
            status: 'danger',
          };  
        });
    }else{
      update(URLS.CLASSROOM_SERVICE, 'jwtClassroomToken', teacher)
        .then(response => {
          console.log(response)
          teacher._id = response.data._id;
        })
        .catch(error => {
          console.error('Error:', error);
        });
    }
  }

  const openImage = () => {
    window.open(teacher.image_url, "_blank");
  }
</script>
<style>
  #imagePreview:hover{
    cursor:pointer;
  }
</style>

<div class="row">
  <h1 class="mb-2 subtitle">
    <a class="return-link" class:active={$activeLink === '/teachers'} href="/teachers" on:click|preventDefault={() => {navigate('/teachers')}}>
      <i class="fa fa-graduation-cap me-2"></i>Gestión de Docentes
    </a> / {title}
  </h1>
</div>
<hr>
{#if alertMessage.text != ''}
  <div class="alert alert-{alertMessage.status}" role="alert">
    {alertMessage.text}
  </div>
{/if}
<!-- Sección de información básica -->
<div class="card mb-4">
  <div class="card-header bg-secondary text-white">
    <i class="fa fa-info-circle"></i> Información Básica
  </div>
  <div class="card-body">
    <div class="row g-3">
      <div class="col-md-3">
        <label for="names" class="form-label">Nombres</label>
        <input type="text" class="form-control" id="names" required bind:value={teacher.names}>
      </div>
      <div class="col-md-3">
        <label for="lastNames" class="form-label">Apellidos</label>
        <input type="text" class="form-control" id="lastNames" required bind:value={teacher.last_names}>
      </div>
      <div class="col-md-2">
        <label for="documentTypeId" class="form-label">Tipo de Documento</label>
        <select id="documentTypeId" class="form-select" bind:value={teacher.document_type_id}>
          <option value="" disabled selected>Seleccione Tipo de Documento</option>
          {#each documentTypes as dt}
            <option value="{dt._id}">{dt.name}</option>
          {/each}
        </select>
      </div>
      <div class="col-md-2">
        <label for="documentNumber" class="form-label">Número de Documento</label>
        <input type="text" class="form-control" id="documentNumber" required bind:value={teacher.document_number}>
      </div>
      <div class="col-md-2">
        <label for="code" class="form-label">Código de Docente</label>
        <input type="text" class="form-control" id="code" required bind:value={teacher.code}>
      </div>
    </div>
    <div class="row mt-3">
      <div class="col-md-6 d-flex align-items-center gap-3">
        <img id="imagePreview" src="{URLS.FILES_SERVICE}{teacher.image_url}" width="90" height="90" alt="Vista previa" class="img-thumbnail me-3" on:click={openImage}>
        <div class="flex-grow-1">
          <InputUpload
            label="Subir fotografía"
            acceptedFormats={['jpg', 'jpeg', 'png']}
            maxSizeMB={2}
            url="{URLS.FILES_SERVICE}/api/v1/public-files"
            fileKey="file"
            extraParams={{folder: 'teachers'}}
            jwt={localStorage.getItem("jwtFilesToken")}
            onSuccess={(data) => afterUpload(data, 'success')}
            onError={(error) => afterUpload(error, 'error')}
          />
        </div>
      </div>
      <div class="col-md-4">
        <label for="email" class="form-label">Correo</label>
        <input type="email" class="form-control" id="email" required bind:value={teacher.email}>
      </div>
      <div class="col-md-2 position-relative" style="min-height: 100px;">
        <!-- Contenido de la columna (si lo hay) -->
        <!-- Botón posicionado abajo a la derecha -->
        <div class="position-absolute bottom-0 end-0 p-2" style="padding-bottom: 25px !important;">
          <button class="btn btn-primary mb-3" on:click={saveForm}>
            <i class="fa fa-save"></i> Guardar Cambios
          </button>
        </div>
      </div>
    </div>
  </div>
</div>

<!-- Sección de Usuario -->
<div class="card mb-4">
  <div class="card-header bg-secondary text-white d-flex justify-content-between align-items-center">
    <span><i class="fa fa-map-marker"></i> Usuario del Docente</span>
  </div>
  <div class="card-body">
    <div class="col-md-4">
      <label for="teacherCode" class="form-label">Usuario Asociado</label>
      <div class="input-group">
        <input type="text" class="form-control" id="teacherCode" disabled placeholder="Código del docente">
        <button class="btn btn-primary" type="button" id="searchTeacherBtn" title="Buscar docente">
          <i class="fa fa-search"></i> Buscar
        </button>
        <button class="btn btn-success" type="button" id="associateTeacherBtn" title="Asociar docente">
          <i class="fa fa-link"></i> Asociar
        </button>
        <button class="btn btn-danger" type="button" id="removeTeacherBtn" title="Remover docente">
          <i class="fa fa-unlink"></i> Remover
        </button>
      </div>
      <small class="text-muted">Busque y asocie un código de docente</small>
    </div>
  </div>
</div>

<!-- Sección de Direcciones -->
<div class="card mb-4">
  <div class="card-header bg-secondary text-white d-flex justify-content-between align-items-center">
    <span><i class="fa fa-map-marker"></i> Direcciones</span>
  </div>
  <div class="card-body">
    <div class="table-responsive">
      <table class="table table-striped" id="addressTable">
        <thead>
          <tr>
            <th>Dirección</th>
            <th>Ubicación</th>
            <th class="table-actions">Acciones</th>
          </tr>
        </thead>
        <tbody>
          <!-- Ejemplo de dirección -->
          <tr>
            <td>Calle Principal 123</td>
            <td>Sede Central</td>
            <td>
              <button class="btn btn-sm btn-outline-danger delete-address">
                <i class="fa fa-trash"></i> Eliminar
              </button>
            </td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</div>

<!-- Sección de Teléfonos -->
<div class="card mb-4">
  <div class="card-header bg-secondary text-white d-flex justify-content-between align-items-center">
    <span><i class="fa fa-phone"></i> Teléfonos</span>
  </div>
  <div class="card-body">
    <div class="table-responsive">
      <table class="table table-striped" id="phoneTable">
        <thead>
          <tr>
            <th>Número</th>
            <th>Tipo</th>
            <th class="table-actions">Acciones</th>
          </tr>
        </thead>
        <tbody>
          <!-- Ejemplo de teléfono -->
          <tr>
            <td>987654321</td>
            <td>Celular</td>
            <td>
              <button class="btn btn-sm btn-outline-danger delete-phone">
                <i class="fa fa-trash"></i> Eliminar
              </button>
            </td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</div>

<!-- Botones de acción -->
<div class="d-flex justify-content-end mb-4">
  <button type="button" class="btn btn-outline-secondary me-2">
    <i class="fa fa-times"></i> Cancelar
  </button>
  <button type="submit" class="btn btn-primary">
    <i class="fa fa-save"></i> Guardar Docente
  </button>
</div>

<!-- Modal para selección de usuario -->
<div class="modal fade" id="userModal" tabindex="-1" aria-labelledby="userModalLabel" aria-hidden="true">
  <div class="modal-dialog modal-lg">
    <div class="modal-content">
      <div class="modal-header">
        <h5 class="modal-title" id="userModalLabel"><i class="fa fa-users"></i> Seleccionar Usuario</h5>
        <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
      </div>
      <div class="modal-body">
        <div class="input-group mb-3">
          <input type="text" class="form-control" placeholder="Buscar usuario..." id="userSearch">
          <button class="btn btn-outline-secondary" type="button" id="searchUserBtn">
            <i class="fa fa-search"></i> Buscar
          </button>
        </div>
        <div class="table-responsive">
          <table class="table table-hover">
            <thead>
              <tr>
                <th>ID</th>
                <th>Email</th>
                <th>Nombre</th>
                <th>Seleccionar</th>
              </tr>
            </thead>
            <tbody id="userTableBody">
              <tr>
                <td>1</td>
                <td>profesor1@escuela.edu</td>
                <td>Juan Pérez</td>
                <td>
                  <button type="button" class="btn btn-sm btn-outline-primary select-user" 
                          data-id="1" data-email="profesor1@escuela.edu" data-name="Juan Pérez">
                    <i class="fa fa-check"></i> Seleccionar
                  </button>
                </td>
              </tr>
              <tr>
                <td>2</td>
                <td>profesor2@escuela.edu</td>
                <td>María Gómez</td>
                <td>
                  <button type="button" class="btn btn-sm btn-outline-primary select-user" 
                          data-id="2" data-email="profesor2@escuela.edu" data-name="María Gómez">
                    <i class="fa fa-check"></i> Seleccionar
                  </button>
                </td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>
      <div class="modal-footer">
        <button type="button" class="btn btn-secondary" data-bs-dismiss="modal">
          <i class="fa fa-times"></i> Cerrar
        </button>
      </div>
    </div>
  </div>
</div>

<!-- Modal para agregar dirección -->
<div class="modal fade" id="addressModal" tabindex="-1" aria-hidden="true">
  <div class="modal-dialog">
    <div class="modal-content">
      <div class="modal-header">
        <h5 class="modal-title"><i class="fa fa-map-marker"></i> Agregar Dirección</h5>
        <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
      </div>
      <div class="modal-body">
        <form id="addressForm">
          <div class="mb-3">
            <label for="addressName" class="form-label">Nombre/Descripción</label>
            <input type="text" class="form-control" id="addressName" required>
          </div>
          <div class="mb-3">
            <label for="addressLocation" class="form-label">Ubicación</label>
            <select class="form-select" id="addressLocation" required>
              <option value="" selected disabled>Seleccione una ubicación</option>
              <option value="1">Sede Principal</option>
              <option value="2">Sede Norte</option>
              <option value="3">Sede Sur</option>
            </select>
          </div>
        </form>
      </div>
      <div class="modal-footer">
        <button type="button" class="btn btn-secondary" data-bs-dismiss="modal">
          <i class="fa fa-times"></i> Cancelar
        </button>
        <button type="button" class="btn btn-primary" id="saveAddressBtn">
          <i class="fa fa-save"></i> Guardar Dirección
        </button>
      </div>
    </div>
  </div>
</div>

<!-- Modal para agregar teléfono -->
<div class="modal fade" id="phoneModal" tabindex="-1" aria-hidden="true">
  <div class="modal-dialog">
    <div class="modal-content">
      <div class="modal-header">
        <h5 class="modal-title"><i class="fa fa-phone"></i> Agregar Teléfono</h5>
        <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
      </div>
      <div class="modal-body">
        <form id="phoneForm">
          <div class="mb-3">
            <label for="phoneNumber" class="form-label">Número de Teléfono</label>
            <input type="tel" class="form-control" id="phoneNumber" required>
          </div>
          <div class="mb-3">
            <label for="phoneType" class="form-label">Tipo de Teléfono</label>
            <select class="form-select" id="phoneType" required>
              <option value="" selected disabled>Seleccione un tipo</option>
              <option value="Celular">Celular</option>
              <option value="Casa">Casa</option>
              <option value="Oficina">Oficina</option>
              <option value="Otro">Otro</option>
            </select>
          </div>
        </form>
      </div>
      <div class="modal-footer">
        <button type="button" class="btn btn-secondary" data-bs-dismiss="modal">
          <i class="fa fa-times"></i> Cancelar
        </button>
        <button type="button" class="btn btn-primary" id="savePhoneBtn">
          <i class="fa fa-save"></i> Guardar Teléfono
        </button>
      </div>
    </div>
  </div>
</div>