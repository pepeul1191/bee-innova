<svelte:options accessors={true} />
<script>
  import { onMount } from 'svelte';
  import EasyMDE from 'easymde';
  import { createEventDispatcher } from 'svelte';
  import { fetchAll as fetchAllTags } from '../../services/app/tags_service';
  import { fetchAll as fetchAllPriorities } from '../../services/app/priorities_service';
  import { fetchAll as fetchAllIssueStates } from '../../services/app/issue_states_service';
  import { create as createIssue, patchTags, editIssue, addDocument, removeDocument } from '../../services/app/issues_services';
  import { localDateTimeToISOString } from '../../helpers/datetime';
  import { getInfo } from '../../services/app/auth_services';
	import FileTable from '../widgets/FileTable.svelte';

  const dispatch = createEventDispatcher();
  export let issue = {}
  let issueId = '';
  let resume = '';
  let description = '';
  let issueStateId = '';
  let moment = '';
  let priorityId = '';
  let priorities = [];
  let issueStates = [];
  let selectedTags = [];
  let tags = [];
  let documents;
  let employee = {
    _id: '',
    names: '',
    last_names: '',
  }
  let employeeFullName = '';
  let disabled = true;
  let editorElement;
  let easymde;

  export const updateView = () => {
    issueId = issue._id;
    resume = issue.resume;
    description = issue.description;
    employeeFullName = issue.reporter.name; 
    issueStateId = issue.issue_state._id;
    priorityId = issue.priority._id;
    moment = issue.reportered;
    selectedTags = issue.tags.map(tag => tag._id);
    disabled = false;
    // documentos
    documents.setExtraData({
      folder: issue
    });
    documents.setFiles(issue.documents);
  }

  export const loadMarkDownEditor = () => {
    easymde = new EasyMDE({
      element: editorElement,
      initialValue: description,
      spellChecker: false,
      autosave: {
        enabled: false,
      }
    });

    easymde.codemirror.on('change', () => {
      description = easymde.value();
    });
  }

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

    getInfo(URLS.TICKETS_SERVICE, 'jwtTicketsToken')
      .then(response => {
        //console.log('Estados:', response.data);
        employee = response.data.employee;
        employeeFullName = `${employee.names} ${employee.last_names}`;
        //user = response.data.user;
      })
      .catch(error => {
        console.error('Error:', error);
      });
  });

  const saveForm = (event) => {
    event.preventDefault();
    const body = {
      resume: resume,
      description: description, 
      issue_state_id: issueStateId, 
      priority_id: priorityId, 
      reporter_id: employee._id,
      reportered: localDateTimeToISOString(moment)
    }
    if (issueId == ''){
      createIssue(URLS.TICKETS_SERVICE, 'jwtTicketsToken', body)
        .then(response => {
          //console.log('Estados:', response.data);
          console.log(response.data);
          issueId = response.data._id;
          disabled = false;
          //user = response.data.user;
        })
        .catch(error => {
          console.error('Error:', error);
        });
    }else{
      editIssue(URLS.TICKETS_SERVICE, 'jwtTicketsToken', body, issueId)
        .then(response => {
          //console.log('Estados:', response.data);
          console.log(response.data);
          issueId = response.data._id; 
          //user = response.data.user;
        })
        .catch(error => {
          console.error('Error:', error);
        });
    }
    //dispatch('search', {name, description, initDate, endDate, issueStateId, priorityId});
  }

  const saveTags = (event) => {
    event.preventDefault();
    patchTags(URLS.TICKETS_SERVICE, 'jwtTicketsToken', issueId, selectedTags)
      .then(response => {
        //console.log('Estados:', response.data);
        console.log(response.data);
        //user = response.data.user;
      })
      .catch(error => {
        console.error('Error:', error);
      });
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

  const addDocumentToIssue = (file) => {
    //console.log(file)
    console.log(documents)
    let data = {
      name: file.original_filename,
      description: documents.textInput,
      size: file.size,
      url: file.path,
      mime: file.mime_type,
    }

    addDocument(URLS.TICKETS_SERVICE, 'jwtTicketsToken', issueId, data)
      .then(response => {
        //console.log('Estados:', response.data);
        //user = response.data.user;
        let document = {};
        document.created = response.data.created;
        document.description = response.data.description;
        document.name = response.data.name;
        document._id = response.data._id;
        document.url = response.data.url;
        documents.addSavedDocument(document);
      })
      .catch(error => {
        console.error('Error:', error);
      });
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

  const deleteDocumentCallback = (callback) => {
    let documentId = callback.detail._id;
    removeDocument(URLS.TICKETS_SERVICE, 'jwtTicketsToken', issueId, documentId)
      .then(response => {
          //console.log('Estados:', response.data);
          console.log(response.data);
          //user = response.data.user;
          documents.deleteFromData(documentId);
        })
        .catch(error => {
          console.error('Error:', error);
        });
  }
</script>

<style>

</style>

<div class="row subtitle-row">
  <h4 class="subtitle">Datos de la Incidencia</h4>
</div>
<div class="row">
  <form class="mb-4">
    <div class="row">
      <div class="col-md-9">
        <label for="resume" class="form-label">Resumen</label>
        <input type="text" class="form-control" id="resume" placeholder="Resumen de la incidencia" bind:value={resume}>
      </div>
      <div class="col-md-3">
        <label for="reportered" class="form-label">Reportado por</label>
        <input type="text" disabled class="form-control" id="reportered" placeholder="Nombre" bind:value={employeeFullName}>
      </div>
    </div>
    <div class="row mt-3">
      <div class="col-md-12">
        <label for="descriptionTxt" class="form-label">Descripción</label>
        <textarea class="form-control" id="descriptionTxt" bind:value={description} bind:this={editorElement} rows="8" placeholder="Descripción"></textarea>
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
        <button type="submit" class="btn btn-success me-2" on:click={saveForm}>
          <i class="fa fa-floppy-o me-2"></i> Guardar Datos Generales
        </button>
      </div>
    </div>
  </form>
</div>
<div class="row subtitle-row">
  <h4 class="subtitle">Etiquetas</h4>
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
          {disabled}
          id={"tag-" + tag._id}
        />
        <label class="form-check-label" for={"tag-" + tag._id}>
          {tag.name}
        </label>
      </div>
    </div>
  {/each}
  <div class="col-md-12 d-flex justify-content-end align-items-end">
    <button type="submit" class="btn btn-success me-2" on:click={saveTags} {disabled}>
      <i class="fa fa-floppy-o me-2"></i> Guardar Etiquetas
    </button>
  </div>
</div>
<div class="row subtitle-row mt-3">
  <h4 class="subtitle mb-3">Documentos</h4>
  <FileTable 
    bind:this={documents}
    columnKeys={['_id', 'name', 'description', 'created']}
    columnTypes={['id', 'td', 'td', 'td']}
    columnNames={['ID', 'Nombre', 'Descripción', 'Fecha', 'Acciones']}
    columnStyles={['max-width: 50px;', 'max-width: 250px;', 'max-width: 400px;', 'max-width: 150px;', 'text-align: right !important; padding-right: 70px;']}
    columnClasses={['d-none', '', '', '','','text-end']}
    tdStyles={['max-width: 50px;', 'max-width: 250px;', 'max-width: 400px;', 'max-width: 150px;']}
    messages = {{
      success: 'Datos actualizados', 
      errorNetwork: 'No se pudo listar los sistemas. No hay conexión con el servidor.',
      notFound: 'No se pudo listar los sistemas. Recurso no encontrado.',
      serverError:'No se pudo listar los sistemas. Error interno del servidor',
      requestError: 'No se pudo listar los sistemas. No se recibió respuesta del servidor',
      otherError: 'No se pudo listar los sistemas. Ocurrió un error no esperado al traer los datos del servidor',
    }}
    baseURL={URLS.FILES_SERVICE + 'api/v1/files/'}
    saveURL={URLS.FILES_SERVICE + 'api/v1/files'}
    fetchURL={URLS.TICKETS_SERVICE + 'api/v1/files'}
    tokenStorageId={'jwtFilesToken'}
    afterUploadCallback={addDocumentToIssue}
    on:deleteCallback={deleteDocumentCallback}
    rowId={'_id'}
    extraData={{
      folder: issueId,
    }} />
</div>
