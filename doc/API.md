## 基础信息
服务器地址: localhost:8080
最大请求体大小: 10 MB
请求 GET, POST


## prompt
基于我给出的API.md接口定义完成前端网页的开发
功能要求如下

- 基于vite的svelte 模板进行开发 , 不要使用sveltekit
- header栏显示登录状态，如果未登录显示登录按钮，如果已登录显示队伍名和退出按钮
- 点击登录按钮，跳转到登录页面，用户可使用team_id进行登录(/login),登录成功后跳转回首页，显示队伍名和退出按钮, 登录相关需求见接口定义
- 点击退出按钮，退出登录状态，跳转到首页
- 点击队伍名，跳转到队伍详情页面，显示队伍信息
- 队伍详情页面显示队伍信息和队伍成员列表(/auth/team/${team_id}),队伍提交列表(/auth/submit/${team_id})
- 在localstorage中存储token和team_id，刷新页面时自动登录
- submit页面，显示提交文件的表单，提交文件到指定队伍, 提交文件的相关需求见接口定义
- 提交文件后，显示提交成功的提示信息
- 提交文件后，跳转到队伍详情页面
- 排行榜页面，显示所有队伍的提交文件列表，后端返回的结果已按照成绩排序


样式表要求如下
- 使用tailwindcss进行样式设计
- 需要支持移动端和PC端
- 需要支持暗黑模式
- 需要有流畅简单的动效, 对用户的操作有明显的反馈
- 需要有合理的排版和配色, 现代化的设计风格
- 需要有合理的交互设计, 用户操作流程清晰明了


## 功能模块
- 注册/登录
  - 显示队伍信息
  - 需要校验手机号(11位) 学号(9位) 姓名(<=64字节)字段长度、非空
  - 队伍人数 1-2人

- 提交文件
  - 显示本队伍提交列表
  - 检查文件类型 (zip)
  
- 获取排行榜
  - 显示排行榜信息


## 接口列表
所有需要认证的接口都必须在请求头中包含有效的 JWT token。
JWT token 在登录接口成功响应后获得，并在后续请求的 Authorization 头部中以 Bearer {token} 的格式发送。


- 获取排行榜
/board GET
描述: 获取排行榜信息。
参数: 无
返回值(示例): 
```json
[
    {
        "team_id": "b705235b-a373-481c-9650-905b6a10ffb3",
        "file_upload_time": "2025-03-23 17:01:47",
        "judge_result_time": "2025-03-23 17:01:50",
        "score": 18.719585
    },
            // ...若干条记录
]
```

   

- Ping 接口
/ping GET
描述: 用于检查服务是否在线。
参数: 无
返回值(示例):
```json
{"message": "pong"}
```

- 登录
/login POST
描述: 用户登录，获取 JWT 认证 token。
参数:
form: uuid (必填): 队伍ID // 检查: 是否为36字符
返回值示例:
```json
{
    {"code": 200, "expire": "2025-03-23T18:01:53+08:00", "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NDI3MjQxMTMsIm9yaWdfaWF0IjoxNzQyNzIwNTEzfQ.qdjoYFGMmjO5N2v3lgOnBzzorXYENfDbTFW4PBCBRgc"}
}
```


- 注册
/register POST
描述: 用户注册新账户。
参数:
返回值(示例):

```json
{
    "team_id": "b705235b-a373-481c-9650-905b6a10ffb3", 
    "teamname": "旅团vxxqx",
    "players": [
        {"name": "沈kfbsq", "phonenumber": "64377324677", "student_id": "245110193"},
        {"name": "李qzq", "phonenumber": "64377324677", "student_id": "245110193"},
    ]
    // 检查: 36位id, 64字节name (utf8), 11位phonenumber, 9位student_id
}
```


- 登出
/logout POST
描述: 用户登出，使 token 失效。
参数: 无
需要 JWT
返回值(示例):无

- 获取队伍信息
/auth/team/:team_id GET
描述: 获取指定队伍的详细信息。
参数:
Param: team_id (必填): 队伍ID（路径参数）// 检查: 36位
需要 JWT
返回值(示例):
```json
{"team_id": "119a8904-6014-438a-aaee-94f51616aa25", "teamname": "战队opyox", "players": [{"name": "郑isqss", "phonenumber": "72655224593", "student_id": "122026793"}, {"name": "吴xpjhu", "phonenumber": "69631735551", "student_id": "542851011"}]}
```

- 获取提交列表
/auth/submit/:team_id GET
描述: 获取指定队伍的提交列表。
参数:
Param: team_id (必填): 队伍ID（路径参数）// 检查: 36位
需要 JWT
返回值(示例):
```json
{"submit_list": [{"team_id": "119a8904-6014-438a-aaee-94f51616aa25", "time": "2025-03-23 17:01:47.915 +0800 CST", "status": "处理完成"}, {"team_id": "119a8904-6014-438a-aaee-94f51616aa25", "time": "2025-03-23 17:01:48.073 +0800 CST", "status": "处理完成"},...]}

```


- 提交文件
/auth/submit POST
描述: 提交文件到指定队伍。
参数:
form:"team_id" // 检查: 36位
form:"file" // 检查: 文件为zip类型, 大小不超过10MB
需要JWT
返回值(示例):
```json
{"id": "8", "message": "上传成功"}
```



- 刷新 Token
/auth/refresh POST
描述: 使用旧 token 获取新的 JWT token。
参数: 无
JWT
返回值(示例):
```json
{"code": 200, "expire": "2025-03-24T20:45:55+08:00", "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NDI4MjAzNTUsIm9yaWdfaWF0IjoxNzQyODE2NzU1fQ.TxdEE_idsxn8rX26byBOA2ZwcSS-0SlS9Szz-F5sHqA"}
```



