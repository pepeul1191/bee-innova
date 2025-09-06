<script>
  import { onMount } from 'svelte';
  import { navigate } from 'svelte-routing';
  import Logo from '../../svgs/Logo.svelte';
  import InputText from '../../widgets/InputText.svelte';
  import InputPassword from '../../widgets/InputPassword.svelte';
  import { validateUser } from '../../../services/access/user_service.js';

  let username = ''; let usernameInput;
  let password = ''; let passwordInput;
  let message = '';
  let messageClass = '';
  
  onMount(() => {
    const search = location.search;
    // Parsear los query params
    const urlParams = new URLSearchParams(search);
    const errorParam = urlParams.get('error');
    if (errorParam) {
      message = decodeURIComponent(errorParam);
      messageClass = 'text-danger';
    } else {
      message = '';
      messageClass = '';
    }
  });

  const access = (event) => {
    event.preventDefault();
    // run validations
    usernameInput.validate();
    passwordInput.validate();
    // if ok, ajax
    if(usernameInput.isValid && passwordInput.isValid){
      let data = {
        username: username, 
        password: password, 
      };
      validateUser(data).then((resp) => {
        //console.log(resp)
        message = resp.data.message;
        localStorage.setItem('jwtFilesToken', resp.data.files_token);
        localStorage.setItem('jwtAccessToken', resp.data.access_token);
        localStorage.setItem('jwtTicketsToken', resp.data.tickets_token);
        messageClass = 'text-success';
        setTimeout(() => {
          message = '';
          messageClass = '';
          window.location.href = '/';
        }, 1500);
      }).catch((resp) =>  {
        console.log(resp)
        message = resp.message;
        messageClass = 'text-danger';
        setTimeout(() => {
          message = '';
          messageClass = '';
        }, 4000);
      })
    }
  }
</script>
<style>

</style>

<svelte:head>
	<title>Ingresar al Sistema</title>
</svelte:head>
<div class="container mt-5">
  <div class="row justify-content-center mb-4">
    <Logo size=48 color='#0000002d'/>
  </div>
  <div class="row justify-content-center">
    <div class="col-md-4">
      <div class="card">
        <div class="card-header">
          <h3 class="text-center">Bienvenido</h3>
        </div>
        <div class="card-body">
          <form>
            <div class="mb-3">
              <InputText 
                id="username" 
                label="Usuario" 
                bind:value={username} 
                bind:this={usernameInput}
                onInputValidation={false}
                validations = {[
                  {type: 'notEmpty', message: 'Ingresar un usuario'},
                ]} />
            </div>
            <div class="mb-3">
              <InputPassword 
                id="password" 
                label="Contraseña" 
                bind:value={password} 
                bind:this={passwordInput}
                onInputValidation={false}
                validations = {[
                  {type: 'notEmpty', message: 'Ingresar una contraseña'},
                ]} />
            </div>
            <button  class="btn btn-primary w-100" on:click={access}>Ingresar</button>
          </form>
          <div class="text-center mt-3 {messageClass}">
            {message}
          </div>
          <div class="text-center mt-3">
            <a class="navbar-brand" href="/reset-password" on:click|preventDefault={() => {navigate('/reset-password')}}>Recuperar Contraseña</a>
            <span class="mx-2">|</span>
            <a class="navbar-brand" href="/sign-up" on:click|preventDefault={() => {navigate('/sign-up')}}>Crear Cuenta</a>
          </div>
        </div>
        <div class="card-footer text-center">
          <small>&copy; 2024 Your Brand</small>
        </div>
      </div>
    </div>
  </div>
</div>