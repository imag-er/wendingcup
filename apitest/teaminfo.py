# fc4a7d0e-bb15-489c-b008-435069fb5c6d

import requests
from requests.exceptions import RequestException

response = requests.get(
    'http://localhost:8080/teaminfo',
    data={'team_id': 'fc4a7d0e-bb15-489c-b008-435069fb5c6d'},
    headers={'Content-Type': 'application/x-www-form-urlencoded'},
)



print('登录:', response.json())
token = response.json()['token']
print(token,file=open('token.txt','w+'))

