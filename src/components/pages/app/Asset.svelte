<script>
  import { onMount } from 'svelte';
  import { Modal } from 'bootstrap';
  import AssetsFilters from '../../forms/AssetFilters.svelte';
  import DataTable from '../../widgets/DataTable.svelte';
  import AssetDetail from '../../forms/AssetDetail.svelte';
  import { activeLink } from '../../../helpers/store'; 

  let assetDetailModalInstance;
  let assetFormInstance;
  let assetDetailModal;
  let alertMessage = {
    text: '',
    status: '',
  };
  let assetDataTable;
  let modalTitle;

  onMount(() => {
    activeLink.set('/assets');
    document.title = "Gestión de Activos";
    // montar acciones de la tabla
      // ejemplos
      //assetDataTable.addButton.action = () => assetDataTable.addRow();
      //assetDataTable.addButton.action = () => assetDataTable.goToLink('/users');
      //assetDataTable.addButton.action = () => assetDataTable.goToHref(BASE_URL + 'hola');
      //assetDataTable.addButton.action = () => assetDataTable.openTab(BASE_URL + 'hola');
    assetDetailModalInstance = new Modal(assetDetailModal);
    // table action buttons
    assetDataTable.actionButtons = [
      {
        class: 'btn-secondary',
        icon: 'fa-pencil',
        label: 'Editar',
        action: editAsset
      },
      {
        class: 'btn-danger',
        icon: 'fa-trash',
        label: 'Eliminar',
        action: (record) => {
          assetDataTable.askToDeleteRow(record, '_id');
        }
      },
    ];
    assetDataTable.list();
  });

  const editAsset = (asset) => {
    modalTitle = 'Editar Activo'
    assetFormInstance.clean();
    assetFormInstance.loadAsset(asset);
    assetDetailModalInstance.show();
  }

  const handleSearchFilter = (event) => {
    const { name, description } = event.detail;
    console.log(event.detail);
    assetDataTable.queryParams = {name,description};
    assetDataTable.list();
  }
  
  const handleCleanFilter= () => {
    assetDataTable.queryParams = {};
    assetDataTable.list();
  }

  const addIssue = () => {
    modalTitle = 'Agregar Activo'
    assetFormInstance.clean();
    assetDetailModalInstance.show();
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
    assetDataTable.list();
    assetDetailModalInstance.hide();
  };
</script>
<style></style>

<div bind:this={assetDetailModal} class="modal fade" tabindex="-1">
  <div class="modal-dialog modal-lg modal-dialog">
    <div class="modal-content">
      <div class="modal-header">
        <h5 class="modal-title">{modalTitle}</h5>
        <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Cerrar"></button>
      </div>
      <div class="modal-body">
        <AssetDetail 
          bind:this={assetFormInstance} 
          on:saved={handleFormSave} />
      </div>
    </div>
  </div>
</div>

<div class="row">
  <h1 class="mb-2 subtitle"><i class="fa fa-cogs me-2"></i>Gestión de Activos</h1>
</div>
<hr>
{#if alertMessage.text != ''}
  <div class="alert alert-{alertMessage.status}" role="alert">
    {alertMessage.text}
  </div>
{/if}
<div class="row">
  <AssetsFilters 
    on:search={handleSearchFilter} 
    on:clean={handleCleanFilter} />
</div>
<div class="row subtitle-row">
  <h4 class="subtitle">Listado de Activos</h4>
</div>
<div class="row">
  <DataTable 
    bind:this={assetDataTable}
    recordId={'_id'}
    fetchURL={URLS.TICKETS_SERVICE + 'api/v1/assets'}
    deleteURL={URLS.TICKETS_SERVICE + 'api/v1/assets'}
    tokenStorageId={'jwtTicketsToken'}
    columnKeys={['_id', 'name', 'description']}
    columnTypes={['id', 'td', 'td']}
    columnNames={['ID', 'Nombre', 'Descripción', 'Acciones']}
    columnStyles={['max-width: 50px;', 'max-width: 250px;', 'max-width: 150px;']}
    columnClasses={['d-none', '', '', 'text-end']}
    tdStyles={['max-width: 50px;', 'max-width: 250px;', 'max-width: 150px;']}
    messages = {{
      success: 'Datos actualizados', 
      errorNetwork: 'No se pudo listar los sistemas. No hay conexión con el servidor.',
      notFound: 'No se pudo listar los sistemas. Recurso no encontrado.', 
      serverError:'No se pudo listar los sistemas. Error interno del servidor',
      requestError: 'No se pudo listar los sistemas. No se recibió respuesta del servidor',
      otherError: 'No se pudo listar los sistemas. Ocurrió un error no esperado al traer los datos del servidor',
    }}
    addButton={{
      display: true,
      disabled: false,
      action: addIssue
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