<script>
  import { onMount } from 'svelte';
  import { navigate } from 'svelte-routing';
  import Logo from '../../svgs/Logo.svelte';
  import InputText from '../../widgets/InputText.svelte';
  import InputPassword from '../../widgets/InputPassword.svelte';
  import InputEmail from '../../widgets/InputEmail.svelte';
  import { createUser } from '../../../services/access/user_service.js';

  let message = '';
  let messageClass = '';

  let username = ''; let usernameInput;
  let email = ''; let emailInput;
  let password1 = ''; let password1Input;
  let password2 = ''; let password2Input;
  
  onMount(() => {

  });

  const create = (event) => {
    event.preventDefault();
    // run validations
    usernameInput.validate();
    emailInput.validate();
    password1Input.validate();
    password2Input.validate();
    // if ok, ajax
    if(
      usernameInput.isValid && 
      emailInput.isValid && 
      password1Input.isValid && 
      password2Input.isValid
    ){
      if(password1 == password2){
        let data = {
          username: username,
          email: email, 
          password: password1,
        };
        createUser(data).then((resp) => {
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
      }else{
        password1Input.isValid = false; 
        password1Input.message = ''; 
        password2Input.isValid = false;
        password2Input.message = 'Contraseñas deben de coincidir';
        setTimeout(() => {
          password1Input.isValid = true; 
          password2Input.isValid = true;
        }, 4000);
      }
    }
  }
</script>
<style></style>

<svelte:head>
	<title>Crear Cuenta</title>
</svelte:head>

<div class="container mt-5">
  <div class="row justify-content-center mb-4">
    <Logo size=48 color='#0000002d'/>
  </div>
  <div class="row justify-content-center">
    <div class="col-md-4">
      <div class="card">
        <div class="card-header">
          <h3 class="text-center">Agregar Cuenta</h3>
        </div>
        <div class="card-body">
          <form>
            <div class="mb-3">
              <InputText 
                id="username" 
                label="Usuario" 
                bind:value={username} 
                bind:this={usernameInput}
                onInputValidation={true}
                validations = {[
                  {type: 'notEmpty', message: 'Ingresar un usuario'},
                  {type: 'minLength', message: 'Mínimo 5 caracteres', length: 5},
                  {type: 'maxLength', message: 'Máximo 20 caracteres', length: 20},
                ]} />
            </div>
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
            <div class="mb-3">
              <InputPassword 
                id="password1" 
                label="Contraseña" 
                bind:value={password1} 
                bind:this={password1Input}
                onInputValidation={true}
                validations = {[
                  {type: 'notEmpty', message: 'Ingresar una contraseña'},
                  {type: 'minLength', message: 'Mínimo 5 caracteres', length: 5},
                  {type: 'maxLength', message: 'Máximo 20 caracteres', length: 20},
                  {type: 'isSecure', message: 'Contraseña no segura'},
                ]} />
            </div>
            <div class="mb-3">
              <InputPassword 
                id="password2" 
                label="Repetir Contraseña" 
                bind:value={password2} 
                bind:this={password2Input}
                onInputValidation={true}
                validations = {[
                  {type: 'notEmpty', message: 'Ingresar una contraseña'},
                  {type: 'minLength', message: 'Mínimo 5 caracteres', length: 5},
                  {type: 'maxLength', message: 'Máximo 20 caracteres', length: 20},
                  {type: 'isSecure', message: 'Contraseña no segura'},
                ]} />
            </div>
            <button type="submit" class="btn btn-primary w-100" on:click={create}>Crear Cuenta</button>
          </form>
          <div class="text-center mt-3 {messageClass}">
            {message}
          </div>
          <div class="text-center mt-3">
            <a class="navbar-brand" href="/sign-in" on:click|preventDefault={() => {navigate('/sign-in')}}>Ingresar</a>
            <span class="mx-2">|</span>
            <a class="navbar-brand" href="/reset-password" on:click|preventDefault={() => {navigate('/reset-password')}}>Recuperar Contraseña</a>
          </div>
        </div>
        <div class="card-footer text-center">
          <small>&copy; 2024 Your Brand</small>
        </div>
      </div>
    </div>
  </div>
</div>