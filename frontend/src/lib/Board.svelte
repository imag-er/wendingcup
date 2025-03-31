<script>
    import { ListGroup, ListGroupItem } from "@sveltestrap/sveltestrap";
    import { onMount } from "svelte";
    import { toastError, toastSuccess } from "./ToastMessage/ToastManager";

    let board = [];
    function getBoard() {
        let res = fetch("http://localhost:8080/board", {
            method: "GET",
        })
            .then((res) => res.json())
            .then((data) => {
                if (data === undefined) {
                    return;
                }
                board = data;
            });

        return res;
    }

    onMount(getBoard);
    setInterval(getBoard, 5000);
</script>

{#if board.length === 0}
    <p>loading...</p>
{:else}
    <ListGroup class="list-group-horizontal">
        <ListGroupItem class="list-group-item-success mb-1 col-md-3"
            >队伍名称</ListGroupItem
        >
        <ListGroupItem class="list-group-item-info mb-1 col-md-3"
            >队伍成绩</ListGroupItem
        >
        <ListGroupItem class="list-group-item-warning mb-1 col-md-3"
            >上传时间</ListGroupItem
        >
        <ListGroupItem class="list-group-item-warning mb-1 col-md-3"
            >评测结束时间</ListGroupItem
        >
    </ListGroup>
    {#each board as item}
        <!-- 为前三个添加primary标签 -->

        <ListGroup class="list-group-horizontal">
            <ListGroupItem
                class="list-group-item-{item.team_name ==
                localStorage.getItem('team_name')
                    ? 'active'
                    : 'success'} mb-1 col-md-3">{item.team_name}</ListGroupItem
            >
            <ListGroupItem class="list-group-item-info mb-1 col-md-3 "
                >{item.score}</ListGroupItem
            >
            <ListGroupItem class="list-group-item-warning mb-1 col-md-3"
                >{item.file_upload_time}</ListGroupItem
            >
            <ListGroupItem class="list-group-item-warning mb-1 col-md-3"
                >{item.judge_result_time}</ListGroupItem
            >
        </ListGroup>
    {/each}
{/if}
