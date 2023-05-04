package elastic

import (
	"log"

	"github.com/elastic/go-elasticsearch/v7"
)

var esClient *elasticsearch.Client

func init() {
	var err error
	log.Printf("init elastic")
	// 创建Elasticsearch客户端
	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://47.108.160.82:9200",
		},
	}
	esClient, err = elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}
}

func GetEsClient() *elasticsearch.Client {
	return esClient
}
