<script>
  import { onMount } from 'svelte';
  import { activeLink } from '../../../helpers/store'; 

  onMount(() => {
    activeLink.set('/employees');
    document.title = "Gestión de Colaboradores";
  });
</script>

<style></style>

<div class="row">
  <h1 class="mb-2 subtitle"><i class="fa fa-pie-chart me-2"></i>Gestión de Empleados</h1>
</div>
<hr>