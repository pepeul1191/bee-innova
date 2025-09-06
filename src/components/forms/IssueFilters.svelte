<svelte:options accessors={true} />
<script>
  import { onMount } from 'svelte';
  import { createEventDispatcher } from 'svelte';
  import { fetchAll as fetchAllPriorities } from '../../services/app/priorities_service';
  import { fetchAll as fetchAllIssueStates } from '../../services/app/issue_states_service';
  import { fetchAll as fetchAllTags } from '../../services/app/tags_service';

  const dispatch = createEventDispatcher();
  let name = '';
  let description = '';
  let initDate = '';
  let endDate = '';
  let issueStateId = '';
  let priorityId = '';
  let priorities = [];
  let issueStates = [];
  let selectedTags = [];
  let tags = [];

  onMount(() => {
    fetchAllPriorities(URLS.TICKETS_SERVICE, 'jwtTicketsToken')
      .then(response => {
        //console.log('Prioridades:', response.data);
        priorities = response.data;
      })
      .catch(error => {
        console.error('Error:', error);
      });

    fetchAllIssueStates(URLS.TICKETS_SERVICE, 'jwtTicketsToken')
      .then(response => {
        //console.log('Estados:', response.data);
        issueStates = response.data;
      })
      .catch(error => {
        console.error('Error:', error);
      });
      
    fetchAllTags(URLS.TICKETS_SERVICE, 'jwtTicketsToken')
      .then(response => {
        //console.log('Estados:', response.data);
        tags = response.data;
      })
      .catch(error => {
        console.error('Error:', error);
      });
  });

  const searchClick = (event) => {
    event.preventDefault();
    //console.log({name, initDate, endDate, issueStateId, priorityId});
    //console.log(selectedTags)
    dispatch('search', {name, initDate, endDate, issueStateId, priorityId, selectedTags});
  }

  const cleanClick = (event) => {
    event.preventDefault();
    if (name || initDate || endDate || issueStateId || priorityId || selectedTags) {
      // Si al menos uno de los campos tiene valor, limpiamos ambos y disparamos el evento
      name = '';
      initDate = '';
      endDate = '';
      issueStateId = '';
      priorityId = '';
      selectedTags = [];
      dispatch('clean');
    }
  }
</script>

<style>

</style>

<form class="mb-4">
  <div class="row subtitle-row mt-3">
    <h4 class="subtitle">Filtros de Búsqueda</h4>
  </div>
  <div class="row">
    <div class="col-md-4">
      <label for="name" class="form-label">Buscar por Nombre</label>
      <input type="text" class="form-control" id="name" placeholder="Nombre" bind:value={name}>
    </div>
    <div class="col-md-2">
      <label for="status" class="form-label">Estado de Ticket</label>
      <select id="issueStateId" class="form-select" bind:value={issueStateId}>
        <option value="" disabled selected>Seleccione Estado</option>
        {#each issueStates as state}
          <option value="{state._id}">{state.name}</option>
        {/each}
      </select>
    </div>
    <div class="col-md-2">
      <label for="status" class="form-label">Prioridad</label>
      <select id="priorityId" class="form-select" bind:value={priorityId}>
        <option value="" disabled selected>Seleccione prioridad</option>
        {#each priorities as priority}
          <option value="{priority._id}">{priority.name}</option>
        {/each}
      </select>
    </div>
    <div class="col-md-2">
      <label for="status" class="form-label">Fecha Creación - Inicio</label>
      <input type="date" class="form-control" min="2024-01-01" max="2025-12-31" bind:value={initDate}>
    </div>
    <div class="col-md-2">
      <label for="status" class="form-label">Fecha Creación - Fin</label>
      <input type="date" class="form-control" min="2024-01-01" max="2025-12-31" bind:value={endDate}>
    </div>
  </div>
  <div class="row subtitle-row mt-3">
    <p>Etiquetas de Incidencias</p>
  </div>
  <div class="row">
    {#each tags as tag}
      <div class="col-md-2 mb-2">
        <div class="form-check">
          <input
            class="form-check-input"
            type="checkbox"
            bind:group={selectedTags}
            value={tag._id}
            id={"tag-" + tag._id}
          />
          <label class="form-check-label" for={"tag-" + tag._id}>
            {tag.name}
          </label>
        </div>
      </div>
    {/each}
  </div>
  
  <div class="row mt-3">
    <div class="col-md-12 d-flex align-items-end justify-content-end">
      <button type="submit" class="btn btn-primary me-2" on:click={searchClick}>
        <i class="fa fa-search me-2"></i> Buscar
      </button>
      <button type="reset" class="btn btn-secondary" on:click={cleanClick}>
        <i class="fa fa-eraser me-2"></i> Limpiar
      </button>
    </div>
  </div>
</form>