package main

import (
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v9"
	"strings"
	"time"
)

func main() {
	cf := elasticsearch.Config{
		Addresses: []string{"http://localhost:9201"},
		Username:  "elastic",
		Password:  "Xiaomi123@#",
	}
	es, err := elasticsearch.NewClient(cf)
	if err != nil {
		panic(err)
	}
	//createIndex(es)
	GetIndexes(es)
	//Search(es)
}

type test struct {
	Context string  `json:"context"`
	Price   float64 `json:"price"`
	Name    string  `json:"name"`
}

// 创建索引
func createIndex(es *elasticsearch.Client) {
	mapping := `{
	"settings": {
          "number_of_shards":   1,
          "number_of_replicas": 1
	},
	"mappings": {
		"properties": {
			"context":    { "type": "text"},
			"price":    { "type": "float" },
			"name":  { "type": "keyword" }
		}
	}
}`
	res, err := es.Indices.Create("my-index", es.Indices.Create.WithBody(strings.NewReader(mapping)))
	if err != nil {
		panic(fmt.Sprintf("create index error: %s\n", err))

	}
	defer res.Body.Close()

	fmt.Printf("res %#v\n", res)
}

// 获取索引

func GetIndexes(es *elasticsearch.Client) {
	res, err := es.Indices.Get([]string{"_all"}) // _all:所有的索引
	if err != nil {
		panic(fmt.Sprintf("get indexes error: %s", err))
	}
	defer res.Body.Close()
	var indices map[string]any
	if err = json.NewDecoder(res.Body).Decode(&indices); err != nil {
		panic(fmt.Sprintf("parse decode error: %s", err))
	}
	fmt.Printf("%#v\n", indices)
}

// 创建文档
func CreateDoc(es *elasticsearch.Client) {
	t := test{
		Context: "this is a test",
		Price:   10,
		Name:    "lisi",
	}
	bt, _ := json.Marshal(t)

	res, err := es.Create("my-index", fmt.Sprintf("%d", time.Now().Nanosecond()), strings.NewReader(string(bt)))
	if err != nil {
		panic(fmt.Sprintf("create index error: %s\n", err))

	}
	defer res.Body.Close()

	fmt.Printf("res %#v\n", res)
	fmt.Println(res.String())
}

//获取文档

func GetOnDoc(es *elasticsearch.Client, id string) {
	var rs map[string]any
	res, err := es.Get("my-index", id)
	if err != nil {
		panic(fmt.Sprintf("get one doc error: %s", err))

	}
	defer res.Body.Close()
	if err = json.NewDecoder(res.Body).Decode(&rs); err != nil {
		panic(fmt.Sprintf("parse error: %s", err))
	}
	fmt.Printf("res:%#v\n", rs)
}

//查询

func Search(es *elasticsearch.Client) {
	query := `
{
	"query":{
		"match_all":{}
	}
}
	`
	var rs map[string]any
	res, err := es.Search(
		es.Search.WithIndex("my-index"),
		es.Search.WithBody(strings.NewReader(query)),
	)
	if err != nil {
		panic(fmt.Sprintf("search error: %s", err))
	}
	defer res.Body.Close()
	if err = json.NewDecoder(res.Body).Decode(&rs); err != nil {
		panic(fmt.Sprintf("parse error: %s", err))
	}
	fmt.Printf("res:%#v\n", rs)
	fmt.Println(">>>>>>>>>>>>>>>>>>")
	hits := rs["hits"].(map[string]interface{})["hits"].([]interface{})
	for _, v := range hits {
		vv := v.(map[string]interface{})
		vvv := vv["_source"].(map[string]interface{})
		fmt.Printf("_id=%s,context=%s,name=%s,price=%f \n", vv["_id"], vvv["context"], vvv["name"], vvv["price"])
	}
}
