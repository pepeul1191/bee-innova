<svelte:options accessors={true} />
<script>
  import { onMount, createEventDispatcher } from 'svelte';
  import { Modal } from 'bootstrap';
  import axios from 'axios';
  //import { deleteOne } from '../../../services/app/files_service'; // Asegúrate de importar tu servicio

  export let files = [];
  let fileInput = null;
  let responseMessage = '';
  let isLoading = false;

  export let textInput = '';
  export let rowId = '';
  export let fetchURL = '';
  export let deleteURL = null;
  export let baseURL = '';
  export let data = [];
  export let columnKeys = [];
  export let columnTypes = [];
  export let columnNames = [];
  export let columnClasses = [];
  export let columnStyles = [];
  export let afterUploadCallback = null;
  export let saveURL = '';
  export let tokenStorageId = '';
  export let extraData = {};

  // delete confirmation modal
  let deleteConfirmationInstance;
  let deleteConfirmationModal;
  let messageConfirmationModal = {
    text: '',
    status: ''
  };
  let idForDeleting = null;
  let btnDisabledDeleteConfirmation = false;
  // dispatch
  const dispatch = createEventDispatcher();

  onMount(() => {
    // list();
    deleteConfirmationInstance = new Modal(deleteConfirmationModal);
    //console.log(addButton)
    //console.log(pagination)
  });

  const fileChange = (event) => {
    const { name, value } = event.target;

    if (name === 'file') {
      fileInput = event.target.files[0]; // Guardamos el archivo seleccionado
    } else {
      textInput = value; // Guardamos el texto del input
    }
  };

  export const setExtraData = (data) => {
    extraData = data;
  };

  export const addSavedDocument = (doc) => {
    console.log(files);
    files.push(doc);
    console.log(files)
    files = files;
  }

  export const setFiles = (docs) => {
    files = docs;
  };

  const uploadFile = async () => {
    if (!fileInput) {
      alert('Por favor, selecciona un archivo');
      return;
    }

    isLoading = true;
    responseMessage = '';

    const formData = new FormData();
    formData.append('file', fileInput);
    formData.append('name', textInput);

    // Agregar datos adicionales dinámicamente
    console.log(extraData)
    for (const [key, value] of Object.entries(extraData || {})) {
      formData.append(key, value);
    }

    const token = localStorage.getItem(tokenStorageId); // Obtén el JWT

    try {
      console.log(saveURL)
      const res = await fetch(saveURL, {
        method: 'POST',
        headers: {
          Authorization: `Bearer ${token}`
          // No establezcas Content-Type con FormData
        },
        body: formData,
        credentials: 'include'
      });

      if (!res.ok) throw new Error(`Error del servidor: ${res.status}`);
      const data = await res.json();
      //console.log(data)
      // Actualizamos la lista de archivos si recibimos uno nuevo
      if (afterUploadCallback) {
        afterUploadCallback(data);
      }
      responseMessage = `✅ Éxito: ${JSON.stringify(data)}`;
    } catch (err) {
      responseMessage = `❌ Error: ${err.message}`;
    } finally {
      isLoading = false;
    }
  };

  let messageImageAlert = {
    show: false,
    message: '',
    class: ''
  };

  const radioThClicked = (event, key) => {
    const isChecked = event.target.checked;
    //console.log(isChecked);
    //console.log(key);
    data.forEach((record) => {
      record[key] = isChecked;
      let rowKey = record[recordId]
      if(isChecked){
        if(!observer.edit.includes(rowKey)){
          observer.edit.push(rowKey)
        }
      }else{
        if(!observer.edit.includes(rowKey)){
          observer.edit.push(rowKey)
        }
      }
    });
    data = data;
  };

  async function viewFile(fileUrl) {
    const token = localStorage.getItem(tokenStorageId); // Obtén el JWT

    try {
      const response = await axios.get(fileUrl, {
        responseType: 'blob', // Pide una respuesta binaria
        headers: {             // <--- Add the headers object
          Authorization: `Bearer ${token}` // <--- Your authorization header
        }
      });

      const contentType = response.headers['content-type'] || 'application/octet-stream';
      const blob = new Blob([response.data], { type: contentType });
      const url = URL.createObjectURL(blob);

      window.open(url, '_blank');

      // Opcional: Liberar la URL de objeto después de un tiempo
      // setTimeout(() => URL.revokeObjectURL(url), 60000); // Libera después de 1 minuto

    } catch (error) {
      console.error("Error al obtener o mostrar el archivo:", error);
      alert("No se pudo cargar el archivo. Por favor, inténtalo de nuevo.");
    }
  }

  export const askToDeleteRow = (record, key) => {
    idForDeleting = record[key];
    deleteConfirmationInstance.show();
  }

  const deleteCallack = () => {
    dispatch('deleteCallback', { 
      [rowId]: idForDeleting,
    });
    deleteConfirmationInstance.hide();
  }

  export const deleteFromData = (id) => {
    files = files.filter(file => file[rowId] !== id);
  }
</script>

<style>
  .file {
    justify-content: space-between;
    align-items: center;
  }
</style>

<div bind:this={deleteConfirmationModal} class="modal fade" tabindex="-1">
  <div class="modal-dialog">
    <div class="modal-content">
      <div class="modal-header">
        <h5 class="modal-title">Confirmación de Eliminación</h5>
        <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Cerrar"></button>
      </div>
      <div class="modal-body">
        {#if messageConfirmationModal.text != ''}
          <div class="alert alert-{messageConfirmationModal.status}" role="alert">
            {messageConfirmationModal.text}
          </div>
        {/if}
        ¿Seguro que quiere borrar el documento?
      </div>
      <div class="modal-footer">
        <button type="button" class="btn btn-secondary" data-bs-dismiss="modal" disabled={btnDisabledDeleteConfirmation}>
          <i class="fa fa-times"></i>Cancelar</button>
        <button type="button" class="btn btn-danger" disabled={btnDisabledDeleteConfirmation}
        on:click={deleteCallack}>
          <i class="fa fa-trash"></i> Eliminar
        </button>
      </div>
    </div>
  </div>
</div>

<div class="container">
  <div class="mb-12">
    <div class="row mb-3">
      <div class="col-md-4">
        <input type="file" class="form-control" id="formFile" name="file" on:change={fileChange} />
      </div>
      <div class="col-md-6">
        <input
          type="text"
          class="form-control ms-2"
          placeholder="Descripción del archivo"
          id="fileName"
          name="name"
          bind:value={textInput}
          on:input={fileChange}
        />
      </div>
      <div class="col-md-2 d-flex justify-content-end align-items-end">
        <button type="button" class="btn btn-primary ms-3" on:click={uploadFile} disabled={isLoading}>
          <span><i class="fa fa-upload"></i> Subir Archivo</span>
        </button>
      </div>
    </div>
  </div>

  <!-- Mensaje de respuesta -->
  {#if responseMessage}
    <div class="alert alert-info">{responseMessage}</div>
  {/if}

  <div class="mb-12">
    <table class="table table-striped">
      <thead>
        <tr>
          {#each columnNames as key, i}
            {#if columnTypes[i] == 'radiobuttonAll'}
              <th class="{columnClasses[i]}" style="{columnStyles[i]}" scope="col">
                {columnNames[i]} 
                <input
                  class="form-check-input form-check-input-all"
                  type="checkbox"
                  on:change={(event) => radioThClicked(event, columnKeys[i])}
                />
              </th>
            {:else}
              <th class="{columnClasses[i]}" style="{columnStyles[i]}" scope="col">{columnNames[i]}</th>
            {/if}
          {/each}
        </tr>
      </thead>
      <tbody>
        {#each files as file}
          <tr>
            <td>{file.name}</td>
            <td>{file.description}</td>
            <td>{new Date(file.created).toLocaleString()}</td>
            <td class="text-end">
              <a class="btn btn-secondary" href="{baseURL}{file.url}" target="_blank" rel="noopener noreferrer" on:click|preventDefault={viewFile(baseURL + file.url)}>
                <i class="fa fa-search"></i> Ver 
              </a>
              <button class="btn btn-danger" on:click={askToDeleteRow(file, rowId)}>
                <i class="fa fa-trash"></i> Borrar 
              </button>
            </td>
          </tr>
        {/each}
      </tbody>
    </table>
  </div>
</div>