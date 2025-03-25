<script >
  import { fade } from 'svelte/transition';
  
  export let team_id;
  
  // 明确指定 file 变量的类型为 File 或 null
  let file;
  let error = '';
  let success = '';
  
  const handleUpload = async () => {
    if (!file) {
      error = '请选择要上传的文件';
      return;
    }

    const formData = new FormData();
    formData.append('team_id', team_id);
    formData.append('file', file);

    try {
      const token = localStorage.getItem('jwt');
      const res = await fetch('http://localhost:8080/auth/submit', {
        method: 'POST',
        headers: { Authorization: `Bearer ${token}` },
        body: formData
        
      });

      if (res.ok) {
        success = '文件上传成功';
        error = '';
        file = null;
        setTimeout(() => success = '', 3000);
      } else {
        error = `上传失败: ${await res.text()}`;
      }
    } catch (err) {
      error = '网络连接异常';
    }
  };
</script>

<div class="bg-white rounded-xl shadow-sm p-6 mb-8" in:fade>
  <h2 class="text-xl font-semibold mb-4">提交新文件</h2>
  
  <div class="space-y-4">
    {#if error}
      <div class="p-3 bg-red-50 text-red-700 rounded-md">{error}</div>
    {/if}
    
    {#if success}
      <div class="p-3 bg-green-50 text-green-700 rounded-md">{success}</div>
    {/if}

    <div class="flex gap-4 items-center">
      <input
        type="file"
        accept=".zip"
        bind:files={file}
        class="block w-full text-sm text-gray-500
          file:mr-4 file:py-2 file:px-4
          file:rounded-md file:border-0
          file:text-sm file:font-semibold
          file:bg-blue-50 file:text-blue-700
          hover:file:bg-blue-100"
      />
      <button
        on:click={handleUpload}
        class="px-6 py-2 bg-blue-600 text-white rounded-md
          hover:bg-blue-700 transition-colors"
      >
        上传
      </button>
    </div>
    
    <p class="text-sm text-gray-500">
      仅支持ZIP格式，文件大小不超过10MB
    </p>
  </div>
</div>