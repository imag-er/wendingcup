# fc4a7d0e-bb15-489c-b008-435069fb5c6d

import random
import requests
from requests.exceptions import RequestException


response = requests.get('http://localhost:8080/ping')
print(response.json())


team_ids = [line[:-1] for line in open('team_id.txt').readlines()]
tid = random.choice(team_ids)


response = requests.post(
    'http://localhost:8080/login',
    data={'team_id': tid},
    headers={'Content-Type': 'application/x-www-form-urlencoded'},
)
print('登录:', response.json())
token = response.json()['token']


response = requests.get(
    'http://localhost:8080/auth/team/' + tid,
    headers={'Content-Type': 'application/x-www-form-urlencoded',
             'Authorization': f'Bearer {token}'},
)
print('获取队伍信息:', response.json())


response = requests.get(
    'http://localhost:8080/board',
    data={'team_id': tid},
    headers={'Content-Type': 'application/x-www-form-urlencoded',
             'Authorization': f'Bearer {token}'},
)
print('获取排行榜:', response.json())


response = requests.get(
    'http://localhost:8080/auth/submit/' + tid,
    data={'team_id': tid},
    headers={'Content-Type': 'application/x-www-form-urlencoded',
             'Authorization': f'Bearer {token}'},
)
print('获取提交记录:', response.json())
