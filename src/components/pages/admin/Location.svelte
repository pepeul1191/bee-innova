<script>
  import { onMount } from 'svelte';
  import DataTable from '../../widgets/DataTable.svelte';

  let department = '';

  let alertMessage = {
    text: '',
    status: ''
  };

  let departmentId;
  let provinceId;

  let departament;
  let province;

  let departmentDataTable;
  let provincesDataTable;
  let districtsDataTable;
  let displayProvinces = false;
  let displayDistricts = false;

  const addRole = () => {
    departmentDataTable.addRow();
  }

  const handleTableAlert = (callback) => { 
    alertMessage = callback.detail;
    console.log(callback.detail)
    setTimeout(() => {
      alertMessage = {
        text: '',
        status: '',
      };
    }, 4300);
  }

  onMount(() => {
    departmentDataTable.list();
    // montar acciones de la tabla
      // ejemplos
    departmentDataTable.addButton.action = () => departmentDataTable.addRow();
    //systemDataTable.addButton.action = () => systemDataTable.goToLink('/users');
    //systemDataTable.addButton.action = () => systemDataTable.goToHref(URLS.CLASSROOM_SERVICE + 'hola');
    //systemDataTable.addButton.action = () => systemDataTable.openTab(URLS.CLASSROOM_SERVICE + 'hola');
    // table action buttons
    departmentDataTable.actionButtons = [
      {
        class: 'btn-secondary', 
        icon: 'fa-list',
        label: 'Provincias',
        action: (record) => {
          //systemDataTable.askToDeleteRow(record, 'id');
          department = record.name;
          provincesDataTable.fetchURL = URLS.CLASSROOM_SERVICE + `api/v1/locations/departments/${record._id}/provinces`;
          provincesDataTable.saveURL = URLS.CLASSROOM_SERVICE + `api/v1/locations/departments/${record._id}/provinces`;
          provincesDataTable.list();
          displayProvinces = true;
          displayDistricts = false;
        }
      },
      {
        class: 'btn-danger',
        icon: 'fa-trash',
        label: 'Eliminar',
        action: (record) => {
          //systemDataTable.askToDeleteRow(record, 'id');
          //console.log(record);
          departmentDataTable.deleteRow(record, '_id');
          if(record.name == departament){
            displayProvinces = false;
            displayDistricts = false;
          }
        }
      },
    ];

    provincesDataTable.addButton.action = () => provincesDataTable.addRow();
    provincesDataTable.actionButtons = [
      {
        class: 'btn-secondary',
        icon: 'fa-list',
        label: 'Distritos',
        action: (record) => {
          //systemDataTable.askToDeleteRow(record, 'id');
          province = record.name;
          districtsDataTable.fetchURL = URLS.CLASSROOM_SERVICE + `api/v1/locations/provinces/${record._id}/districts`;
          districtsDataTable.saveURL = URLS.CLASSROOM_SERVICE + `api/v1/locations/provinces/${record._id}/districts`;
          districtsDataTable.list();
          console.log(districtsDataTable.data.length)
          displayDistricts = true;
        }
      },
      {
        class: 'btn-danger',
        icon: 'fa-trash',
        label: 'Eliminar',
        action: (record) => {
          //systemDataTable.askToDeleteRow(record, 'id');
          //console.log(record);
          provincesDataTable.deleteRow(record, '_id');
          if(record.name == province){
            displayDistricts = false;
          }
        }
      },
    ];

    districtsDataTable.addButton.action = () => districtsDataTable.addRow();
    districtsDataTable.actionButtons = [
      {
        class: 'btn-danger',
        icon: 'fa-trash',
        label: 'Eliminar',
        action: (record) => {
          //systemDataTable.askToDeleteRow(record, 'id');
          //console.log(record);
          districtsDataTable.deleteRow(record, '_id');
        }
      },
    ];
  });
</script>

<style></style>

<div class="row">
  <h1 class="mb-2 subtitle"><i class="fa fa-map me-2"></i>Gestión de Locaciones (Departamentos, Provincias y Distritos)</h1>
</div>
<hr>

{#if alertMessage.text != ''}
  <div class="alert alert-{alertMessage.status}" role="alert">
    {alertMessage.text}
  </div>
{/if}

<div class="card mb-4">
  <div class="card-header bg-secondary text-white">
    <i class="fa fa-info-circle"></i> Información de Locaciones
  </div>
  <div class="card-body">
    <div class="row g-2">
      <div class="col-md-4">
        <div class="row subtitle-row">
          <h4 class="subtitle">Lista de Departamentos</h4>
        </div>
        <DataTable 
          recordId={'_id'}
          bind:this={departmentDataTable}
          fetchURL={URLS.CLASSROOM_SERVICE + `api/v1/locations/departments`}
          saveURL={URLS.CLASSROOM_SERVICE + 'api/v1/locations/departments'}
          columnKeys={['_id', 'name', ]}
          columnTypes={['id', 'input[text]', ]}
          columnNames={['ID', 'Nombre', 'Acciones']}
          columnStyles={['max-width: 50px;', 'width: 70px;', 'width: 240px;']}
          columnClasses={['d-none', '', 'text-end']}
          messages = {{
            success: 'Datos actualizados', 
            errorNetwork: 'No se pudo listar los departamentos. No hay conexión con el servidor.',
            notFound: 'No se pudo listar los departamentos. Recurso no encontrado.',
            serverError:'No se pudo listar los departamentos. Error interno del servidor',
            requestError: 'No se pudo listar los departamentos. No se recibió respuesta del servidor',
            otherError: 'No se pudo listar los departamentos. Ocurrió un error no esperado al traer los datos del servidor',
          }}
          addButton={{
            display: true,
            disabled: false,
            action: null
          }}
          saveButton={{
            display: true,
            disabled: false,
            action: null
          }}
          pagination = {{
            display: false,
            step: 10,
            totalPages: 0,
            actualPage: 1
          }}
          actionButtons={[]} 
          on:alert={handleTableAlert}
          jwtToken={localStorage.getItem('jwtClassroomToken')}
        />
      </div>
      <div class="col-md-4 {!displayProvinces ? 'd-none' : ''}" style="padding-left: 30px; padding-right: 20px;">
        <div class="row subtitle-row">
          <h4 class="subtitle">Provincias del Departamento de '{department}'</h4>
          <DataTable 
            recordId={'_id'}
            bind:this={provincesDataTable}
            fetchURL={URLS.CLASSROOM_SERVICE + `api/v1/locations/departments/${departmentId}/provinces`}
            saveURL={URLS.CLASSROOM_SERVICE + `api/v1/locations/departments/${departmentId}/provinces`}
            columnKeys={['_id', 'name', ]}
            columnTypes={['id', 'input[text]', ]}
            columnNames={['ID', 'Nombre', 'Acciones']}
            columnStyles={['max-width: 50px;', 'max-width: 150px;', 'max-width: 240px;']}
            columnClasses={['d-none', '', 'text-end']}
            messages = {{
              success: 'Datos actualizados', 
              errorNetwork: `No se pudo listar las provincias de departamento ${departament}. No hay conexión con el servidor.`,
              notFound: `No se pudo listar las provincias de departamento ${departament}. Recurso no encontrado.`,
              serverError:`No se pudo listar las provincias de departamento ${departament}. Error interno del servidor`,
              requestError: `No se pudo listar las provincias de departamento ${departament}. No se recibió respuesta del servidor`,
              otherError: `No se pudo listar las provincias de departamento ${departament}. Ocurrió un error no esperado al traer los datos del servidor`,
            }}
            addButton={{
              display: true,
              disabled: false,
              action: null
            }}
            saveButton={{
              display: true,
              disabled: false,
              action: null
            }}
            pagination = {{
              display: false,
              step: 10,
              totalPages: 0,
              actualPage: 1
            }}
            actionButtons={[]} 
            on:alert={handleTableAlert}
            jwtToken={localStorage.getItem('jwtToken')}
          />
        </div>
      </div>
      <div class="col-md-4 {!displayDistricts ? 'd-none' : ''}" style="padding-left: 30px; padding-right: 20px;">
        <div class="row subtitle-row">
          <h4 class="subtitle">Distritos de la provincia de '{province}'</h4>
          <DataTable 
            recordId={'_id'}
            bind:this={districtsDataTable}
            fetchURL={URLS.CLASSROOM_SERVICE + `api/v1/locations/provinces/${provinceId}/districts`}
            saveURL={URLS.CLASSROOM_SERVICE + `api/v1/locations/provinces/${provinceId}/districts`}
            columnKeys={['_id', 'name', ]}
            columnTypes={['id', 'input[text]', ]}
            columnNames={['ID', 'Nombre', 'Acciones']}
            columnStyles={['max-width: 50px;', 'max-width: 150px;', 'max-width: 240px;']}
            columnClasses={['d-none', '', 'text-end']}
            messages = {{
              success: 'Datos actualizados', 
              errorNetwork: `No se pudo listar los distritos de la provincia ${province}. No hay conexión con el servidor.`,
              notFound: `No se pudo listar los distritos de la provincia ${province}. Recurso no encontrado.`,
              serverError:`No se pudo listar los distritos de la provincia ${province}. Error interno del servidor`,
              requestError: `No se pudo listar los distritos de la provincia ${province}. No se recibió respuesta del servidor`,
              otherError: `No se pudo listar los distritos de la provincia ${province}. Ocurrió un error no esperado al traer los datos del servidor`,
            }}
            addButton={{
              display: true,
              disabled: false,
              action: null
            }}
            saveButton={{
              display: true,
              disabled: false,
              action: null
            }}
            pagination = {{
              display: false,
              step: 10,
              totalPages: 0,
              actualPage: 1
            }}
            actionButtons={[]} 
            on:alert={handleTableAlert}
            jwtToken={localStorage.getItem('jwtToken')}
          />
        </div>
      </div>
    </div>
  </div>
</div>