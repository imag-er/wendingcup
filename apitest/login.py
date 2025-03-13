# fc4a7d0e-bb15-489c-b008-435069fb5c6d

import requests
from requests.exceptions import RequestException

response = requests.post(
    'http://localhost:8080/auth/login',
    data={'team_id': 'fc4a7d0e-bb15-489c-b008-435069fb5c6d'},
    headers={'Content-Type': 'application/x-www-form-urlencoded'},
)

if response.status_code == 200:
    print('登录成功:', response.json())

token = response.json()['token']
print(token,file=open('token.txt','w+'))

