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
        fetch(`http://123.207.207.136:20000/login`, {
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
            `http://123.207.207.136:20000/auth/team/${localStorage.getItem("team_id")}`,
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
    <Card class="">
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
    <Card class="p-1">
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

<style global>
    /* 全局基础样式 */
    :global(body) {
        background-color: #f8f9fa;
        font-family: 'Segoe UI', system-ui, sans-serif;
    }

    /* 卡片容器样式 */
    :global(.card) {
        border: none;
        border-radius: 12px;
        box-shadow: 0 4px 6px rgba(0, 0, 0, 0.05);
        transition: transform 0.2s ease, box-shadow 0.2s ease;
        background: white;
        overflow: hidden;
    }

    :global(.card:hover) {
        transform: translateY(-2px);
        box-shadow: 0 8px 12px rgba(0, 0, 0, 0.1);
    }

    /* 卡片头部样式 */
    :global(.card-header) {
        background-color: #4a90e2;
        color: white;
        font-weight: 600;
        padding: 1rem 1.5rem;
        border-bottom: 2px solid rgba(255,255,255,0.1);
    }

    /* 输入框样式 */
    :global(.form-control) {
        border: 2px solid #e9ecef;
        border-radius: 8px;
        padding: 0.75rem 1rem;
        transition: all 0.2s ease;
    }

    :global(.form-control:focus) {
        border-color: #4a90e2;
        box-shadow: 0 0 0 3px rgba(74, 144, 226, 0.1);
    }

    /* 按钮基础样式 */
    :global(.btn) {
        border-radius: 8px;
        padding: 0.75rem 1.25rem;
        font-weight: 500;
        transition: all 0.2s ease;
        transform-style: preserve-3d;
        border: none;
    }

    :global(.btn-primary) {
        background-color: #4a90e2;
        color: white;
    }

    :global(.btn-primary:hover) {
        background-color: #357abd;
        transform: translateY(-1px);
    }

    :global(.btn-primary:active) {
        transform: translateY(0);
    }

    :global(.btn-danger) {
        background-color: #ff4757;
        color: white;
        padding: 0.5rem 0.75rem;
        position: absolute;
        right: 1rem;
        top: 1rem;
    }

    /* 信息展示区域 */
    :global(.card-title) {
        color: #2d3436;
        font-size: 1.25rem;
        margin-bottom: 1rem;
        font-weight: 600;
    }

    :global(.card-text) {
        color: #636e72;
        padding: 0.5rem 0;
        border-bottom: 1px solid #f1f2f6;
    }

    /* 布局优化 */
    :global(.row-auto) {
        display: flex;
        flex-direction: column;
        gap: 0.75rem;
    }

    /* 动画关键帧 */
    @keyframes fadeIn {
        from { opacity: 0; transform: translateY(10px); }
        to { opacity: 1; transform: translateY(0); }
    }

    :global(.card-body) {
        animation: fadeIn 0.3s ease-out;
        padding: 1.5rem;
    }

    /* 响应式处理 */
    @media (max-width: 576px) {
        :global(.card) {
            margin: 0.5rem;
        }
        
        :global(.card-header) {
            font-size: 1rem;
        }
    }
</style>
