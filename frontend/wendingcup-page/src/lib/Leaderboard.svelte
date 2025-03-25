<script>
  import { onMount } from "svelte";
  import { fade } from "svelte/transition";

  let leaderboardData = [];
  let isLoading = true;
  let error = "";
  let sortBy = "score";

  const fetchLeaderboard = async () => {
    try {
      const res = await fetch("http://localhost:8080/board");

      if (res.ok) {
        const data = await res.json();
        leaderboardData = data.sort((a, b) => b.score - a.score);
      } else {
        error = "获取排行榜失败";
      }
    } catch (err) {
      error = "网络连接异常";
    } finally {
      isLoading = false;
    }
  };



  onMount(fetchLeaderboard);
</script>

<main class="max-w-4xl mx-auto p-6">
  {#if isLoading}
    <div class="text-center py-8">
      <div
        class="animate-spin inline-block w-8 h-8 border-4 border-blue-500 rounded-full border-t-transparent"
      ></div>
    </div>
  {:else if error}
    <div class="p-4 bg-red-50 text-red-700 rounded-lg" in:fade>{error}</div>
  {:else}
    <div class="space-y-6" in:fade>
      <h1 class="text-3xl font-bold text-gray-800">实时排行榜</h1>


      <div class="divide-y bg-white rounded-xl shadow-sm">
        {#each leaderboardData as item, index}
          <div class="p-4 hover:bg-gray-50 transition-colors grid grid-cols-2 gap-4 items-center">
            <div class="space-y-1">
              <p class="font-medium text-gray-900">
                #{index + 1}
                {item.teamname || "匿名队伍"}
              </p>
              <p class="text-xs text-gray-500 font-mono">{item.team_id.slice(0, 8)}</p>
            </div>
            <div class="text-right space-y-1">
              <p class="text-lg font-semibold text-blue-600">
                {item.score.toFixed(2)}分
              </p>
              <p class="text-xs text-gray-500">
                {new Date(item.judge_result_time).toLocaleString()}
              </p>
            </div>
          </div>
        {/each}
      </div>
    </div>
  {/if}
</main>
