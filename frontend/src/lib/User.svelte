<script>
    import {
        Button,
        Card,
        CardBody,
        CardHeader,
        CardText,
        CardTitle,
        Form,
    } from "@sveltestrap/sveltestrap";
    import { LoginStatusAvailable, CheckJWT, LogOut } from "./auth";
    import RegisterPopUp from "./RegisterPopUp.svelte";
    import { toastError, toastSuccess } from "./ToastMessage/ToastManager";
    let teamId = null;
    function login(e) {
        // 提交form data
        e.preventDefault();
        fetch(`http://localhost:8080/login`, {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify({
                team_id: teamId,
            }),
        })
            .then((res) => res.json())
            .then((data) => {
                if (data.error != undefined) {
                    toastError("登录失败", data.message);
                    return;
                }
                localStorage.setItem("team_id", teamId);
                localStorage.setItem("expire", data.expire);
                localStorage.setItem("token", data.token);
                getTeamInfo();
                window.location.reload();
            })
            .catch((error) => {
                console.error("Error:", error);
                toastError("消息", "登录失败");
            });
    }

    function getTeamInfo() {
        return fetch(
            `http://localhost:8080/auth/team/${localStorage.getItem("team_id")}`,
            {
                method: "GET",
                headers: {
                    Authorization: `Bearer ${localStorage.getItem("token")}`,
                },
            },
        )
            .then(CheckJWT)
            .then((res) => {
                if (res.error != undefined) {
                    toastError("注册失败", res.message);
                    return;
                }
                toastSuccess("消息", "获取队伍信息成功");
                localStorage.setItem("team_name", res.teamname);
                return res;
            });
    }

    let RegisterOpen = false;
</script>

<RegisterPopUp {RegisterOpen} />

{#if LoginStatusAvailable()}
    <Card class="text-bg-info">
        <CardHeader>队伍信息</CardHeader>
        <CardBody style="display: flex; flex-direction: column;">
            {#await getTeamInfo() then data}
                <CardTitle>{data.teamname}</CardTitle>
                {#each data.players as p}
                    <CardText>{p.student_id} {p.name}</CardText>
                {/each}
            {/await}
        </CardBody>
        <Button type="button" class="btn-danger" on:click={LogOut}>X</Button>
    </Card>
{:else}
    <Card class="text-bg-info p-1">
        <CardHeader>队伍信息</CardHeader>
        <CardBody style="display: flex; flex-direction: column;">
            <Form class="row g-3">
                <div class="row-auto m-0 p-0 b-0">
                    <input
                        type="text"
                        class="form-control mb-1"
                        id="team_id"
                        placeholder="Team ID"
                        bind:value={teamId}
                    />
                    <button
                        type="submit"
                        class="btn btn-primary mb-3"
                        on:click={login}>登录</button
                    >
                </div>
            </Form>
            <div class="col-auto">
                <Button
                    class="btn btn-primary mb-3"
                    on:click={() => {
                        RegisterOpen = !RegisterOpen;
                    }}
                >
                    注册
                </Button>
            </div>
        </CardBody>
    </Card>
{/if}
