<script>
  import { onMount } from 'svelte';
  import apiConfig from './apiConfig.json';

  let endpoints = [];
  let selectedEndpoint = null;
  let formData = {};
  let response = null;
  let error = null;

  onMount(async () => {
    endpoints = apiConfig;
  });

  async function submitRequest() {
    try {
      const form = new FormData();
      
      selectedEndpoint.requestFields.forEach(field => {
        if (field.type === 'bytes') {
          form.append(field.name, formData[field.name]);
        } else {
          form.append(field.name, formData[field.name]);
        }
      });

      const res = await fetch(selectedEndpoint.path, {
        method: selectedEndpoint.method,
        body: form
      });

      response = await res.json();
      error = null;
    } catch (err) {
      error = err.message;
      response = null;
    }
  }
</script>

<main class="container">
  <h1>API测试工具</h1>
  
  <div class="endpoint-select">
    <select bind:value={selectedEndpoint}>
      <option value={null}>选择接口</option>
      {#each endpoints as endpoint}
        <option value={endpoint}>{endpoint.name} ({endpoint.method})</option>
      {/each}
    </select>
  </div>

  {#if selectedEndpoint}
    <form on:submit|preventDefault={submitRequest}>
      {#each selectedEndpoint.requestFields as field}
        <div class="form-field">
          <label>{field.name} ({field.type})</label>
          {#if field.type === 'bytes'}
            <input 
              type="file"
              on:change={(e) => formData[field.name] = e.target.files[0]}
            />
          {:else}
            <input 
              type="text"
              bind:value={formData[field.name]}
              placeholder={`输入${field.name}`}
            />
          {/if}
        </div>
      {/each}
      <button type="submit">发送请求</button>
    </form>
  {/if}

  {#if error}
    <div class="error">{error}</div>
  {:else if response}
    <div class="response">
      <pre>{JSON.stringify(response, null, 2)}</pre>
    </div>
  {/if}
</main>

<style>
  .container {
    max-width: 800px;
    margin: 0 auto;
    padding: 20px;
  }
  .form-field {
    margin-bottom: 15px;
  }
  label {
    display: block;
    margin-bottom: 5px;
  }
  input, select {
    width: 100%;
    padding: 8px;
  }
  button {
    background: #007bff;
    color: white;
    padding: 10px 20px;
    border: none;
    cursor: pointer;
  }
  .error {
    color: #dc3545;
    margin-top: 20px;
  }
  .response {
    margin-top: 20px;
    background: #f8f9fa;
    padding: 15px;
    border-radius: 4px;
  }
</style>