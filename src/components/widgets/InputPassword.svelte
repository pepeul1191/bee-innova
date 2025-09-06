<svelte:options accessors={true} />
<script>
  import { createEventDispatcher } from 'svelte';

  export let value = "";
  export let label = "";
  export let id = "";
  export let validations = [];
  export let onInputValidation = true;
  export let isValid = true;
  export let message = '';

  const dispatch = createEventDispatcher();

  export const validate = () => {
    message = ''; 
    validations.forEach(validation => {
      // validation
      if(validation.type == 'notEmpty'){
        if (value.trim() === ''){
          message = validation.message || 'Debe de llenar este campo';
          isValid = false;
        }else{
          isValid = true;
        }
      }
      if(isValid && validation.type == 'minLength'){
        //console.log('emptyOrMaxLength')
        if(value.length < validation.length){
          message = validation.message || 'No superó la cantidad mínima de caracteres';
          isValid = false;
        }else{
          isValid = true;
        }
      }
      if(isValid && validation.type == 'maxLength'){
        //console.log('emptyOrMaxLength')
        if(value.length > validation.length){
          message = validation.message || 'Superó la cantidad máxima de caracteres';
          isValid = false;
        }else{
          isValid = true;
        }
      }
      if(isValid && validation.type == 'isSecure'){
        //console.log('emptyOrMaxLength')
        if(! /^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[@$!%*?&_\-+=(){}[\]:;<>,./\\|~^"'])[\w@$!%*?&_\-+=(){}[\]:;<>,./\\|~^"']{8,}$/.test(value) ) {
          message = validation.message || 'Contraseña no segura, debe de tener al menos 8 caracteres, un caracter especial, un número, letras minusculas y mayúsculas';
          isValid = false;
        }else{
          isValid = true;
        }
      }
      if(isValid && validation.type == 'custom'){
        validation.func(value).then(resp => {
          if(resp == false){
            message = validation.message || 'No pasó esta validación';
            isValid = false;
          }else{
            isValid = true;
          }
        });
      }
    });
  };

  dispatch('input');
</script>

<div>
  <label for={id} class="form-label">{label}</label>
  <input
    type="password"
    class="form-control { !isValid ? 'is-invalid' : '' }"
    id={id}
    bind:value={value}
    on:input={onInputValidation ? validate : ''} >
  {#if isValid == false}
    <div class="{ !isValid ? 'text-danger' : '' }">
      {message}
    </div>
  {/if}
</div>

<style>
  .is-invalid {
    border-color: #dc3545;
  }

  .text-danger{
    margin-top: 5px;
    color: #842029;
  }
</style>
