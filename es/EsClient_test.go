package es

import (
	"github.com/stretchr/testify/require"
	"log"
	"testing"
)

//func TestNewEsClient(t *testing.T) {
//	cfg := &elasticsearch.Config{
//		Addresses: []string{
//			"http://121.37.119.47:9200",
//		},
//	}
//	es := NewEsClient(cfg)
//	resp, err := es.Info()
//	require.NoError(t, err)
//	log.Print(resp)
//}

//func TestAdd(t *testing.T) {
//	err := Add("shopping", "12", "{\"name\":\"apple\",\"slogan\":\"apple\"}")
//	require.NoError(t, err)
//}

func TestGetSize(t *testing.T) {
	id, err := GetSize("policy")
	log.Print(id)
	require.NoError(t, err)
}

func TestSearch(t *testing.T) {
	q := map[string]string{
		"POLICY_TITLE": "经济",
		"POLICY_GRADE": "省级",
		//"POLICY_TYPE":  "其他",
		//"PROVINCE":     "四川",
	}
	Search("policy", 0, 10, 10, "PUB_TIME:desc", q)
}

func TestRangeSearch(t *testing.T) {
	q := map[string]string{
		"POLICY_TITLE": "经济",
		"POLICY_GRADE": "省级",
		//"POLICY_TYPE":  "其他",
		//"PROVINCE":     "四川",
	}
	d := map[string]string{
		"gt":  "2018/01/07",
		"lte": "2018/01/09",
	}
	RangeSearch("policy", 0, 10, "PUB_TIME:desc", q, d)
}
