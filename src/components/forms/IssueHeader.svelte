<svelte:options accessors={true} />
<script>
  import { onMount } from 'svelte';
  import { createEventDispatcher } from 'svelte';
  import { fetchAll as fetchAllPriorities } from '../../services/app/priorities_service';
  import { fetchAll as fetchAllIssueStates } from '../../services/app/issue_states_service';

  const dispatch = createEventDispatcher();
  let name = '';
  let description = '';
  let initDate = '';
  let endDate = '';
  let issueStateId = '';
  let moment = '';
  let priorityId = '';
  let priorities = [];
  let issueStates = [];
  let reporteredBy = '';

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
      });;
  });

  const searchClick = (event) => {
    event.preventDefault();
    console.log({name, description, initDate, endDate, issueStateId, priorityId});
    dispatch('search', {name, description, initDate, endDate, issueStateId, priorityId});
  }

  const cleanClick = (event) => {
    event.preventDefault();
    if (name || description) {
      // Si al menos uno de los campos tiene valor, limpiamos ambos y disparamos el evento
      name = '';
      description = '';
      dispatch('clean');
    }
  }
</script>

<style>

</style>

<form class="mb-4">
  <div class="row">
    <div class="col-md-9">
      <label for="name" class="form-label">Resumen</label>
      <input type="text" class="form-control" id="name" placeholder="Nombre" bind:value={name}>
    </div>
    <div class="col-md-3">
      <label for="reportered" class="form-label">Reportado por</label>
      <input type="text" disabled class="form-control" id="reportered" placeholder="Descripción" bind:value={reporteredBy}>
    </div>
  </div>
  <div class="row mt-3">
    <div class="col-md-12">
      <label for="description" class="form-label">Descripción</label>
      <textarea
        class="form-control"
        id="description"
        placeholder="Descripción"
        rows="8">
      </textarea>
    </div>
  </div>
  <div class="row mt-3">
    <div class="col-md-3">
      <label for="status" class="form-label">Estado de Ticket</label>
      <select id="issueStateId" class="form-select" bind:value={issueStateId}>
        <option value="" disabled selected>Seleccione Estado</option>
        {#each issueStates as state}
          <option value="{state._id}">{state.name}</option>
        {/each}
      </select>
    </div>
    <div class="col-md-3">
      <label for="status" class="form-label">Prioridad</label>
      <select id="priorityId" class="form-select" bind:value={priorityId}>
        <option value="" disabled selected>Seleccione prioridad</option>
        {#each priorities as priority}
          <option value="{priority._id}">{priority.name}</option>
        {/each}
      </select>
    </div>
    <div class="col-md-3">
      <label for="status" class="form-label">Momento</label>
      <input
        type="datetime-local"
        class="form-control"
        min="2024-01-01T00:00"
        max="2025-12-31T23:59"
        bind:value={moment}
      />
    </div>
    <div class="col-md-3 d-flex justify-content-end align-items-end">
      <button type="submit" class="btn btn-success me-2" on:click={searchClick}>
        <i class="fa fa-floppy-o me-2"></i> Guardar Datos Generales
      </button>
    </div>
  </div>
</form>