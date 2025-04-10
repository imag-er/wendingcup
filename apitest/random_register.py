
import sys
import requests
from requests.exceptions import RequestException
import random
import string

# 生成随机队伍名称的函数


def generate_team_name():
    prefixes = ['赛博', '电子', '未来', '量子', '时空', '星河', '智慧',
                '探索', '战队', '先锋', '旅团', '小组', '联盟', '协会', '学派', '军团']
    name = random.choice(prefixes) + \
        ''.join(random.choices(string.ascii_lowercase, k=5))
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


# 从命令行参数读取队伍数量


data = [{
    'team_name': generate_team_name(),
    'players': [generate_player_info() for _ in range(random.randint(1, 2))]
} for _ in range(int(sys.argv[1]))
]

try:
    with open('team_id.txt', 'w') as file:
        for d in data:
            response = requests.post(
                'http://localhost:20000/register',
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
