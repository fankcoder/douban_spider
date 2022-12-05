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

type Bookimg struct {
	Name string `json:"name"`
	Img []string `json:"img"`
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
		var img_list []string
		for _, v := range *data{
			_img := strings.Replace(v.Image, "/s/", "/l/", 1)
			img_list = append(img_list, _img)
		}
		fmt.Printf("img len %d\n", len(img_list))
		var book_img Bookimg
		book_img.Name = v.Tag
		book_img.Img = img_list

		book_data, _ := json.Marshal(book_img)

		err = ioutil.WriteFile(dir+"/data.json", book_data, 0644)
		util.HandleError(err, "create data error")
	}
}