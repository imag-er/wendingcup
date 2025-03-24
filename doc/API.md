## 基础信息
服务器地址: localhost:8080（配置文件中指定）
最大请求体大小: 100 MB
请求 GET, POST

## 接口列表


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



所有需要认证的接口都必须在请求头中包含有效的 JWT token。
JWT token 通常在登录接口成功响应后获得，并在后续请求的 Authorization 头部中以 Bearer {token} 的格式发送。