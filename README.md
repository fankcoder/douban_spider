# 豆瓣爬虫
## 帮设计师朋友爬下书籍封面
#### 使用

先获取书籍分类

```
crawl.GetTag() 
```

再通过分类获取书籍信息

```
crawl.GetBook()
```

再通过书籍信息Image字段获得图片下载链接
```
crawl.Download()
```