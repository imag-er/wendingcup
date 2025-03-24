# 25 wendingcup


技术栈: cwgo(hertz kitex) prometheus opentelemetry mysql consul 


### 服务划分
- api
    - 提供api供前端调用, 分发jwt, rpc调用后端服务
    - [api文档](./doc/API.md)
- user
    用于用户的注册、登录, 以及相关认证
- submit
    用于文件提交与评测, 主要记录文件提交, 评测结果交由board, 
- board
    负责评测结果的排序与储存, 防止因评测机问题导致排行榜无法访问

