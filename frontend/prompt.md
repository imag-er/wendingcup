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