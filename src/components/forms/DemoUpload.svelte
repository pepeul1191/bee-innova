<script>
  let textInput = '';
  let fileInput = null;
  let responseMessage = '';
  let isLoading = false;

  async function handleSubmit() {
    if (!fileInput) {
      alert('Por favor, selecciona un archivo');
      return;
    }

    isLoading = true;
    responseMessage = '';

    const formData = new FormData();
    formData.append('file', fileInput);
    formData.append('folder', textInput);
    const token = localStorage.getItem('jwtFilesToken'); // obtén el JWT
    console.log(token)

    try {
      const res = await fetch(
        URLS.FILE_SERVICE + 'api/v1/files', 
        {
          credentials: 'include', // si usas cookies/sesión
          method: 'POST',
          body: formData,
          headers: {
            // no agregues Content-Type, usa FormData
            'Authorization': `Bearer ${token}`  // << aquí
          },
        });

      if (!res.ok) throw new Error(`Error del servidor: ${res.status}`);
      const data = await res.json();
      responseMessage = `✅ Éxito: ${JSON.stringify(data)}`;
    } catch (err) {
      responseMessage = `❌ Error: ${err.message}`;
    } finally {
      isLoading = false;
    }
  }
</script>

<!-- Resto del HTML igual -->

<style>
  .loading {
    color: blue;
  }

  .response {
    margin-top: 1rem;
    color: green;
  }
</style>

<h2>Formulario de Subida</h2>

<form on:submit|preventDefault={handleSubmit}>
  <div>
    <label>
      Texto:
      <input type="text" bind:value={textInput} placeholder="Ingresa un texto" />
    </label>
  </div>

  <div>
    <label>
      Archivo:
      <input type="file" on:change={(e) => {
        const files = e.target.files;
        if (files && files.length > 0) {
          fileInput = files[0];
        }
      }} />
    </label>
  </div>

  <button type="submit" disabled={isLoading}>
    {isLoading ? 'Enviando...' : 'Enviar'}
  </button>
</form>

{#if responseMessage}
  <div class="response">{responseMessage}</div>
{/if}