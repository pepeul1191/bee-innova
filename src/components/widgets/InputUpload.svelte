<script>
  export let label = "Seleccionar imagen";
  export let acceptedFormats = ["jpg", "jpeg", "png"];
  export let maxSizeMB = 2;
  export let url = "/api/upload";
  export let fileKey = "file";
  export let extraParams = {};
  export let jwt = "";
  export let helpText = `Formatos aceptados: ${acceptedFormats.join(", ").toUpperCase()} (Máx. ${maxSizeMB}MB)`;
  export let onSuccess = (response) => console.log("Upload success:", response);
  export let onError = (error) => console.error("Upload error:", error);

  let file = null;
  let isLoading = false;
  let errorMessage = "";
  let successMessage = "";

  const handleFileChange = (event) => {
    const selectedFile = event.target.files[0];
    if (!selectedFile) return;

    // Validar formato
    const fileExt = selectedFile.name.split('.').pop().toLowerCase();
    if (!acceptedFormats.includes(fileExt)) {
      errorMessage = `Formato no válido. Use: ${acceptedFormats.join(", ").toUpperCase()}`;
      return;
    }

    // Validar tamaño
    if (selectedFile.size > maxSizeMB * 1024 * 1024) {
      errorMessage = `El archivo excede el tamaño máximo de ${maxSizeMB}MB`;
      return;
    }

    errorMessage = "";
    file = selectedFile;
  };

  const uploadFile = async () => {
    if (!file) {
      errorMessage = "Por favor seleccione un archivo";
      return;
    }

    isLoading = true;
    errorMessage = "";
    successMessage = "";

    const formData = new FormData();
    formData.append(fileKey, file);

    // Agregar parámetros extra
    Object.keys(extraParams).forEach(key => {
      formData.append(key, extraParams[key]);
    });

    try {
      const response = await fetch(url, {
        method: "POST",
        headers: {
          "Authorization": `Bearer ${jwt}`
        },
        body: formData
      });

      const data = await response.json();

      if (!response.ok) {
        throw new Error(data.message || "Error al subir el archivo");
      }

      successMessage = "Archivo subido correctamente";
      onSuccess(data);
    } catch (error) {
      errorMessage = error.message;
      onError(error);
    } finally {
      isLoading = false;
    }
  };
</script>

<div class="mb-3">
  <label for="imageFile" class="form-label">{label}</label>
  <div class="input-group">
    <input 
      type="file" 
      class="form-control" 
      id="imageFile" 
      accept={acceptedFormats.map(f => `.${f}`).join(',')}
      on:change={handleFileChange}
      disabled={isLoading}
    >
    <button 
      class="btn btn-secondary" 
      type="button" 
      on:click={uploadFile}
      disabled={isLoading || !file}
    >
      {#if isLoading}
        <i class="fa fa-spinner fa-spin"></i> Subiendo...
      {:else}
        <i class="fa fa-upload"></i> Subir
      {/if}
    </button>
  </div>
  <small class="text-muted">{helpText}</small>

</div>