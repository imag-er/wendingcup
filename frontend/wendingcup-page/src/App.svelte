<script>
  import Router from "svelte-spa-router";
  import { link } from "svelte-spa-router";
  import Login from "./lib/Login.svelte";
  import TeamDashboard from "./lib/TeamDashboard.svelte";
  import Leaderboard from "./lib/Leaderboard.svelte";
  import FileUpload from "./lib/FileUpload.svelte";
  import { fade } from "svelte/transition";

  const routes = {
    "/": Login,
    "/dashboard": TeamDashboard,
    "/leaderboard": Leaderboard,
    "/submit": FileUpload,
  };
</script>

<main>
  <nav class="bg-white shadow-sm">
    <div class="max-w-7xl mx-auto px-4">
      <div class="flex justify-between h-16">
        <a
          use:link
          href="/dashboard"
          class="flex items-center text-gray-700 hover:text-blue-600 transition-colors"
        >
          队伍中心
        </a>
        <a
          use:link
          href="/leaderboard"
          class="flex items-center text-gray-700 hover:text-blue-600 transition-colors"
        >
          实时榜单
        </a>
        <a
          use:link
          href="/submit"
          class="flex items-center text-gray-700 hover:text-blue-600 transition-colors"
        >
          提交文件
        </a>

      </div>
    </div>
  </nav>

  <div class="max-w-7xl mx-auto px-4 py-8" in:fade>
    <Router {routes} />
  </div>
</main>

