import re
import requests
from requests.exceptions import RequestException
import sys


def make_request(method, url, data=None, json=None, headers=None, files=None):
    try:
        response = requests.Response()
        if method == 'GET':
            response = requests.get(url, headers=headers)
        elif method == 'POST':
            response = requests.post(
                url, data=data, json=json, headers=headers, files=files)
        response.raise_for_status()

        return response.json()
    except RequestException as e:
        print("请求失败:", e)
        sys.exit(1)


print("pinging...\n==============")
ping_response = make_request('GET', 'http://localhost:8080/ping')
print(ping_response)

team_id = []
for id in open("team_id.txt", "r").readlines():
    id = id.strip()
    team_id.append(id)

print("登录中...\n==============")
login_response = make_request('POST', 'http://localhost:8080/login', data={
                              'team_id': team_id[0]}, headers={'Content-Type': 'application/x-www-form-urlencoded'})
print('登录成功:', login_response)
print("登录成功\n==============")

token = login_response['token']

# auth.GET("/teaminfo", handler.TeamInfoHandler)
# auth.POST("/submit", infra.AuthMiddleware.MiddlewareFunc(), handler.SubmitHandler)
# auth.POST("/refresh", infra.AuthMiddleware.RefreshHandler)
print("获取队伍信息...\n==============")
teaminfo_response = make_request(
    'GET', f'http://localhost:8080/auth/teaminfo?team_id={team_id[0]}',
    headers={'Authorization': f'Bearer {token}'},)
print('获取队伍信息', teaminfo_response)

print("提交...\n==============")
submit_response = make_request('POST', 'http://localhost:8080/auth/submit',
                               headers={'Authorization': f'Bearer {token}'},
                               data={'team_id': team_id[0]},
                               files={
                                   'file': ('somename.zip', open('somename.zip', 'rb'))}
                               )
print(submit_response)

print("刷新token...\n==============")
refresh_response = make_request('POST', 'http://localhost:8080/auth/refresh',
                                headers={'Authorization': f'Bearer {token}'},
                                data={'team_id': team_id[0]})
token = refresh_response['token']
print(refresh_response)
