import re
from urllib import error
import requests
from bs4 import BeautifulSoup
import json


# 根据地址去查找 对应的图片的信息
def Find(url, A):
    global List  # 保存信息的列表
    print('正在检测图片总数，请稍等.....')
    t = 0
    s = 0
    while t < 1000:
        # 时间戳 不简单刷新访问网址
        Url = url + str(t)
        try:
            # get获取数据
            Result = A.get(Url, timeout=7, allow_redirects=False)
        except BaseException:
            t = t + 60
            continue
        else:
            # 拿到网站的数据
            result = Result.text
            # 找到图片url
            pic_url = re.findall('"objURL":"(.*?)",', result, re.S)
            # 图片总数
            s += len(pic_url)
            if len(pic_url) == 0:
                break
            else:
                List.append(pic_url)
                t = t + 60
    return s


# 记录相关数据
def recommend(url):
    Re = []
    try:
        html = requests.get(url, allow_redirects=False)
    except error.HTTPError as e:
        return e
    else:
        html.encoding = 'utf-8'
        # html文件解析
        bsObj = BeautifulSoup(html.text, 'html.parser')
        imgDivList = bsObj.find_all("div", attrs={'class': 'txList'})

        for imgDiv in imgDivList:
            imgUrl = "https:"+imgDiv.find("img", attrs={'class': 'lazy'}).get("src")
            imgTitle = imgDiv.find("a", attrs={'class': 'imgTitle'}).get("title")
            Re.append({"url": imgUrl, "title": imgTitle})
        return Re


# 下载图片


# def dowmloadPicture(html, keyword):
#     global num
#     # 找到图片url
#     pic_url = re.findall('"objURL":"(.*?)",', html, re.S)
#     print('找到第:' + keyword + '页的图片，开始下载图片....')
#     for each in pic_url:
#         print('正在下载第' + str(num + 1) + '张图片，图片地址:' + str(each))
#         try:
#             if each is not None:
#                 pic = requests.get(each, timeout=7)
#             else:
#                 continue
#         except BaseException:
#             print('错误，当前图片无法下载')
#             continue
#         else:
#             string = file + r'\\' + str(num) + '.jpg'
#             fp = open(string, 'wb')
#             fp.write(pic.content)
#             fp.close()
#             num += 1
#         if num >= numPicture:
#             return

def write_json(filename, list):
    with open(filename, 'a+', encoding='utf-8') as f:
        json.dump(list, f, ensure_ascii=False)


def craw_img(pageIndex):
    headers = {
        'Accept-Language': 'zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6',
        'Connection': 'keep-alive',
        'User-Agent':
        'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36 Edg/110.0.1587.41',
        'Upgrade-Insecure-Requests': '1'
    }

    # 创建一个请求的会话
    A = requests.Session()
    # 设置头部信息
    A.headers = headers

    if pageIndex == '1':
        pageIndex = ''
    # 拼接路径
    url = 'https://www.woyaogexing.com/touxiang/index_' + pageIndex + '.html'

    Recommend = recommend(url)
    # print('经过检测第%s页, 图片共有%d张' % (pageIndex, len(Recommend)))

    file = "img.json"
    print("将数据写入"+file+"文件中...")
    imgFile = open(file, encoding='utf-8', mode='a+')

    # write_json(file, Recommend)

    for item in Recommend:
        imgFile.write(item['url']+","+item['title']+"\n")
        print("url:", item['url'], "   title:", item['title'])
    return len(Recommend)


if __name__ == '__main__':  # 主函数入口
    craw_img('1')
    craw_img('2')
    craw_img('3')
    craw_img('4')
    craw_img('5')

    