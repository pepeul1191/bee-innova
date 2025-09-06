<script>
  import { onMount } from 'svelte';
  import { navigate } from "svelte-routing";
  import IssueFilters from '../../forms/IssueFilters.svelte';
  import DataTable from '../../widgets/DataTable.svelte';
  import { activeLink } from '../../../helpers/store'; 

  let issueDetailModalInstance;
  let issueFormInstance;
  let issueDetailModal;
  let alertMessage = {
    text: '',
    status: '',
  };
  let issueDataTable;
  let modalTitle;

  onMount(() => {
    activeLink.set('/teachers');
    document.title = "Gestión de Docentes";
    // montar acciones de la tabla
      // ejemplos
      //issueDataTable.addButton.action = () => issueDataTable.addRow();
      //issueDataTable.addButton.action = () => issueDataTable.goToLink('/users');
      //issueDataTable.addButton.action = () => issueDataTable.goToHref(BASE_URL + 'hola');
      //issueDataTable.addButton.action = () => issueDataTable.openTab(BASE_URL + 'hola');
    
    // table action buttons
    issueDataTable.actionButtons = [
      {
        class: 'btn-secondary',
        icon: 'fa-pencil',
        label: 'Editar',
        action: editIssue
      },
      {
        class: 'btn-danger',
        icon: 'fa-trash',
        label: 'Eliminar',
        action: (record) => {
          issueDataTable.askToDeleteRow(record, '_id');
        }
      },
    ];
    issueDataTable.list();
  });

  const editIssue = (issue) => {
    navigate(`/issues/${issue._id}/edit`);
  }

  const handleSearchFilter = (event) => {
    const { name, initDate, endDate, issueStateId, priorityId, selectedTags } = event.detail;
    console.log(event.detail);
    issueDataTable.queryParams = {name, initDate, endDate, issueStateId, priorityId, selectedTags};
    issueDataTable.list();
  }
  
  const handleCleanFilter= () => {
    issueDataTable.queryParams = {};
    issueDataTable.list();
  }

  const addIssue = () => {
    navigate(`/teachers/new`);
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
</script>
<style></style>

<div class="row">
  <h1 class="mb-2 subtitle"><i class="fa fa-graduation-cap me-2"></i>Gestión de Docentes</h1>
</div>
<hr>
{#if alertMessage.text != ''}
  <div class="alert alert-{alertMessage.status}" role="alert">
    {alertMessage.text}
  </div>
{/if}
<div class="row">
  <IssueFilters 
    on:search={handleSearchFilter} 
    on:clean={handleCleanFilter} />
</div>
<div class="row subtitle-row">
  <h4 class="subtitle">Listado de Docentes</h4>
</div>
<div class="row">
  <DataTable 
    bind:this={issueDataTable}
    fetchURL={URLS.TICKETS_SERVICE + 'api/v1/issues/summary'}
    deleteURL={URLS.TICKETS_SERVICE + 'api/v1/issues'}
    tokenStorageId={'jwtClassroomToken'}
    columnKeys={['_id', 'resume', 'reporter.name', 'reportered', 'priority.name', 'issue_state.name']}
    columnTypes={['id', 'td', 'td']}
    columnNames={['ID', 'Resumen', 'Reportado por', 'Fecha', 'Prioridad', 'Estado', 'Acciones']}
    columnStyles={['max-width: 50px;', 'max-width: 250px;', 'max-width: 400px;', 'max-width: 150px;']}
    columnClasses={['d-none', '', '', '','','','text-end']}
    tdStyles={['max-width: 50px;', 'max-width: 250px;', 'max-width: 400px;', 'max-width: 150px;']}
    messages = {{
      success: 'Datos actualizados', 
      errorNetwork: 'No se pudo listar los docentes. No hay conexión con el servidor.',
      notFound: 'No se pudo listar los docentes. Recurso no encontrado.',
      serverError:'No se pudo listar los docentes. Error interno del servidor',
      requestError: 'No se pudo listar los docentes. No se recibió respuesta del servidor',
      otherError: 'No se pudo listar los docentes. Ocurrió un error no esperado al traer los datos del servidor',
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
    recordId={'_id'}
    actionButtons={[]} 
    on:alert={handleTableAlert}
    token={localStorage.getItem('jwtClassroomToken')}
  />
</div>