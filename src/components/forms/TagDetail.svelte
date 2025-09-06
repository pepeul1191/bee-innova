<svelte:options accessors={true} />
<script>
  import axios from 'axios';
  import { createEventDispatcher } from 'svelte';
  import { toDatetimeLocalWithSeconds } from '../../helpers/datetime.js';
  
  const dispatch = createEventDispatcher();

  const jwtToken = localStorage.getItem('jwtTicketsToken');

  export let id = null;
  let name = '';
  let created = '';
  let updated = '';
  let btnDisabled = false;
  let message = {
    text: '',
    status: ''
  };

  export const clean = () => {
    id = null;
    name = '';
    created = '';
    updated = '';
    btnDisabled = false;
    message = {
      text: '',
      status: ''
    };
  }

  export const loadTag = (tag) => {
    id = tag._id;
    name = tag.name;
    created = toDatetimeLocalWithSeconds(tag.created);
    updated = toDatetimeLocalWithSeconds(tag.updated);
  }  

  const cleanMessage = (dispatchToParent) => {
    setTimeout(() => {
      message = {
        text: '',
        status: ''
      };
      btnDisabled = false;
      if (dispatchToParent){
        dispatch('saved', { id });
      }
    }, 4300);
  }

  const handleSubmit = async (event) => {
    event.preventDefault();  // Evita el comportamiento por defecto del formulario (recargar la página)
    // Crea un objeto con los datos del formulario
    try {
      // Enviar datos usando Axios
      const formData = {
        id,
        name,
      };
      let response;
      btnDisabled = true;
      if(id == null){
        response = await axios.post(URLS.TICKETS_SERVICE + 'api/v1/tags', formData, {
          headers: {
            'Content-Type': 'application/json',
            'Authorization': `Bearer ${jwtToken}`
          },
        });
        created = toDatetimeLocalWithSeconds(response.data.created);
        updated = toDatetimeLocalWithSeconds(response.data.updated);
        id = response.data._id;
        message.text = 'Se ha creado la etiqueta';
        message.status = 'success';
        // notificar al padre que se ha actualizado algo
        cleanMessage(true);
      }else{
        response = await axios.put(URLS.TICKETS_SERVICE + 'api/v1/tags/' + id, formData, {
          headers: {
            'Authorization': `Bearer ${jwtToken}`,
            'Content-Type': 'application/json',
          },
        });  
        updated = toDatetimeLocalWithSeconds(response.data.updated);
        message.text = 'Se ha editado la etiqueta';
        message.status = 'success';
        cleanMessage(true);
      }
      console.log('Datos enviados con éxito:', response.data);
      // Puedes manejar la respuesta aquí, por ejemplo, mostrar un mensaje de éxito
    } catch (error) {
      console.error('Error al enviar los datos:', error);
      message.text = 'Error al grabar el activo';
      message.status = 'danger';
      cleanMessage(false);
      // Maneja el error (puedes mostrar un mensaje de error en la interfaz de usuario)
    }
  };
</script>

<style></style>

<form on:submit|preventDefault={handleSubmit}>
  {#if message.text != ''}
    <div class="alert alert-{message.status}" role="alert">
      {message.text}
    </div>
  {/if}
  <div class="mb-3">
    <label for="name" class="form-label">Nombre</label>
    <input type="text" class="form-control" id="name" name="name" bind:value={name} required maxlength="40" placeholder="Ingrese el nombre del activo">
  </div>
  <!-- Fila para las fechas -->
  <div class="row mb-3">
    <div class="col-6">
      <label for="created" class="form-label">Fecha de Creación</label>
      <input type="datetime-local" class="form-control" id="created" name="created" bind:value={created} required disabled>
    </div>
    <div class="col-6">
      <label for="updated" class="form-label">Fecha de Actualización</label>
      <input type="datetime-local" class="form-control" id="updated" name="updated" bind:value={updated} required disabled>
    </div>
  </div>
  <div class="text-end">
    <button type="submit" class="btn btn-success mt-1"  disabled={btnDisabled}>
      <i class="fa fa-check"></i>Guardar Activo
    </button>
  </div>
</form>
  