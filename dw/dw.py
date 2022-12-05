import requests
import json
import os
import shutil
import time
import asyncio
from io import BytesIO

headers = {
    'User-Agent': 'Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.132 Safari/537.36 QIHU 360SE',
    "Accept": "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9",
	"Accept-Language": "zh-CN,zh;q=0.9,en;q=0.8,ru;q=0.7",
	"Host": "img2.doubanio.com",
}
f = open('./data.json')
data = json.load(f)
f.close()

Proxy = False #开启代理Ture/False

def request_download(tag, name, url):
    try:
        if Proxy:
            proxies = {
                'http': 'http://121.13.252.60:41564',
                'https': 'http://121.13.252.60:41564',
            }

            r = requests.get(url, headers=headers, stream=True, timeout=10, proxies=proxies)
        else:
            r = requests.get(url, headers=headers, stream=True, timeout=10)
        _file = './{}/{}.jpg'.format(tag, name)

        with open(_file, 'wb') as out_file:
            shutil.copyfileobj(BytesIO(r.content), out_file)
    except:
        return

async def main():
    num = 0
    for item in data.get("Book"):
        tag = item.get("tag")
        folder = os.path.exists(tag)
        if not folder:
            os.makedirs(tag)     

        imgs = item.get("imgs")
        for i in imgs:
            # print(tag, i.get("Name"), i.get("Url"))
            print('正在下载第{}个...'.format(num))
            num += 1
            request_download(tag, i.get("Name"), i.get("Url"))

if __name__ == "__main__":
    coroutine = main()
    loop = asyncio.get_event_loop()
    loop.run_until_complete(coroutine) # 将协程对象添加到事件循环，运行直到结束
    print(os.getpid())
    loop.close() # 关闭事件循环