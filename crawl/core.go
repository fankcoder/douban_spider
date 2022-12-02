package crawl

import (
	"douban_spider/util"
	"douban_spider/model"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
	"strings"
	"strconv"
	"rand"

	// "regexp"

	"github.com/antchfx/htmlquery"
)

func getBookHtml(url string) string {
	client := &http.Client{}
	request, err := http.NewRequest("GET", url, nil)
	util.HandleError(err, "new req error")

	// 设置header属性
	request.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.212 Safari/537.36 Edg/90.0.818.62")
	request.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	request.Header.Add("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8,ru;q=0.7")
	request.Header.Add("Host", "book.douban.com")
	request.Header.Add("Referer", "https://book.douban.com/tag/?view=type&icn=index-sorttags-all")
	resp, _ := client.Do(request)
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Error status code: %d", resp.StatusCode)
	}
	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// panic(err)
		util.HandleError(err, "read error")
	}
	return fmt.Sprintf("%s", result)
}

func clean(html, tag string) {
	root, _ := htmlquery.Parse(strings.NewReader(html))
	var x_img = "//li[@class='subject-item']/div[@class='pic']/a[@class='nbg']/img/@src"
	img_res, _ := htmlquery.QueryAll(root, x_img)

	var x_name = "//li[@class='subject-item']/div[@class='info']/h2/a"
	name_res, _ := htmlquery.QueryAll(root, x_name)
	var x_author = "//li[@class='subject-item']/div[@class='info']/div[@class='pub']"
	aut_res, _ := htmlquery.QueryAll(root, x_author)
	var x_desc = "//ul[@class='subject-list']/li[@class='subject-item']/div[@class='info']/p"
	desc_res, _ := htmlquery.QueryAll(root, x_desc)
	var x_point = "//li[@class='subject-item']/div[@class='info']/div[@class='star clearfix']/span[@class='rating_nums']"
	point_res, _ := htmlquery.QueryAll(root, x_point)
	for i:=0;i<len(img_res);i++ {
		var book_data model.BookData
		book_data.Name = strings.TrimSpace(htmlquery.InnerText(name_res[i]))
		book_data.Author = strings.TrimSpace(htmlquery.InnerText(aut_res[i]))
		book_data.Describe = strings.TrimSpace(htmlquery.InnerText(desc_res[i]))
		book_data.Tag = tag
		book_data.Point = htmlquery.InnerText(point_res[i])
		book_data.Image = htmlquery.InnerText(img_res[i])
		// fmt.Printf("book ,%s, %s, %s, %s, %s,\n", book_data.Name, 
		// 			book_data.Author, book_data.Tag, book_data.Point, book_data.Image)
		model.DB_insert(&book_data)
	}
}

func GetBook() {
	// var tags *[]model.TagData
	tags := model.DB_fetch_tags()
	for _, v := range *tags {
		url := fmt.Sprintf("https://book.douban.com/tag/%s", v.Tag)
		// fmt.Printf("res data %s\n", url)
		html := getBookHtml(url)
		root, _ := htmlquery.Parse(strings.NewReader(html))
		// fmt.Printf("res data %v\n", root)

		var x_page = "//div[@id='subject_list']/div[@class='paginator']/a"
		page_res, _ := htmlquery.QueryAll(root, x_page)
		last_page := page_res[len(page_res)-1]
		last_page_int, err := strconv.Atoi(htmlquery.InnerText(last_page))
		util.HandleError(err, "read error")
		// fmt.Printf("last page %d",  last_page_int)
		for i:=2;i<last_page_int+1;i++ {
			param := fmt.Sprintf("?start=%d&type=T", 20*(i-1))
			fmt.Printf("sub param %s\n",  param)
			suburl := url + param
			fmt.Printf("sub url %s\n",  suburl)
			_html := getBookHtml(suburl)
			clean(_html, v.Tag)
			time.Sleep((rand.Intn(5) +1) * time.Second)
		}

		clean(html, v.Tag)
		time.Sleep((rand.Intn(10) +1) * time.Second)
	}
}