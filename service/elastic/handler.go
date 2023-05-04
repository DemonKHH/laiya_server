package elastic

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/elastic/go-elasticsearch/v7/esapi"
)

const indexName = "laiya"

func Insert(doc interface{}) {
	// log.Printf("doc: %v", doc)
	es := GetEsClient()
	// 添加文档
	docJSON, err := json.Marshal(doc)
	if err != nil {
		log.Printf("Error encoding document: %s", err)
	}
	addDocRequest := esapi.IndexRequest{
		Index:   indexName,
		Body:    bytes.NewReader(docJSON),
		Refresh: "true",
	}
	res, err := addDocRequest.Do(context.Background(), es)
	if err != nil {
		log.Printf("Error adding document: %s", err)
	}
	defer res.Body.Close()
}

func Update(doc interface{}) {
	es := GetEsClient()
	newDocJSON, _ := json.Marshal(doc)
	updateDocRequest := esapi.UpdateRequest{
		Index:   indexName,
		Body:    bytes.NewReader(newDocJSON),
		Refresh: "true",
	}
	res, err := updateDocRequest.Do(context.Background(), es)
	if err != nil {
		log.Fatalf("Error updating document: %s", err)
	}
	defer res.Body.Close()
}

func Search() {
	es := GetEsClient()
	// 查询文档
	searchRequest := esapi.SearchRequest{
		Index: []string{indexName},
	}
	res, err := searchRequest.Do(context.Background(), es)
	if err != nil {
		log.Fatalf("Error searching for documents: %s", err)
	}
	defer res.Body.Close()

	// 处理查询结果
	if res.IsError() {
		log.Fatalf("Error searching for documents: %s", res.Status())
	}
	var searchResult map[string]interface{}
	json.NewDecoder(res.Body).Decode(&searchResult)
	fmt.Println(searchResult)
}

func Delete() {
	es := GetEsClient()
	// 删除文档
	deleteDocRequest := esapi.DeleteRequest{
		Index:   indexName,
		Refresh: "true",
	}
	res, err := deleteDocRequest.Do(context.Background(), es)
	if err != nil {
		log.Fatalf("Error deleting document: %s", err)
	}
	defer res.Body.Close()
}
