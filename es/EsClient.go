package es

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"log"
	"police_search/Models"
	"strings"
	"time"
)

var EsCli *elasticsearch.Client

func init() {
	cfg := &elasticsearch.Config{
		Addresses: []string{
			"http://121.37.119.47:9200",
		},
	}
	tmp, err := elasticsearch.NewClient(*cfg)
	if err != nil {
		log.Printf("Error creating the client: %s", err)
	}

	_, err = tmp.Info()
	if err != nil {
		log.Printf("Error getting response: %s", err)
	}

	EsCli = tmp
}

func GetSize(index string) (string, error) {
	resp, err := EsCli.Cat.Indices(
		EsCli.Cat.Indices.WithIndex(index),
		EsCli.Cat.Indices.WithPretty(),
	)
	if err != nil {
		return "", err
	}

	res := strings.Split(resp.String(), " ")

	return res[8], nil

}

func Add(index string, id string, body *bytes.Buffer) error {

	resp, err := EsCli.Create(
		index, id, body,
		EsCli.Create.WithContext(context.Background()),
		EsCli.Create.WithDocumentType("_doc"),
	)
	if err != nil {
		return err
	}
	log.Println(resp.StatusCode)
	//id已经存在
	if resp.StatusCode == 409 {
		return fmt.Errorf("policyId %s already exist", id)
	}
	return nil

}

func buildRangeSearchQuery(from int, query map[string]string, r map[string]string) *bytes.Buffer {
	var buf bytes.Buffer
	var musts []map[string]interface{}
	for k, v := range query {
		must := map[string]interface{}{
			"terms": map[string][]string{
				k: {v},
			},
		}
		musts = append(musts, must)
	}
	//日期范围查询
	dateRange := map[string]interface{}{
		"PUB_TIME": map[string]string{
			"gte": r["gte"],
			"lte": r["lte"],
		},
	}

	q := map[string]interface{}{
		"from": from,
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"should": musts,
				"filter": map[string]interface{}{
					"range": dateRange,
				},
			},
		},
	}

	err := json.NewEncoder(&buf).Encode(q)
	if err != nil {
		log.Printf("cannot encoding query: %v", err)
		return nil
	}
	return &buf
}

func buildSearchQuery(from int, u int, query map[string]string) *bytes.Buffer {

	var buf bytes.Buffer
	var shoulds []map[string]interface{}
	var q map[string]interface{}
	for k, v := range query {
		should := map[string]interface{}{
			"terms": map[string][]string{
				k: {v},
			},
		}
		shoulds = append(shoulds, should)
	}
	if u != -1 {
		t, _ := Models.FindTendency(u)
		if t != nil {
			should := map[string]interface{}{
				"terms": map[string][]string{
					"POLICY_BODY": t,
				},
			}
			q = map[string]interface{}{
				"from": from,
				"query": map[string]interface{}{
					"bool": map[string]interface{}{
						"should": append(shoulds, should),
					},
				},
				"highlight": map[string]interface{}{
					"fields": map[string]interface{}{
						"POLICY_TITLE": map[string][]string{
							"pre_tags":  []string{"<em>"},
							"post_tags": []string{"</em>"},
						},
					},
				},
			}
		}
	} else {
		q = map[string]interface{}{
			"from": from,
			"query": map[string]interface{}{
				"bool": map[string]interface{}{
					"should": shoulds,
				},
			},
			"highlight": map[string]interface{}{
				"fields": map[string]interface{}{
					"POLICY_TITLE": map[string][]string{
						"pre_tags":  []string{"<em>"},
						"post_tags": []string{"</em>"},
					},
				},
			},
		}
	}

	//日期范围查询
	err := json.NewEncoder(&buf).Encode(q)
	if err != nil {
		log.Printf("cannot encoding query: %v", err)
		return nil
	}
	return &buf
}

func doResp(res *esapi.Response) []interface{} {
	var r map[string]interface{}
	err := json.NewDecoder(res.Body).Decode(&r)
	if err != nil {
		log.Printf("cannot parsing the response body: %v", err)
	}
	return r["hits"].(map[string]interface{})["hits"].([]interface{})
}

func Search(index string, u int, from int, size int, direction string, query map[string]string) ([]interface{}, error) {

	buf := buildSearchQuery(from, u, query)
	res, err := EsCli.Search(
		EsCli.Search.WithContext(context.Background()),
		EsCli.Search.WithBody(buf),
		EsCli.Search.WithIndex(index),
		EsCli.Search.WithPretty(),
		EsCli.Search.WithSize(size),
		EsCli.Search.WithSort(direction),
		//EsCli.Search.WithAnalyzer("ik_max_word"),
	)
	if err != nil {
		return nil, err
	}
	r := doResp(res)
	return r, nil
}

func RangeSearch(index string, from int, size int, direction string, query map[string]string, d map[string]string) ([]interface{}, error) {

	buf := buildRangeSearchQuery(from, query, d)
	res, err := EsCli.Search(
		EsCli.Search.WithContext(context.Background()),
		EsCli.Search.WithBody(buf),
		EsCli.Search.WithIndex(index),
		EsCli.Search.WithPretty(),
		EsCli.Search.WithSize(size),
		EsCli.Search.WithSort(direction),
	)
	if err != nil {
		return nil, err
	}
	r := doResp(res)
	return r, nil
}

func Delete(index string, id string) error {
	_, err := EsCli.Delete(
		index, id,
		EsCli.Delete.WithContext(context.Background()),
	)
	if err != nil {
		return err
	}
	return nil
}

//func NewEsClient(cfg *elasticsearch.Config) *elasticsearch.Client {
//	es, err := elasticsearch.NewClient(*cfg)
//	if err != nil {
//		log.Printf("Error creating the client: %s", err)
//		return nil
//	}
//
//	_, err = es.Info()
//	if err != nil {
//		log.Printf("Error getting response: %s", err)
//		return nil
//	}
//	return es
//}

type ElasticDocs struct {
	id   int
	freq string
	time time.Time
}

func jsonStruct(doc ElasticDocs) string {
	docStruct := &ElasticDocs{
		id:   doc.id,
		freq: doc.freq,
		time: doc.time,
	}
	b, err := json.Marshal(docStruct)
	if err != nil {
		log.Printf("json.Marshal ERROR: %s", err)
	}
	return string(b)
}

//func search(id string) {
//	cfg := &elasticsearch.Config{
//		Addresses: []string{
//			"http://121.37.119.47:9200",
//		},
//	}
//	client := NewEsClient(cfg)
//	recom, _ := Redis.GetAllRecommend(id)
//	recomId := make([]string, len(recom))
//	i := 0
//	for id, _ := range recom {
//		recomId[i] = "10" + id
//		i += 1
//	}
//	var buf bytes.Buffer
//	query := map[string]interface{}{
//		"query": map[string]interface{}{
//			"ids": map[string]interface{}{
//				"values": recomId,
//			},
//		},
//	}
//	err := json.NewEncoder(&buf).Encode(query)
//	if err != nil {
//		log.Printf("cannot encoding query: %v", err)
//	}
//	res, _ := client.Search(
//		client.Search.WithContext(context.Background()),
//		client.Search.WithBody(&buf),
//		client.Search.WithIndex("policy"),
//		client.Search.WithPretty(),
//	)
//	log.Print(res)
//}
