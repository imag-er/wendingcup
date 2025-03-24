# fc4a7d0e-bb15-489c-b008-435069fb5c6d

import sys
import requests
from requests.exceptions import RequestException


def make_request(url, data=None, json=None, headers=None, files=None):
    print(f"{data=}")
    try:
        response = requests.post(
            url, data=data, json=json, headers=headers, files=files)
        response.raise_for_status()
        return response.json()
    except RequestException as e:
        print("请求失败:", e)
        sys.exit(1)


# 构造表单数据
team_ids = [
    line[:-1] for line in open('team_id.txt').readlines()
]

tokens = [
    make_request('http://localhost:8080/login',
                 data={'team_id': tid},
                 headers={'Content-Type': 'application/x-www-form-urlencoded'})['token']
    for tid in team_ids
]

print(tokens)
