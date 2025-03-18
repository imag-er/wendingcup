import requests
from requests.exceptions import RequestException

# 构造表单数据
data = [
    {
    'team_name': '赛博太上老君队',
    'players' : [
        {
            'name': '郭东',
            'phonenumber': '13800138000',
            'student_id': '220110102'
        },
        {
            'name': '陈奕',
            'phonenumber': '13900139000',
           'student_id': '220110101'
        },
    ]

},
{
    'team_name': '赛博炼丹炉',
    'players' : [
        {
            'name': '陈浩',
            'phonenumber': '13800138011',
            'student_id': '220110103'
        },
        {
            'name': '陈奕',
            'phonenumber': '13900139000',
           'student_id': '220110107'
        },
    ]

},
{
    'team_name': '东华东华',
    'players' : [
        {
            'name': '严谦',
            'phonenumber': '13800138000',
            'student_id': '220110110'
        },
        {
            'name': '唐志',
            'phonenumber': '13900139000',
           'student_id': '220110111'
        },
    ]

},
]

try:
    for d in data:
        response = requests.post(
            'http://localhost:8080/register',
            json=d,
            headers={'Content-Type': 'application/json'}
        )
        
        if response.status_code == 200:
            print('注册成功:', response.json())
            print(response.json()['team_id'],file=open('team_id.txt','w+'))
        else:
            print(f'请求失败，状态码：{response.status_code}')
            print('响应内容:', response.text)

except RequestException as e:
    print(f'请求异常：{str(e)}')
except Exception as e:
    print(f'发生错误：{str(e)}')