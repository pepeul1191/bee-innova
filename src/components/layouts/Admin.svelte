<script>
  import { Router, Route, navigate } from 'svelte-routing';
  import { onMount } from 'svelte';
  import Sidebar from '../widgets/Sidebar.svelte';
  import Navbar from '../widgets/Navbar.svelte';
  import Home from '../pages/app/Home.svelte';
  import Asste from '../pages/app/Asset.svelte';
  import Issue from '../pages/app/Issue.svelte';
  import IssueDetail from '../pages/app/IssueDetail.svelte';
  import Employee from '../pages/app/Employee.svelte';
  import Tag from '../pages/app/Tag.svelte';
	import Asset from '../pages/app/Asset.svelte';
	import Location from '../pages/admin/Location.svelte';
	import Teacher from '../pages/admin/Teacher.svelte';
	import TeacherDetail from '../pages/admin/TeacherDetail.svelte';
  export let basepath = '/';

  let userRoles = [];
  let username = '';
  let userImage = '';

  const fetchInfoIfMissing = () => {
    // Verificar si los tokens ya existen
    var jwtToken = localStorage.getItem('jwtToken');

    // Si ambos tokens ya están guardados, no hacer nada
    if (jwtToken) {
      console.log('Token ya existen en localStorage.');
      return Promise.resolve(); // Salir sin hacer la petición
    }

    // Si alguno falta, hacer el GET a /tokens
    return new Promise(function(resolve, reject) {
      var xhr = new XMLHttpRequest();
      xhr.open('GET', '/session', true);
      xhr.setRequestHeader('Content-Type', 'application/json');
      
      xhr.onload = function() {
        if (xhr.status >= 200 && xhr.status < 300) {
          var response = JSON.parse(xhr.responseText);
          var tokens = response.tokens;
          var person = response.person;
          var user = response.user;

          // Guardar ambos tokens en localStorage
          if (tokens) {
            localStorage.setItem('jwtClassroomToken', tokens.classroom);
            localStorage.setItem('jwtFilesToken', tokens.files);
            localStorage.setItem('jwtAccessToken', tokens.access);
            localStorage.setItem('full_name', person.last_names + ', ' + person.names);
            localStorage.setItem('image_url', URLS.FILES_SERVICE + person.image_url);
            localStorage.setItem('roles', JSON.stringify(user.roles));
            username = person.last_names + ', ' + person.names;
            userImage = URLS.FILES_SERVICE + person.image_url;
            userRoles = user.roles;
          }

          console.log('Token guardado en localStorage.');
          resolve();
        } else {
          reject(new Error('Error HTTP: ' + xhr.status));
        }
      };

      xhr.onerror = function() {
        reject(new Error('Error de red'));
      };

      xhr.send();
    })
    .catch(function(error) {
      console.error('Error al obtener tokens:', error);
      return Promise.reject(error);
    });
  }

  onMount(() => {
    // Verificar si los tokens ya existen
    const jwtTicketToken = localStorage.getItem('jwtTicketsToken');
    const jwtFileToken = localStorage.getItem('jwtFilesToken');
    const jwtAccessToken = localStorage.getItem('jwtAccessToken');
    const roles = localStorage.getItem('roles');
    const fullName = localStorage.getItem('full_name');
    const imageUrl = localStorage.getItem('image_url');

    // Si los tokens ya están presentes, no hacemos nada
    if (!jwtTicketToken || !jwtFileToken || !jwtAccessToken || !fullName || !imageUrl || !roles) {
      fetchInfoIfMissing()
        .then(function() {
          console.log('Tokens listos para usar.');
        })
        .catch(function(err) {
          console.error('No se pudieron obtener los tokens:', err);
        });
    }

    if(!userImage || !username){
      username = localStorage.getItem('full_name');
      userImage = localStorage.getItem('image_url');
    }
  });
</script>
  
<style>

</style>

<Sidebar {userRoles} />
<Navbar {username} {userImage} />

<main class="main-content mt-5">
  <div class="container-fluid">
    <Router basepath="{basepath}">
      <Route path="/" component={Home} />
      <Route path="/locations" component={Location} />
      <Route path="/teachers" component={Teacher} />
      <Route path="/teachers/new" component={TeacherDetail} />

      <Route path="/assets" component={Asset} />
      <Route path="/issues" component={Issue} />
      <Route path="/tags" component={Tag} />
      <Route path="/employees" component={Employee} />
      <Route path="/issues/new" component={IssueDetail} />
      <Route path="/issues/:_id/edit/" let:params><IssueDetail _id={params._id}/></Route>
    </Router>
  </div>
</main>

<!-- Footer -->
<footer class="bg-light text-center py-3 mt-auto">
  <div class="container">
    <p class="mb-0">© 2025 Innova ULima. Todos los derechos reservados.</p>
  </div>
</footer>