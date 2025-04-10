<script>
  import Board from "./lib/Board.svelte";
  import Submit from "./lib/Submit.svelte";
  import { ListGroup, ListGroupItem } from "@sveltestrap/sveltestrap";
  import User from "./lib/User.svelte";
  import ToastManager from "./lib/ToastMessage/ToastManager.svelte";
</script>

<main style="display: flex; height: 100vh; margin: 0; padding: 0; width: 100%;">
  <ToastManager />
  <div style="flex: 8; padding: 10px;">
    <Board />
  </div>
  <div style="flex: 4; display: flex; flex-direction: column;">
      <User />
      <Submit />
  </div>
</main>

<style>

  /* main标签样式 */
  main {
    display: flex;
    flex: 1;
    width: 100%;
    height: 100%;
  }

  /* 左侧Board组件样式 */
  main > div:first-child {
    flex: 8;
    padding: 10px;
    overflow-y: auto; /* 如果内容过多，则显示滚动条 */
  }

  /* 右侧User和Submit组件容器样式 */
  main > div:last-child {
    flex: 4;
    display: flex;
    flex-direction: column;
  }


  /* 响应式设计 */
  @media (max-width: 768px) {
    main {
      flex-direction: column;
    }

    main > div:first-child,


    main > div:last-child {
      flex-direction: row;
    }
  }
</style>
