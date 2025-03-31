

function LoginStatusAvailable() {
    if (localStorage.getItem("team_id") === null) {
        return false;
    }
    if (localStorage.getItem("token") === null) {
        return false;
    }

    // 判断token是否过期
    let expire = localStorage.getItem("expire");
    if (expire === null) {
        return false;
    }
    let expireDate = new Date(expire);
    let now = new Date();
    // 如果当前时间大于过期时间，返回false
    if (now.getTime() > expireDate.getTime()) {
        return false;
    }

    return true;
}
function clearLocalStorage() {
    localStorage.removeItem("team_id");
    localStorage.removeItem("expire");
    localStorage.removeItem("token");
    localStorage.removeItem("team_name");
}
function CheckJWT(res) {
    if (!res.ok) {
        // 捕获 401 未授权状态码
        if (res.status === 401) {
            clearLocalStorage();
            window.location.reload(); // 刷新页面以显示登录表单
        }
        throw new Error(`HTTP error! status: ${res.status}`);
    }
    return res.json();
}

function LogOut() {
    clearLocalStorage();
    window.location.reload(); // 刷新页面以显示登录表单
}

function register() {
    // TODO
}

export { LoginStatusAvailable,CheckJWT,LogOut };