<script>
  import { onMount } from 'svelte';
  import { navigate } from 'svelte-routing';
  import Logo from '../../svgs/Logo.svelte';
  import InputEmail from '../../widgets/InputEmail.svelte';
  import { resetPassword } from '../../../services/access/user_service.js';
  
  let message = '';
  let messageClass = '';
  let email = ''; let emailInput;

  onMount(() => {

  });

  const reset = (event) => {
    event.preventDefault();
    // run validations
    emailInput.validate();
    // if ok, ajax
    if(emailInput.isValid){
      let data = {
        email: email, 
      };
      resetPassword(data).then((resp) => {
          //console.log(resp)
          message = resp.data;
          messageClass = 'text-success';
          setTimeout(() => {
            message = '';
            messageClass = '';
          }, 4000);
        }).catch((resp) =>  {
          //console.log(resp)
          message = resp.data;
          messageClass = 'text-danger';
          setTimeout(() => {
            message = '';
            messageClass = '';
          }, 4000);
        })
    }
  }
</script>
<style></style>

<svelte:head>
	<title>Recuperar Contraseña</title>
</svelte:head>

<div class="container mt-5">
  <div class="row justify-content-center mb-4">
    <Logo size=48 color='#0000002d'/>
  </div>
  <div class="row justify-content-center">
    <div class="col-md-4">
      <div class="card">
        <div class="card-header">
          <h3 class="text-center">Olvidó su Contraseña</h3>
        </div>
        <div class="card-body">
          <form>
            <div class="mb-3">
              <InputEmail 
                id="email" 
                label="Correo Electrónico" 
                bind:value={email} 
                bind:this={emailInput}
                onInputValidation={true}
                validations = {[
                  {type: 'notEmpty', message: 'Ingresar un correo'},
                  {type: 'validEmail', message: 'Correo no válido'},
                ]} />
            </div>
            <button type="submit" on:click={reset} class="btn btn-primary w-100">Enviar Correo</button>
          </form>
          <div class="text-center mt-3 {messageClass}">
            {message}
          </div>
          <div class="text-center mt-3">
            <a class="navbar-brand" href="/sign-in" on:click|preventDefault={() => {navigate('/sign-in')}}>Ingresar</a>
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