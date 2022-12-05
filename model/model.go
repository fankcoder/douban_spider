package model

import (
	"douban_spider/util"
	"context"
	"fmt"

	"github.com/qiniu/qmgo"
	"go.mongodb.org/mongo-driver/bson"
)

type TagData struct {
	Tag string `bson:"tag"`
}

type BookData struct {
	Name string `bson:"name"`
	Author string `bson:"author"`
	Describe string `bson:"desc"`
	Tag string `bson:"tag"`
	Point string `bson:"point"`
	Image string `bson:"image"`
}


var db_name string = "douban"
var tagcol string = "tag"
var bookcol string = "book"
var uri string = "mongodb://localhost:27017"

func DB_insert(BookData *BookData) {
	ctx := context.Background()
	cli, err := qmgo.Open(ctx, &qmgo.Config{Uri: uri, Database: db_name, Coll: bookcol})
	util.HandleError(err, "connection error")

	defer func() {
		if err = cli.Close(ctx); err != nil {
			panic(err)
		}
	}()


	name := DB_find(BookData.Name)
	if name == "" {
		_, err := cli.InsertOne(ctx, BookData)
		util.HandleError(err, "insert ex listing error")
		fmt.Printf("insert %s\n", BookData.Name)
	} else {
		fmt.Printf("exists %s\n", name)
	}
}

func DB_find(name string) string {
	ctx := context.Background()
	cli, err := qmgo.Open(ctx, &qmgo.Config{Uri: uri, Database: db_name, Coll: bookcol})
	util.HandleError(err, "connection error")

	defer func() {
		if err = cli.Close(ctx); err != nil {
			panic(err)
		}
	}()

	one := BookData{}
	err = cli.Find(ctx, bson.M{"name": name}).One(&one)
	// fmt.Printf("res data %s", one.Name)
	return one.Name
}

func DB_fetch_by_tag(name string) *[]BookData {
	ctx := context.Background()
	cli, err := qmgo.Open(ctx, &qmgo.Config{Uri: uri, Database: db_name, Coll: bookcol})
	util.HandleError(err, "connection error")

	defer func() {
		if err = cli.Close(ctx); err != nil {
			panic(err)
		}
	}()

	list := []BookData{}
	err = cli.Find(ctx, bson.M{"tag": name}).All(&list)
	return &list
}


func DB_insert_tag(Tags *[]TagData) {
	ctx := context.Background()
	cli, err := qmgo.Open(ctx, &qmgo.Config{Uri: uri, Database: db_name, Coll: tagcol})
	util.HandleError(err, "connection error")

	defer func() {
		if err = cli.Close(ctx); err != nil {
			panic(err)
		}
	}()
	
	result, err := cli.InsertMany(ctx, *Tags)
	util.HandleError(err, "insert ex listing error")
	fmt.Printf("%v", result)
}


func DB_fetch_tags() *[]TagData {
	ctx := context.Background()
	cli, err := qmgo.Open(ctx, &qmgo.Config{Uri: uri, Database: db_name, Coll: tagcol})
	util.HandleError(err, "connection error")

	defer func() {
		if err = cli.Close(ctx); err != nil {
			panic(err)
		}
	}()

	list := []TagData{}
	err = cli.Find(ctx, bson.M{}).All(&list)

	return &list
}