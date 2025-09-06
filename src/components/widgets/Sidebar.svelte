<script>
  import { navigate } from 'svelte-routing';
  import { activeLink } from '../../helpers/store.js';
	import { onMount } from 'svelte';

  export let userRoles;

  onMount(() => {

  });
</script>

<div class="sidebar d-none d-md-block">
  <div class="sidebar-header">
    <a class="navbar-brand text-dark" href="/" on:click|preventDefault={() => {navigate('/')}}>Aula Virtual</a>
  </div>
  <ul class="nav flex-column mt-3">
    <li class="nav-item">
      <a class="nav-link" class:active={$activeLink === '/'} href="/" on:click|preventDefault={() => {navigate('/')}}>
        <i class="fa fa-pie-chart me-2" aria-hidden="true"></i> Dashboard
      </a>
    </li>
    {#each userRoles as r}
      {#if r.name == 'admin'}
        <li class="nav-item">
          <a class="nav-link" href="#">
            <i class="fa fa-users me-2"></i> Usuarios???
          </a>
        </li>
        <!-- Gestión de Locaciones -->
        <li class="nav-item">
          <a class="nav-link" class:active={$activeLink === '/locations'} href="/locations" on:click|preventDefault={() => {navigate('/locations')}}>
            <i class="fa fa-map me-2"></i> Locaciones
          </a>
        </li>

        <!-- Gestión de Docentes -->
        <li class="nav-item">
          <a class="nav-link" class:active={$activeLink === '/teachers'} href="/teachers" on:click|preventDefault={() => {navigate('/teachers')}}>
            <i class="fa fa-graduation-cap me-2"></i> Docentes
          </a>
        </li> 

        <!-- Gestión de Alumnos -->
        <li class="nav-item">
          <a class="nav-link" class:active={$activeLink === '/students'} href="/students" on:click|preventDefault={() => {navigate('/students')}}>
            <i class="fa fa-users me-2"></i> Alumnos
          </a>
        </li>

        <!-- Gestión de Apoderados -->
        <li class="nav-item">
          <a class="nav-link" class:active={$activeLink === '/guardians'} href="/guardians" on:click|preventDefault={() => {navigate('/guardians')}}>
            <i class="fa fa-user-plus me-2"></i> Apoderados
          </a>
        </li>

        <!-- Gestión de Cursos -->
        <li class="nav-item">
          <a class="nav-link" class:active={$activeLink === '/courses'} href="/courses" on:click|preventDefault={() => {navigate('/courses')}}>
            <i class="fa fa-book me-2"></i> Cursos y Sessciones
          </a>
        </li>
      {/if}
      {#if r.name == 'teacher'}
        <!-- Gestión de Secciones -->
        <li class="nav-item">
          <a class="nav-link" class:active={$activeLink === '/sections'} href="/sections" on:click|preventDefault={() => {navigate('/sections')}}>
            <i class="fa fa-th-large me-2"></i> Secciones
          </a>
        </li>
        <li class="nav-item">
          <a class="nav-link" class:active={$activeLink === '/issues'} href="/issues" on:click|preventDefault={() => {navigate('/issues')}}>
            <i class="fa fa-bug me-2"></i> Mensajes
          </a>
        </li>
      {/if}
    {/each}

  </ul>
</div>
