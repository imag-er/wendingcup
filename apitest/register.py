# import requests
# from requests.exceptions import RequestException

# # 构造表单数据
# data = [
#     {
#     'team_name': '赛博太上老君队',
#     'players' : [
#         {
#             'name': '郭东',
#             'phonenumber': '13800138000',
#             'student_id': '220110102'
#         },
#         {
#             'name': '陈奕',
#             'phonenumber': '13900139000',
#            'student_id': '220110101'
#         },
#     ]

# },
# {
#     'team_name': '赛博炼丹炉',
#     'players' : [
#         {
#             'name': '陈浩',
#             'phonenumber': '13800138011',
#             'student_id': '220110103'
#         },
#         {
#             'name': '陈奕',
#             'phonenumber': '13900139000',
#            'student_id': '220110107'
#         },
#     ]

# },
# {
#     'team_name': '东华东华',
#     'players' : [
#         {
#             'name': '严谦',
#             'phonenumber': '13800138000',
#             'student_id': '220110110'
#         },
#         {
#             'name': '唐志',
#             'phonenumber': '13900139000',
#            'student_id': '220110111'
#         },
#     ]

# },
# ]

# try:
#     for d in data:
#         response = requests.post(
#             'http://localhost:8080/register',
#             json=d,
#             headers={'Content-Type': 'application/json'}
#         )

#         if response.status_code == 200:
#             print('注册成功:', response.json())
#             print(response.json()['team_id'],file=open('team_id.txt','w+'))
#         else:
#             print(f'请求失败，状态码：{response.status_code}')
#             print('响应内容:', response.text)

# except RequestException as e:
#     print(f'请求异常：{str(e)}')
# except Exception as e:
#     print(f'发生错误：{str(e)}')


import requests
from requests.exceptions import RequestException
import random
import string

# 生成随机队伍名称的函数


def generate_team_name():
    prefixes = ['赛博', '电子', '未来', '量子', '时空', '星河', '智慧',
                '探索', '战队', '先锋', '旅团', '小组', '联盟', '协会', '学派', '军团']
    name = random.choice(prefixes) + ''.join(random.choices(string.ascii_lowercase, k=5))
    return name

# 生成随机玩家信息的函数


def generate_player_info():
    first_names = ['赵', '钱', '孙', '李', '周', '吴', '郑',
                   '王', '冯', '陈', '褚', '卫', '蒋', '沈', '韩', '杨']

    player = {
        'name': random.choice(first_names) + ''.join(random.choices(string.ascii_lowercase, k=5)),
        'phonenumber': ''.join(random.choices(string.digits, k=11)),
        'student_id': ''.join(random.choices(string.digits, k=9))
    }
    return player


# 生成1000个队伍的数据
data = [{'team_name': generate_team_name(), 'players': [generate_player_info()
                                                        for _ in range(random.randint(1,2))]} for _ in range(100)]

try:
    with open('team_id.txt', 'w') as file:
        for d in data:
            response = requests.post(
                'http://localhost:8080/register',
                json=d,
                headers={'Content-Type': 'application/json'}
            )

            if response.status_code == 200:
                print('注册成功:', response.json())
                file.write(response.json()['team_id'] + '\n')
            else:
                print(f'请求失败，状态码：{response.status_code}')
                print('响应内容:', response.text)

except RequestException as e:
    print(f'请求异常：{str(e)}')
except Exception as e:
    print(f'发生错误：{str(e)}')
