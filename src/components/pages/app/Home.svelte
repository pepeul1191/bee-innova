<script>
  import { onMount } from 'svelte';
  import { activeLink } from '../../../helpers/store';

  // Llamada a los datos de ejemplo (esto se simula aquí, en la práctica se usaría una API)
  let totalIssues = 120;
  let totalEmployees = 50;
  let totalAssets = 35;
  let totalTags = 25;

  let issues = [
    { id: 1, resume: "Falla en el servidor", description: "El servidor principal no responde", issueState: "Abierta", priority: "Alta", reportDate: "2025-07-15" },
    { id: 2, resume: "Error en el sistema de login", description: "El sistema de login muestra un error 500", issueState: "En progreso", priority: "Media", reportDate: "2025-07-16" },
    { id: 3, resume: "Problema con la impresora", description: "La impresora no imprime correctamente", issueState: "Cerrada", priority: "Baja", reportDate: "2025-07-17" },
    { id: 4, resume: "Fallo en la red interna", description: "La red interna no permite conectar a algunos equipos", issueState: "Abierta", priority: "Alta", reportDate: "2025-07-18" }
  ];

  // Función para manejar la acción de ver detalles (puedes cambiarla a lo que necesites)
  const viewIssue = (id) => {
    alert(`Ver detalles de la incidencia ${id}`);
  };

  // Función para manejar la acción de editar la incidencia
  const editIssue = (id) => {
    alert(`Editar la incidencia ${id}`);
  };

  // Se establecen los valores del activeLink y el título de la pestaña
  onMount(() => {
    activeLink.set('/dashboard');
    document.title = "Dashboard - Gestión de Incidencias";
  });
</script>

<style>
  .dashboard-card {
    background: #ffffff;
    border-radius: 10px;
    padding: 20px;
    box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
    text-align: center;
  }

  .dashboard-card i {
    font-size: 2rem;
    color: #007bff;
  }

  .dashboard-card .count {
    font-size: 1.5rem;
    font-weight: bold;
    color: #333;
  }

  .row {
    display: flex;
    justify-content: space-between;
    gap: 20px;
  }

  .col-3 {
    flex: 1;
    min-width: 200px;
  }

  .subtitle {
    font-size: 1.75rem;
    color: #333;
  }

  .icon-wrapper {
    display: flex;
    align-items: center;
    justify-content: center;
    margin-bottom: 10px;
  }

  table {
    width: 100%;
    border-collapse: collapse;
    margin: 20px 0;
  }

  th, td {
    padding: 12px;
    text-align: left;
    border-bottom: 1px solid #ddd;
  }

  th {
    background-color: #f2f2f2;
  }

  tr:hover {
    background-color: #f5f5f5;
  }

  .btn {
    padding: 6px 12px;
    margin: 3px;
    text-align: center;
    cursor: pointer;
    border-radius: 5px;
    font-size: 0.9rem;
    color: white;
  }

  .btn-view {
    background-color: #007bff;
  }

  .btn-edit {
    background-color: #28a745;
  }

  .btn-delete {
    background-color: #dc3545;
  }

</style>

<div class="row">
  <h1 class="mb-2 subtitle"><i class="fa fa-tachometer-alt me-2"></i><i class="fa fa-pie-chart me-2" aria-hidden="true"></i> Dashboard de Incidencias</h1>
</div>

<hr>

<!-- Cards de estadísticas del dashboard -->
<div class="row">
  <!-- Total de Incidencias -->
  <div class="col-3">
    <div class="dashboard-card">
      <div class="icon-wrapper">
        <i class="fa fa-exclamation-triangle"></i>
      </div>
      <div class="count">{totalIssues}</div>
      <p>Incidencias en total</p>
    </div>
  </div>

  <!-- Total de Empleados -->
  <div class="col-3">
    <div class="dashboard-card">
      <div class="icon-wrapper">
        <i class="fa fa-users"></i>
      </div>
      <div class="count">{totalEmployees}</div>
      <p>Empleados registrados</p>
    </div>
  </div>

  <!-- Total de Activos -->
  <div class="col-3">
    <div class="dashboard-card">
      <div class="icon-wrapper">
        <i class="fa fa-cogs"></i>
      </div>
      <div class="count">{totalAssets}</div>
      <p>Activos registrados</p>
    </div>
  </div>

  <!-- Total de Etiquetas -->
  <div class="col-3">
    <div class="dashboard-card">
      <div class="icon-wrapper">
        <i class="fa fa-tags"></i>
      </div>
      <div class="count">{totalTags}</div>
      <p>Etiquetas disponibles</p>
    </div>
  </div>
</div>

<!-- Gráfico o estadísticas adicionales -->
<!-- En la práctica, podrías integrar gráficos aquí con librerías como Chart.js o ApexCharts -->

<div class="table-container mt-4">
  <h2>Lista de Incidencias</h2>
  <table>
    <thead>
      <tr>
        <th>ID</th>
        <th>Resumen</th>
        <th>Descripción</th>
        <th>Estado</th>
        <th>Prioridad</th>
        <th>Fecha de Reporte</th>
        <th>Acciones</th>
      </tr>
    </thead>
    <tbody>
      {#each issues as issue}
        <tr>
          <td>{issue.id}</td>
          <td>{issue.resume}</td>
          <td>{issue.description}</td>
          <td>{issue.issueState}</td>
          <td>{issue.priority}</td>
          <td>{issue.reportDate}</td>
          <td>
            <button class="btn btn-view" on:click={() => viewIssue(issue.id)}>Ver</button>
            <button class="btn btn-edit" on:click={() => editIssue(issue.id)}>Editar</button>
            <button class="btn btn-delete" on:click={() => alert(`Eliminar incidencia ${issue.id}`)}>Eliminar</button>
          </td>
        </tr>
      {/each}
    </tbody>
  </table>
</div>