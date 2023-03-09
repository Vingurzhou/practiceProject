import pandas
import requests


def crawl(i):
    pageNum = i
    pageSize = 10
    url = f"http://c.shidastudy.com/api/SHIDA-GOODSORDER-SERVICE/OrdersApi/findList?pageSize={pageSize}&pageNum={pageNum}"

    payload = {}
    headers = {
        'Connection': 'keep-alive',
        'Accept': 'application/json, text/plain, */*',
        'Origin': 'http://120.79.13.151:9000',
        'X-Authorization': 'eyJhbGciOiJIUzUxMiJ9.eyJzdWIiOiJCQUNLU1RBR0VfVVNFUiIsInJlZGlzS2V5IjoiU1NPOkJBQ0tTVEFHRV9VU0VSOjMzNDY3MWVhZjk4NzRkNTlhOTdiZmIzYTAxZmVjYmRlIiwiaXNBZG1pbiI6ZmFsc2UsImV4cCI6MTY3ODg5NTkxMCwib3BlcmF0b3JJZCI6IjMzNDY3MWVhZjk4NzRkNTlhOTdiZmIzYTAxZmVjYmRlIiwiaWF0IjoxNjc4MjkxMTEwfQ.bv-ERby7tOo9bPjWGY8YXlduMrxs2lsSVAlpTHkkddeqUxXnFmAaBxCEl8ZJdPkCYZgwbS1dI5peelDgyYrp5w',
        'jwtToken': 'eyJhbGciOiJIUzUxMiJ9.eyJzdWIiOiJCQUNLU1RBR0VfVVNFUiIsInJlZGlzS2V5IjoiU1NPOkJBQ0tTVEFHRV9VU0VSOjMzNDY3MWVhZjk4NzRkNTlhOTdiZmIzYTAxZmVjYmRlIiwiaXNBZG1pbiI6ZmFsc2UsImV4cCI6MTY3ODg5NTkxMCwib3BlcmF0b3JJZCI6IjMzNDY3MWVhZjk4NzRkNTlhOTdiZmIzYTAxZmVjYmRlIiwiaWF0IjoxNjc4MjkxMTEwfQ.bv-ERby7tOo9bPjWGY8YXlduMrxs2lsSVAlpTHkkddeqUxXnFmAaBxCEl8ZJdPkCYZgwbS1dI5peelDgyYrp5w',
        'User-Agent': 'Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.87 Safari/537.36 SE 2.X MetaSr 1.0',
        'Referer': 'http://120.79.13.151:9000/',
        'Accept-Language': 'zh-CN,zh;q=0.9'
    }

    response = requests.request("GET", url, headers=headers, data=payload)
    # print(response.json())
    for j in response.json()["content"]["list"]:
        url = f"http://c.shidastudy.com/api/SHIDA-GOODSORDER-SERVICE/OrdersApi/get/{j['id']}"

        payload = {}
        headers = {
            'Connection': 'keep-alive',
            'Accept': 'application/json, text/plain, */*',
            'Origin': 'http://120.79.13.151:9000',
            'X-Authorization': 'eyJhbGciOiJIUzUxMiJ9.eyJzdWIiOiJCQUNLU1RBR0VfVVNFUiIsInJlZGlzS2V5IjoiU1NPOkJBQ0tTVEFHRV9VU0VSOjMzNDY3MWVhZjk4NzRkNTlhOTdiZmIzYTAxZmVjYmRlIiwiaXNBZG1pbiI6ZmFsc2UsImV4cCI6MTY3ODg5NTkxMCwib3BlcmF0b3JJZCI6IjMzNDY3MWVhZjk4NzRkNTlhOTdiZmIzYTAxZmVjYmRlIiwiaWF0IjoxNjc4MjkxMTEwfQ.bv-ERby7tOo9bPjWGY8YXlduMrxs2lsSVAlpTHkkddeqUxXnFmAaBxCEl8ZJdPkCYZgwbS1dI5peelDgyYrp5w',
            'jwtToken': 'eyJhbGciOiJIUzUxMiJ9.eyJzdWIiOiJCQUNLU1RBR0VfVVNFUiIsInJlZGlzS2V5IjoiU1NPOkJBQ0tTVEFHRV9VU0VSOjMzNDY3MWVhZjk4NzRkNTlhOTdiZmIzYTAxZmVjYmRlIiwiaXNBZG1pbiI6ZmFsc2UsImV4cCI6MTY3ODg5NTkxMCwib3BlcmF0b3JJZCI6IjMzNDY3MWVhZjk4NzRkNTlhOTdiZmIzYTAxZmVjYmRlIiwiaWF0IjoxNjc4MjkxMTEwfQ.bv-ERby7tOo9bPjWGY8YXlduMrxs2lsSVAlpTHkkddeqUxXnFmAaBxCEl8ZJdPkCYZgwbS1dI5peelDgyYrp5w',
            'User-Agent': 'Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.87 Safari/537.36 SE 2.X MetaSr 1.0',
            'Referer': 'http://120.79.13.151:9000/',
            'Accept-Language': 'zh-CN,zh;q=0.9'
        }

        response = requests.request("GET", url, headers=headers, data=payload)
        # print(response.json())
        r = response.json()["content"]
        lst.append(r)
        print(i)
        # print(
        #     f"""课程名字:{r['goodsName']},姓名:{r['userName']},电话:{r['mobile']},支付方式:{r['payTypeStr']},金额:{r['buyPrice']},订单时间:{r['createTime']}""")


if __name__ == '__main__':
    import concurrent.futures

    lst = []

    # 创建一个线程池
    with concurrent.futures.ThreadPoolExecutor(max_workers=16) as executor:
        # 提交要执行的任务
        for p in range(1, 86514 + 1):
            executor.submit(crawl, p)
    pandas.DataFrame(lst).to_csv('result.csv')
