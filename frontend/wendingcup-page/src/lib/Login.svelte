<script>
  import { fade } from "svelte/transition";

  let teamId = "";
  let error = "";

  async function handleLogin(e) {
    e.preventDefault();
    const formData = new FormData();
    formData.append("team_id", teamId);

    try {
      const res = await fetch("http://localhost:8080/login", {
        method: "POST",
        body: formData,
      });

      if (res.ok) {
        const { token } = await res.json();
        localStorage.setItem("team_id", teamId);
        localStorage.setItem("jwt", token);

        window.location.href = "/dashboard";
      } else {
        error = "登录失败，请检查队伍ID";
      }
    } catch (err) {
      console.log(err)
      error = "网络连接异常";
    }
  }
</script>

<main class="max-w-xl mx-auto mt-20 p-8 bg-white rounded-xl shadow-lg">
  <form on:submit={handleLogin} class="space-y-6" in:fade>
    <h2 class="text-3xl font-bold text-gray-800 mb-6">队伍登录</h2>

    {#if error}
      <div class="p-3 bg-red-100 text-red-700 rounded-md" in:fade>{error}</div>
    {/if}

    <div>
      <input
        id="teamId"
        type="text"
        bind:value={teamId}
        class="w-full px-5 py-3 text-lg border-2 border-gray-200 rounded-lg
          focus:ring-4 focus:ring-blue-500/20 focus:border-blue-500
          transition-all duration-200"
        placeholder="输入36位队伍ID"
        required
      />
    </div>

    <button
      type="submit"
      class="w-full bg-gradient-to-r from-blue-500 to-blue-600 text-white
        py-4 px-6 rounded-lg hover:from-blue-600 hover:to-blue-700
        transition-all duration-200 transform hover:scale-[1.02]"
    >
      登录
    </button>
  </form>
</main>

<style global>
  /* 添加全局过渡动画 */
  :global(main) {
    view-transition-name: main-container;
  }

  @keyframes fadeIn {
    from {
      opacity: 0;
      transform: translateY(20px);
    }
    to {
      opacity: 1;
      transform: translateY(0);
    }
  }

  :global(.fade-in) {
    animation: fadeIn 0.3s ease-out;
  }
</style>
