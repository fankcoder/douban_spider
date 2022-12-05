import requests
import json
import os
import shutil
import time
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

def request_download(tag, name, url):
    import requests
    r = requests.get(url, headers=headers, stream=True)
    _file = './{}/{}.jpg'.format(tag, name)

    with open(_file, 'wb') as out_file:
        shutil.copyfileobj(BytesIO(r.content), out_file)

for item in data.get("Book"):
    tag = item.get("tag")
    folder = os.path.exists(tag)
    if not folder:
        os.makedirs(tag)     

    imgs = item.get("imgs")
    for i in imgs:
        print(tag, i.get("Name"), i.get("Url"))
        request_download(tag, i.get("Name"), i.get("Url"))
        time.sleep(3)