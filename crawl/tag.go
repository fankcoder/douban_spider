package crawl

import (
	"douban_spider/util"
	"douban_spider/model"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	// "regexp"

	"github.com/antchfx/htmlquery"
)

var (
	xmatch = `//tbody/tr/td/a`
)

func getHtml() string {
	client := &http.Client{}
	url := "https://book.douban.com/tag/?view=type&icn=index-sorttags-all"
	request, err := http.NewRequest("GET", url, nil)
	util.HandleError(err, "new req error")

	// 设置header属性
	request.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.212 Safari/537.36 Edg/90.0.818.62")
	request.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	request.Header.Add("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8,ru;q=0.7")
	request.Header.Add("Host", "book.douban.com")
	request.Header.Add("Referer", "https://book.douban.com/")
	fmt.Printf("2")
	resp, _ := client.Do(request)
	fmt.Printf("3")
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Error status code: %d", resp.StatusCode)
	}
	fmt.Printf("4")
	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// panic(err)
		util.HandleError(err, "read error")
	}
	fmt.Printf("5")
	return fmt.Sprintf("%s", result)
}


func GetTag() {
	fmt.Printf("1")
	html := getHtml()
	root, _ := htmlquery.Parse(strings.NewReader(html))
	fmt.Printf(html)
	// re := regexp.MustCompile(xmatch)
	res, _ := htmlquery.QueryAll(root, xmatch)
	// fmt.Printf("%d", len(res))
	var tags []model.TagData
	for _, item := range res {
		var tag model.TagData
		tag.Tag = htmlquery.InnerText(item)
		fmt.Printf("item %s \n", htmlquery.InnerText(item))
		tags = append(tags, tag)
	}
	model.DB_insert_tag(&tags)

}
