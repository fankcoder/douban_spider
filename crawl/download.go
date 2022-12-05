package crawl

import (
	"fmt"
	"os"
	"encoding/json"
	"io/ioutil"
	"strings"

	"douban_spider/util"
	"douban_spider/model"
)

type Img struct {
	Name string
	Url string
}

type Bookimg struct {
	Tag string `json:"tag"`
	Img []Img `json:"imgs"`
}

type BookDW struct {
	Book []Bookimg
}

// 判断文件夹是否存在
func PathExists(path string) (bool, error) {
    _, err := os.Stat(path)
    if err == nil {
        return true, nil
    }
    if os.IsNotExist(err) {
        return false, nil
    }
    return false, err
}

func Download() {
	tags := model.DB_fetch_tags()
	num := 0
	var book_dw BookDW
	for _, v := range *tags {
		fmt.Printf("tag data %s\n", v.Tag)
		data := model.DB_fetch_by_tag(v.Tag)
		if len(*data) == 0 {
			// fmt.Printf("no data: %s\n", v.Tag)
			num++
			continue
		}
		dir := fmt.Sprintf("./data")
		exist, err := PathExists(dir)
		util.HandleError(err, "path error")
		if exist {
			fmt.Printf("dir exists\n")
		} else {
			err := os.Mkdir(dir, os.ModePerm)
			util.HandleError(err, "create path error")
		}
		var img_list []Img
		for _, v := range *data{
			_img := strings.Replace(v.Image, "/s/", "/l/", 1)
			var img Img
			img.Name = v.Name
			img.Url = _img
			img_list = append(img_list, img)
		}
		fmt.Printf("img len %d\n", len(img_list))
		var book_img Bookimg
		book_img.Tag = v.Tag
		book_img.Img = img_list
		book_dw.Book = append(book_dw.Book, book_img)

	}
	book_data, _ := json.Marshal(book_dw)
	err := ioutil.WriteFile("./data/data.json", book_data, 0644)
	util.HandleError(err, "create data error")
}