<script>
  import { onMount } from 'svelte';
  import { navigate } from 'svelte-routing';
  import IssueHeaderTab from '../../forms/IssueHeaderTab.svelte';
  import IssueAssetTab from '../../forms/IssueAssetTab.svelte';
  import { fetchOneById } from '../../../services/app/issues_services';
  import { activeLink } from '../../../helpers/store'; 

  export let title = 'Nueva Incidencia';
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
  
  let modalTitle;

  // Función para manejar el cambio de pestaña
  const handleTabChange = (event) => {
    const tabId = event.target.getAttribute('data-bs-target');
    console.log('Pestaña cambiada a:', tabId);
    
    // Aquí puedes ejecutar lógica específica para cada pestaña
    switch(tabId) {
      case '#header':
        console.log('Pestaña Datos Generales activa');
        break;
      case '#assets':
        console.log('Pestaña Activos Asociados activa');
        // Ejemplo: cargar datos de activos
        loadAssetsTab();
        break;
      case '#users':
        console.log('Pestaña Usuarios activa');
        break;
      case '#histories':
        console.log('Pestaña Historial activa');
        break;
      case '#logs':
        console.log('Pestaña Actividad activa');
        break;
    }
  }

  // Ejemplo de función para cargar datos cuando se cambia a una pestaña específica
  const loadAssetsTab = () => {
    if (assetTabInstance && typeof assetTabInstance.loadData === 'function') {
      assetTabInstance.loadData();
    }
  }

  onMount(() => {
    activeLink.set('/issues');
    document.title = "Nueva Incidencia";

    if (_id != ''){
      document.title = "Editar Incidencia";
      title = 'Editar Incidencia'
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

    // Agregar event listeners para los tabs
    const tabElements = document.querySelectorAll('[data-bs-toggle="tab"]');
    tabElements.forEach(tab => {
      tab.addEventListener('shown.bs.tab', handleTabChange);
    });

    // Limpiar event listeners cuando el componente se destruye
    return () => {
      tabElements.forEach(tab => {
        tab.removeEventListener('shown.bs.tab', handleTabChange);
      });
    };
  });
</script>
<style></style>

<div class="row">
  <h1 class="mb-2 subtitle">
    <a class="return-link" class:active={$activeLink === '/issues'} href="/issues" on:click|preventDefault={() => {navigate('/issues')}}>
      <i class="fa fa-bug me-2"></i>Gestión de Incidencias
    </a> / {title}
  </h1>
</div>
<hr>
{#if alertMessage.text != ''}
  <div class="alert alert-{alertMessage.status}" role="alert">
    {alertMessage.text}
  </div>
{/if}
<ul class="nav nav-tabs mb-3" id="myTabs" role="tablist">
  <li class="nav-item" role="presentation">
    <button class="nav-link active" id="generales-tab" data-bs-toggle="tab" data-bs-target="#header" type="button" role="tab">
      Datos Generales
    </button>
  </li>
  <li class="nav-item" role="presentation">
    <button class="nav-link" id="etiquetas-tab" data-bs-toggle="tab" data-bs-target="#assets" type="button" role="tab">
      Activos Asociados
    </button>
  <li class="nav-item" role="presentation">
    <button class="nav-link" id="documentos-tab" data-bs-toggle="tab" data-bs-target="#users" type="button" role="tab">
      Acceso a Usuarios
    </button>
  </li>
  <li class="nav-item" role="presentation">
    <button class="nav-link" id="documentos-tab" data-bs-toggle="tab" data-bs-target="#histories" type="button" role="tab">
      Seguimiento e Historial
    </button>
  </li>
  <li class="nav-item" role="presentation">
    <button class="nav-link" id="documentos-tab" data-bs-toggle="tab" data-bs-target="#logs" type="button" role="tab">
      Actividad de Usuarios
    </button>
  </li>
</ul>
<div class="tab-content">
  <div class="tab-pane fade show active" id="header" role="tabpanel">
    <IssueHeaderTab bind:this={headerTabInstance} />
  </div>
  <div class="tab-pane fade" id="assets" role="tabpanel">
    <IssueAssetTab bind:this={assetTabInstance} /> 
  </div>
  <div class="tab-pane fade" id="users" role="tabpanel">
    <div class="row subtitle-row mt-3">
      <h4 class="subtitle">Usuarios</h4>
    </div>
  </div>
  <div class="tab-pane fade" id="histories" role="tabpanel">
    <div class="row subtitle-row mt-3">
      <h4 class="subtitle">Historias</h4>
    </div>
  </div>
  <div class="tab-pane fade" id="logs" role="tabpanel">
    <div class="row subtitle-row mt-3">
      <h4 class="subtitle">Actividad de Usuarios</h4>
    </div>
  </div>
</div>