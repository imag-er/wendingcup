# fc4a7d0e-bb15-489c-b008-435069fb5c6d

import requests
from requests.exceptions import RequestException

token = open('token.txt','r').read().strip()

print(token)
try:
    response = requests.get(
        'http://127.0.0.1:8080/ping',
        headers={
                'Authorization': 'Bearer ' + token},
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

# print(response.json())