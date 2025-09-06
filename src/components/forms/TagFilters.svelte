<svelte:options accessors={true} />
<script>
  import { onMount } from 'svelte';
  import { createEventDispatcher } from 'svelte';

  const dispatch = createEventDispatcher();
  let name = '';
  let description = '';


  onMount(() => {
  });

  const searchClick = (event) => {
    event.preventDefault();
    console.log({name, description});
    dispatch('search', {name, description});
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
  <div class="row subtitle-row mt-3">
    <h4 class="subtitle">Filtros de Búsqueda</h4>
  </div>
  <div class="row align-items-end">
    <div class="col-md-3">
      <label for="name" class="form-label">Buscar por Nombre</label>
      <input type="text" class="form-control" id="name" placeholder="Nombre" bind:value={name}>
    </div>
    <div class="col-md-5 d-flex justify-content-start">
      <button type="submit" class="btn btn-primary me-2" on:click={searchClick}>
        <i class="fa fa-search me-2"></i> Buscar
      </button>
      <button type="reset" class="btn btn-secondary" on:click={cleanClick}>
        <i class="fa fa-eraser me-2"></i> Limpiar
      </button>
    </div>
  </div>
</form>