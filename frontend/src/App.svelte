<script>
  import Board from "./lib/Board.svelte";
  import Submit from "./lib/Submit.svelte";

  import { ListGroup, ListGroupItem } from "@sveltestrap/sveltestrap";
  import User from "./lib/User.svelte";
  import ToastManager from "./lib/ToastMessage/ToastManager.svelte";
</script>

<main style="display: flex;">
  <ToastManager />
  <div style="flex: 4;">
    <Board />
  </div>
  <div style="flex: 1;">
    <User />
    <Submit />
  </div>
</main>

<style>
</style>
