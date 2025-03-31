<script>
    import {
        Button,
        Form,
        Modal,
        ModalBody,
        ModalFooter,
    } from "@sveltestrap/sveltestrap";
    import { tick } from "svelte";
    import { toastError, toastSuccess } from "./ToastMessage/ToastManager";

    function autofocus(node) {
        tick();
        node.focus();
    }

    export let RegisterOpen = false;
    export function toggle() {
        RegisterOpen = !RegisterOpen;
        // 重置表单状态
        if (!RegisterOpen) {
            teamName = "";
            players = [{ student_id: "", name: "", phonenumber: "" }];
            errorMessage = "";
        }
    }

    let teamName = "tteam";
    let players = [
        { student_id: "220110101", name: "lsm", phonenumber: "13812341234" },
    ];
    let errorMessage = "";

    function handleRegister() {
        errorMessage = "";

        // 验证队伍名
        if (!teamName.trim()) {
            errorMessage = "请输入队伍名称";
            return;
        }

        // 验证成员数量
        if (players.length < 1 || players.length > 2) {
            errorMessage = "队伍成员必须为1或2人";
            return;
        }

        // 验证成员信息
        for (const [index, member] of players.entries()) {
            // 非空检查
            if (
                !member.student_id.trim() ||
                !member.name.trim() ||
                !member.phonenumber.trim()
            ) {
                errorMessage = `请填写第${index + 1}位成员的完整信息`;
                return;
            }

            // 学号校验（2开头9位数字）
            if (!/^2\d{8}$/.test(member.student_id)) {
                errorMessage = `第${index + 1}位成员的学号格式不正确（必须2开头的9位数字）`;
                return;
            }

            // 手机号校验（1开头11位数字）
            if (!/^1\d{10}$/.test(member.phonenumber)) {
                errorMessage = `第${index + 1}位成员的手机号格式不正确（必须1开头的11位数字）`;
                return;
            }

            // 姓名校验（UTF-8字节数 < 64）
            const nameByteLength = new TextEncoder().encode(member.name).length;
            if (nameByteLength >= 64) {
                errorMessage = `第${index + 1}位成员的姓名过长（超过64字节）`;
                return;
            }
        }

        fetch("http://localhost:8080/register", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify({
                team_name: teamName,
                players: players,
            }),
        })
            .then((res) => res.json())
            .then((data) => {
                toastSuccess("注册成功", "队伍注册成功");
                window.location.reload();

                toggle();
            })
            .catch((error) => {
                console.error("Error:", error);
                toastError("注册失败", "队伍注册失败");
            });
    }

    function addMember() {
        if (players.length < 2) {
            players = [
                ...players,
                { student_id: "", name: "", phonenumber: "" },
            ];
        }
    }
</script>

<Modal header="队伍注册" isOpen={RegisterOpen}>
    <ModalBody>
        <Form class="row g-3" on:submit={handleRegister}>
            <div class="mb-3">
                <input
                    type="text"
                    class="form-control"
                    placeholder="队伍名称"
                    bind:value={teamName}
                    required
                />
            </div>

            {#each players as member, index}
                <div class="member-group mb-3">
                    <div class="text-secondary small mb-1">
                        成员 {index + 1}
                    </div>
                    <input
                        type="text"
                        class="form-control mb-2"
                        placeholder="220101010"
                        bind:value={member.student_id}
                        maxlength="9"
                        inputmode="numeric"
                        pattern="[0-9]*"
                    />
                    <input
                        type="tel"
                        class="form-control mb-2"
                        placeholder="13812345678"
                        bind:value={member.phonenumber}
                        maxlength="11"
                        inputmode="numeric"
                        pattern="[0-9]*"
                    />
                    <input
                        type="text"
                        class="form-control mb-2"
                        placeholder="张三"
                        bind:value={member.name}
                    />
                </div>
            {/each}

            {#if errorMessage}
                <div class="text-danger small mb-3">{errorMessage}</div>
            {/if}

            {#if players.length < 2}
                <button
                    type="button"
                    class="btn btn-outline-secondary btn-sm"
                    on:click={addMember}
                >
                    添加成员 +
                </button>
            {/if}
        </Form>
    </ModalBody>
    <ModalFooter>
        <Button color="primary" on:click={handleRegister}>注册</Button>
        <Button on:click={toggle}>取消</Button>
    </ModalFooter>
</Modal>
