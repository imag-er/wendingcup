<script>
  import { onMount } from "svelte";
  import { fade } from "svelte/transition";

  let teamData = null;
  let submissions = [];
  let isLoading = true;
  let error = "";

  const fetchTeamData = async () => {
    const token = localStorage.getItem("jwt");
    if (!token) window.location.href = "/";

    const team_id = localStorage.getItem("team_id");
    if (!team_id) window.location.href = "/";

    try {
      const res = await fetch(`http://localhost:8080/auth/team/${team_id}`, {
        method: "GET",
        headers: { Authorization: `Bearer ${token}` },
      });

      if (res.ok) {
        teamData = await res.json();
        await fetchSubmissions();
      } else if (res.status === 401) {
        window.location.href = "/";
      } else {
        error = "获取队伍信息失败";
      }
    } catch (err) {
      error = "网络连接异常";
    } finally {
      isLoading = false;
    }
  };

  const fetchSubmissions = async () => {
    const token = localStorage.getItem("jwt");
    if (!token) window.location.href = "/";

    const team_id = localStorage.getItem("team_id");
    if (!team_id) window.location.href = "/";
    try {
      const res = await fetch(`http://localhost:8080/auth/submit/${team_id}`, {
        headers: { Authorization: `Bearer ${token}` },
      });

      if (res.ok) {
        const data = await res.json();
        submissions = data.submit_list;
      } else if (res.status === 401) {
        window.location.href = "/";
      } else {
        error = "获取队伍信息失败";
      }
    } catch (err) {
      console.error("获取提交记录失败:", err);
    }
  };

  onMount(fetchTeamData);
</script>

<main class="max-w-4xl mx-auto p-6">
  {#if isLoading}
    <div class="text-center py-8">
      <div
        class="animate-spin inline-block w-8 h-8 border-4 border-blue-500 rounded-full border-t-transparent"
      ></div>
    </div>
  {:else if error}
    <div class="p-4 bg-red-50 text-red-700 rounded-lg">{error}</div>
  {:else}
    <div class="space-y-8">
      <h1 class="text-3xl font-bold text-gray-800" in:fade>队伍信息</h1>

      <div class="bg-white rounded-xl shadow-sm p-6" in:fade>
        <h2 class="text-xl font-semibold mb-4">{teamData.teamname}</h2>
        <div class="space-y-3">
          {#each teamData.players as player}
            <div
              class="flex items-center space-x-4 p-3 hover:bg-gray-50 rounded-lg transition-colors"
            >
              <div class="flex-1">
                <p class="font-medium">{player.name}</p>
                <p class="text-sm text-gray-600">学号: {player.student_id}</p>
                <p class="text-sm text-gray-600">电话: {player.phonenumber}</p>
              </div>
            </div>
          {/each}
        </div>
      </div>

      <div class="bg-white rounded-xl shadow-sm p-6" in:fade>
        <h2 class="text-xl font-semibold mb-4">提交记录</h2>
        <div class="divide-y">
          {#each submissions as sub}
            <div
              class="py-3 flex items-center justify-between hover:bg-gray-50 transition-colors"
            >
              <div>
                <p class="font-medium">{sub.time}</p>
                <p class="text-sm text-gray-600">状态: {sub.status}</p>
              </div>
            </div>
          {/each}
        </div>
      </div>
    </div>
  {/if}
</main>
