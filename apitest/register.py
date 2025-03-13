import requests
from requests.exceptions import RequestException

# 构造表单数据
data = {
    'team_name': '示例战队',
    'players[0].name': '张三',
    'players[0].phonenumber': '13800138000',
    'players[0].student_id': '20210001',
    'players[1].name': '李四',
    'players[1].phonenumber': '13900139000',
    'players[1].student_id': '20210002'
}

try:
    response = requests.post(
        'http://localhost:8080/auth/register',
        data=data,
        headers={'Content-Type': 'application/x-www-form-urlencoded'}
    )
    
    if response.status_code == 200:
        print('注册成功:', response.json())
    else:
        print(f'请求失败，状态码：{response.status_code}')
        print('响应内容:', response.text)

except RequestException as e:
    print(f'请求异常：{str(e)}')
except Exception as e:
    print(f'发生错误：{str(e)}')