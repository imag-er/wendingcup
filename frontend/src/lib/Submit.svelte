<script>
    import {
        Button,
        Card,
        CardBody,
        CardHeader,
    } from "@sveltestrap/sveltestrap";
    import { LoginStatusAvailable } from "./auth";
    import { CheckJWT } from "./auth";
    import { toastError, toastSuccess } from "./ToastMessage/ToastManager";
    
    function getSubmitList() {
        return fetch(
            `http://localhost:8080/auth/submit/${localStorage.getItem("team_id")}`,
            {
                method: "GET",
                headers: {
                    Authorization: `Bearer ${localStorage.getItem("token")}`,
                },
            },
        )
            .then(CheckJWT)
            .then((res) => {
                toastSuccess("消息", "获取提交列表成功");
                res.submit_list = res.submit_list.reverse();
                return res;
            });
    }

    let cardColor = {
        处理完成: "lightgreen",
        判题中: "lightyellow",
        处理失败: "lightcoral",
        已上传: "lightblue",
    };

    let files;
    function onClick() {
        if (files.length === 0) {
            return;
        }
        if (files[0] === undefined) {
            return;
        }

        let formData = new FormData();
        formData.append("file", files[0]);
        formData.append("team_id", localStorage.getItem("team_id"));

        fetch(`http://localhost:8080/auth/submit`, {
            method: "POST",
            headers: {
                Authorization: `Bearer ${localStorage.getItem("token")}`,
            },
            body: formData,
        }).then((res) => {
            if (res.status === 200) {
                toastSuccess("消息", "提交成功");
            } else {
                toastError("消息", "提交失败");
            }
        });

        // 刷新当前组件
        getSubmitList();
    }
</script>

<Card class="text-bg-success mb-3">
    <CardHeader>提交</CardHeader>
    {#if LoginStatusAvailable()}
        <div class="mb-3">
            <input
                class="form-control form-control-sm"
                id="formFileSm"
                type="file"
                accept=".zip"
                bind:files
            />
            <Button class="mt-3" color="primary" on:click={onClick}
                >Submit</Button
            >
        </div>
        {#await getSubmitList() then data}
            <ul style="list-style: none; padding: 0;">
                {#each data.submit_list as item}
                    <Card
                        class="mb-3"
                        style="background-color: {cardColor[item.status]};"
                    >
                        <CardBody>
                            <p>{item.status}</p>
                            <p>{item.time}</p>
                            {#if item.score !== undefined}
                                <p>{item.score}</p>
                            {/if}
                        </CardBody>
                    </Card>
                {/each}
            </ul>
        {/await}
        <CardBody style="display: flex; flex-direction: column;"></CardBody>
    {:else}
        <CardBody style="display: flex; flex-direction: column;">
            <p>请登录</p>
        </CardBody>
    {/if}
</Card>
