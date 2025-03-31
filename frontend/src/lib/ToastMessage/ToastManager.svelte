<script lang="ts">
    import { onMount } from "svelte";
    import { Toast, ToastBody, ToastHeader } from "@sveltestrap/sveltestrap";
    import { NlpToastMessage, toasts } from "./ToastManager";
    let toastMessages: NlpToastMessage[] = [];
  
    onMount(() => {
      return toasts.subscribe((_entries: NlpToastMessage[]) => {
        toastMessages = _entries;
      });
    });
  </script>
  
  <div class="toast-container position-fixed bottom-0 end-0 p-3">
    {#each toastMessages as entry}
    <Toast autohide>
      <ToastHeader icon="{entry.color}" toggle="{entry.dismiss}">
        {entry.title}
      </ToastHeader>
      {#if entry.body}<ToastBody>{entry.body}</ToastBody>{/if}
    </Toast>
    {/each}
  </div>