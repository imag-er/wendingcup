import requests
from requests.exceptions import RequestException

token = open('token.txt','r').read().strip()

# 构造表单数据
data = {
    'team_id': 'fc4a7d0e-bb15-489c-b008-435069fb5c6d'
}
files = {
    'file': ('example.txt', open('example.txt', 'rb'))  # 使用元组格式 (filename, fileobj)
}

try:
    response = requests.post(
        'http://localhost:8080/submit',
        data=data,
        files=files,  # 现在包含正确的文件名和文件对象
        headers={
            'Authorization': 'Bearer ' + token
        }
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