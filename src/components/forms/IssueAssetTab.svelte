<script>
  import { onMount } from 'svelte';
  import { useLocation } from "svelte-routing";
  import { Modal } from 'bootstrap';
  import DataTable from '../widgets/DataTable.svelte';
	import AssetFilters from './AssetFilters.svelte';
  
  let alertMessage = {
    text: '',
    status: '',
  };
  export let id = null;
  let assetsDataTable;
  let name = '';

  const handleSearchFilter = (event) => {
    const { username, email, status } = event.detail;
    assetsDataTable.queryParams = {username,email, status};
    assetsDataTable.list();
  }
  
  const handleCleanFilter= () => {
    assetsDataTable.queryParams = {};
    assetsDataTable.list();
  }

  onMount(() => {
    // montar acciones de la tabla
      // ejemplos
      //assetsDataTable.addButton.action = () => assetsDataTable.addRow();
      //assetsDataTable.addButton.action = () => assetsDataTable.goToLink('/users');
      //assetsDataTable.addButton.action = () => assetsDataTable.goToHref(BASE_URL + 'hola');
      //assetsDataTable.addButton.action = () => assetsDataTable.openTab(BASE_URL + 'hola');
    // table action buttons
  });

  export const loadData = () => {
    assetsDataTable.list();
    const queryString = window.location.search;
    if (queryString) {
      const params = new URLSearchParams(queryString);
      name = params.get("name");
    }
  }
</script>

<style>

</style>

<div class="row subtitle-row">
  <h4 class="subtitle">Asociar Activos de la Incidencia</h4>
</div>
{#if alertMessage.text != ''}
  <div class="alert alert-{alertMessage.status}" role="alert">
    {alertMessage.text}
  </div>
{/if}
<div class="row">
  <AssetFilters 
    on:search={handleSearchFilter} 
    on:clean={handleCleanFilter} 
    showTitle={false}/>
</div>
<div class="row">
  <div class="col-md-12">
    <DataTable 
      bind:this={assetsDataTable}
      fetchURL={URLS.TICKETS_SERVICE + 'api/v1/systems/' + id + '/users'}
      columnKeys={['id', 'username', 'email', 'registered']}
      columnTypes={['id', 'td', 'td', 'radiobuttonAll']}
      columnNames={['ID', 'Nombre', 'Descripción', 'Registrado', 'Acciones']}
      columnStyles={['max-width: 50px;', 'max-width: 250px;', 'max-width: 150px;', '']}
      tdStyles={['', '', '', 'padding-left: 40px;']}
      columnClasses={['d-none', '', '', '', 'text-end']}
      messages = {{
        success: 'Datos actualizados', 
        errorNetwork: 'No se pudo listar los usuarios. No hay conexión con el servidor.',
        notFound: 'No se pudo listar los usuarios. Recurso no encontrado.',
        serverError:'No se pudo listar los usuarios. Error interno del servidor',
        requestError: 'No se pudo listar los usuarios. No se recibió respuesta del servidor',
        otherError: 'No se pudo listar los usuarios. Ocurrió un error no esperado al traer los datos del servidor',
      }}
      saveButton={{
        display: true,
        disabled: false,
      }}
      pagination = {{
        display: true,
        step: 10,
        totalPages: 0,
        actualPage: 1
      }}
      actionButtons={[]} 
      jwtToken={localStorage.getItem('jwtToken')}
    />
  </div>
</div>
