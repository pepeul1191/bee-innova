<script>
  import { onMount } from 'svelte';
  import { Modal } from 'bootstrap';
  import TagsFilters from '../../forms/TagFilters.svelte';
  import DataTable from '../../widgets/DataTable.svelte';
  import TagDetail from '../../forms/TagDetail.svelte';
  import { activeLink } from '../../../helpers/store'; 

  let tagDetailModalInstance;
  let tagFormInstance;
  let tagDetailModal;
  let alertMessage = {
    text: '',
    status: '',
  };
  let tagDataTable;
  let modalTitle;

  onMount(() => {
    activeLink.set('/tags');
    document.title = "Gestión de Etiquetas";
    // montar acciones de la tabla
      // ejemplos
      //tagDataTable.addButton.action = () => tagDataTable.addRow();
      //tagDataTable.addButton.action = () => tagDataTable.goToLink('/users');
      //tagDataTable.addButton.action = () => tagDataTable.goToHref(BASE_URL + 'hola');
      //tagDataTable.addButton.action = () => tagDataTable.openTab(BASE_URL + 'hola');
    tagDetailModalInstance = new Modal(tagDetailModal);
    // table action buttons
    tagDataTable.actionButtons = [
      {
        class: 'btn-secondary',
        icon: 'fa-pencil',
        label: 'Editar',
        action: editTag
      },
      {
        class: 'btn-danger',
        icon: 'fa-trash',
        label: 'Eliminar',
        action: (record) => {
          tagDataTable.askToDeleteRow(record, '_id');
        }
      },
    ];
    tagDataTable.list();
  });

  const editTag = (tag) => {
    modalTitle = 'Editar Activo'
    tagFormInstance.clean();
    tagFormInstance.loadTag(tag);
    tagDetailModalInstance.show();
  }

  const handleSearchFilter = (event) => {
    const { name, description } = event.detail;
    console.log(event.detail);
    tagDataTable.queryParams = {name,description};
    tagDataTable.list();
  }
  
  const handleCleanFilter= () => {
    tagDataTable.queryParams = {};
    tagDataTable.list();
  }

  const addTag = () => {
    modalTitle = 'Agregar Activo'
    tagFormInstance.clean();
    tagDetailModalInstance.show();
  }

  const handleTableAlert = (callback) => { 
    alertMessage = callback.detail;
    setTimeout(() => {
      alertMessage = {
        text: '',
        status: '',
      };
    }, 4300);
  }

  const handleFormSave = (event) => {
    tagDataTable.list();
    tagDetailModalInstance.hide();
  };
</script>
<style></style>

<div bind:this={tagDetailModal} class="modal fade" tabindex="-1">
  <div class="modal-dialog modal-lg modal-dialog">
    <div class="modal-content">
      <div class="modal-header">
        <h5 class="modal-title">{modalTitle}</h5>
        <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Cerrar"></button>
      </div>
      <div class="modal-body">
        <TagDetail 
          bind:this={tagFormInstance} 
          on:saved={handleFormSave} />
      </div>
    </div>
  </div>
</div>

<div class="row">
  <h1 class="mb-2 subtitle"><i class="fa fa-tags me-2"></i>Gestión de Etiquetas</h1>
</div>
<hr>
{#if alertMessage.text != ''}
  <div class="alert alert-{alertMessage.status}" role="alert">
    {alertMessage.text}
  </div>
{/if}
<div class="row">
  <TagsFilters 
    on:search={handleSearchFilter} 
    on:clean={handleCleanFilter} />
</div>
<div class="row subtitle-row">
  <h4 class="subtitle">Listado de Etiquetas</h4>
</div>
<div class="row">
  <div class="col-md-7">
    <DataTable 
      bind:this={tagDataTable}
      recordId={'_id'}
      fetchURL={URLS.TICKETS_SERVICE + 'api/v1/tags'}
      deleteURL={URLS.TICKETS_SERVICE + 'api/v1/tags'}
      tokenStorageId={'jwtTicketsToken'}
      columnKeys={['_id', 'name', ]}
      columnTypes={['id', 'td', ]}
      columnNames={['ID', 'Nombre', 'Acciones']}
      columnStyles={['max-width: 50px;', 'max-width: 250px;', ]}
      columnClasses={['d-none',, 'text-end']}
      tdStyles={['max-width: 50px;', 'max-width: 250px;', ]}
      messages = {{
        success: 'Datos actualizados', 
        errorNetwork: 'No se pudo listar las etiquetas. No hay conexión con el servidor.',
        notFound: 'No se pudo listar las etiquetas. Recurso no encontrado.', 
        serverError:'No se pudo listar las etiquetas. Error interno del servidor',
        requestError: 'No se pudo listar las etiquetas. No se recibió respuesta del servidor',
        otherError: 'No se pudo listar las etiquetas. Ocurrió un error no esperado al traer los datos del servidor',
      }}
      addButton={{
        display: true,
        disabled: false,
        action: addTag
      }}
      pagination = {{
        display: true,
        step: 10,
        totalPages: 0,
        actualPage: 1
      }}
      actionButtons={[]} 
      on:alert={handleTableAlert}
      token={localStorage.getItem('jwtTicketsToken')}
    />
  </div>
</div>