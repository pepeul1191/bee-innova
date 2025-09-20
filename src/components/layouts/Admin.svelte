<script>
  import { onMount } from 'svelte';
	import NavBar from '../widgets/NavBar.svelte';
  import SideBar from '../widgets/SideBar.svelte';
  import Footer from '../widgets/Footer.svelte';

  export let basepath = '/';

  onMount(() => {
  });
</script>
  
<style>

</style>

<NavBar />

<div class="container-fluid">
    <div class="row">
        <SideBar />

        <main class="col-md-9 ms-sm-auto col-lg-10 px-md-4">
            <div class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
                <h1 class="h2">Contenido Principal</h1>
            </div>
            
            <p>
                Este es el área principal de tu dashboard.
            </p>
        </main>
    </div>
</div>