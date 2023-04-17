package Services

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"io"
	"log"
	"police_search/Redis"
	"police_search/api"
	"police_search/es"
	"strconv"
	"strings"
)

type result map[string]any

type Params map[string]string

type PathParams struct {
	Keyword    string `json:"keyword" form:"keyword"`
	PageNo     int    `json:"pageNo" form:"pageNo"`
	PageSize   int    `json:"pageSize" form:"pageSize"`
	Grade      string `json:"grade" form:"grade"`
	PolicyType string `json:"policyType" form:"policyType"`
	Province   string `json:"province" form:"province"`
}

func TestBind(c *gin.Context) {
	var params = &PathParams{}
	err := c.ShouldBind(&params)
	log.Print(params)
	c.JSON(200, err)
}

func (p PathParams) ConvertToMap() map[string]string {
	params := Params{}
	params["keyword"] = p.Keyword
	params["pageNo"] = fmt.Sprintf("%v", p.PageNo)
	params["pageSize"] = fmt.Sprintf("%v", p.PageSize)
	params["grade"] = p.Grade
	params["type"] = p.PolicyType
	params["province"] = p.Province
	return params
}

func Search(from int, size int, u int, direction string, params map[string]string) ([]interface{}, error) {

	res, err := es.Search("policy", u, from, size, direction, params)
	if err != nil {
		return nil, err
	}
	return res, nil
	//var results []result
	//host, err := consul.FindServer("search")
	//if err != nil {
	//	log.Println(err)
	//	return nil, err
	//} else {
	//	client := resty.New()
	//	resp, err := client.R().SetPathParams(params).SetResult(&results).Get("http://" + host + "/search/{keyword}/{pageNo}/{pageSize}/{grade}/{type}/{province}")
	//	if err != nil {
	//		return nil, err
	//	}
	//	log.Print(resp.Time())
	//	return results, nil
	//}
}

func rSearch(from int, size int, direction string, params map[string]string, r map[string]string) ([]interface{}, error) {

	res, err := es.RangeSearch("policy", from, size, direction, params, r)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func GetProvince(ip string) string {
	client := resty.New()
	var p struct {
		Province []map[string]interface{} `json:"data"`
	}
	_, err := client.R().SetResult(&p).SetQueryString(fmt.Sprintf("query=%s&co=&resource_id=6006&oe=utf8", ip)).Get("http://opendata.baidu.com/api.php")
	if err != nil {
		log.Printf("cannot get location %v", err)
		return ""
	}
	//TODO

	//r, _ := regexp.Compile("/(北京|天津|河北|山西|内蒙古|辽宁|吉林|黑龙江|上海|江苏|浙江|安徽|福建|江西|山东|河南|湖北|湖南|广东|广西|海南|重庆|四川|贵州|云南|西藏|陕西|甘肃|青海|宁夏|新疆|台湾)/.?")
	log.Print(p.Province[0]["location"])
	province, _, _ := strings.Cut(fmt.Sprintf("%v", p.Province[0]["location"]), "省")
	//province := r.FindString(fmt.Sprintf("%v", p.Province[0]["location"]))
	return province
}

type params struct {
	From      int               `json:"from"`
	Size      int               `json:"size"`
	Direction string            `json:"direction"`
	Params    map[string]string `json:"params"`
	Range     map[string]string `json:"range"`
}

func SearchService(c *gin.Context) {

	//var params *PathParams
	//if c.ShouldBind(&params) == nil {
	//m := params.ConvertToMap()
	//var m map[string]string
	var p *params

	if c.ShouldBind(&p) == nil {
		u, _ := strconv.Atoi(c.Param("id"))
		var results []interface{}
		results, err := Search(p.From, p.Size, u, p.Direction, p.Params)
		if err != nil {
			c.JSON(500, &api.ReturnJson{
				Code:    500,
				Data:    results,
				Message: "cannot do search",
			})
		} else {
			//if m["province"] == "" {
			//	m["province"] = GetProvince(api.GetClientIP(c))
			//	res, err := Search(m)
			//	if err != nil {
			//		c.JSON(500, &api.ReturnJson{
			//			Code:    500,
			//			Data:    results,
			//			Message: "cannot do search",
			//		})
			//	}
			//	results = append(results, res...)
			//ip := api.GetClientIP(c)
			//}
			c.JSON(200, &api.ReturnJson{
				Code:    200,
				Data:    results,
				Message: "success",
			})
		}
	}
	//} else {
	//	c.JSON(500, &api.ReturnJson{
	//		Code:    500,
	//		Data:    nil,
	//		Message: "cannot bind json",
	//	})
	//}

	//	err := search()
	//	//resp, err := client.R().SetResult(&results).Get("http://" + host + "/hello")
	//	if err != nil {
	//		c.JSON(500, &api.ReturnJson{
	//			Code:    500,
	//			Data:    nil,
	//			Message: "illegal request",
	//		})
	//	} else {
	//		log.Println(resp.Body())
	//		c.JSON(200, &api.ReturnJson{
	//			Code:    200,
	//			Data:    results,
	//			Message: "success",
	//		})
	//	}
	//}
}

func RangeSearch(c *gin.Context) {
	var p *params

	if c.ShouldBind(&p) == nil {

		var results []interface{}
		results, err := rSearch(p.From, p.Size, p.Direction, p.Params, p.Range)
		if err != nil {
			c.JSON(500, &api.ReturnJson{
				Code:    500,
				Data:    results,
				Message: "cannot do search",
			})
		} else {
			//if m["province"] == "" {
			//	m["province"] = GetProvince(api.GetClientIP(c))
			//	res, err := Search(m)
			//	if err != nil {
			//		c.JSON(500, &api.ReturnJson{
			//			Code:    500,
			//			Data:    results,
			//			Message: "cannot do search",
			//		})
			//	}
			//	results = append(results, res...)
			//ip := api.GetClientIP(c)
			//}
			c.JSON(200, &api.ReturnJson{
				Code:    200,
				Data:    results,
				Message: "success",
			})
		}
	}
}

func buildRecomQuery(recom *map[string]string) *bytes.Buffer {

	recomId := make([]string, len(*recom))
	i := 0
	for id, _ := range *recom {
		recomId[i] = "10" + id
		i += 1
	}
	var buf bytes.Buffer
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"ids": map[string]interface{}{
				"values": recomId,
			},
		},
	}
	err := json.NewEncoder(&buf).Encode(query)
	if err != nil {
		log.Printf("cannot encoding query: %v", err)
		return nil
	}
	return &buf
}

func AutoRecom(c *gin.Context) {

	id := c.Param("UserId")
	recom, err := Redis.GetAllRecommend(id)
	if err != nil {
		log.Printf("cannot get recommend for user: %v err: %v", id, err)
		c.JSON(500, &api.ReturnJson{
			Code:    500,
			Data:    nil,
			Message: "failed",
		})
	}

	buf := buildRecomQuery(&recom)

	if buf == nil {
		c.JSON(500, &api.ReturnJson{
			Code:    500,
			Data:    nil,
			Message: "cannot recommend",
		})
	}

	res, err := es.EsCli.Search(
		es.EsCli.Search.WithContext(context.Background()),
		es.EsCli.Search.WithBody(buf),
		es.EsCli.Search.WithIndex("policy"),
		es.EsCli.Search.WithPretty(),
	)
	if err != nil {
		log.Printf("cannot do search for user: %v err: %v", id, err)
		c.JSON(500, &api.ReturnJson{
			Code:    500,
			Data:    nil,
			Message: "cannot recommend",
		})
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Printf("cannot close response body %v", err)
		}
	}(res.Body)

	if res.IsError() {
		var e map[string]interface{}
		err := json.NewDecoder(res.Body).Decode(&e)
		if err != nil {
			log.Printf("cannot parsing the response body: %v", err)
		} else {
			log.Printf("[%s] %s: %s",
				res.Status(),
				e["error"].(map[string]interface{})["type"],
				e["error"].(map[string]interface{})["reason"])
		}
	} else {
		var r map[string]interface{}
		err := json.NewDecoder(res.Body).Decode(&r)
		if err != nil {
			log.Printf("cannot parsing the response body: %v", err)
		}
		log.Print(res)
		idNum, _ := strconv.Atoi(id)
		rec, err := Search(1, 10-len(recom), idNum, "PUB_TIME:desc", nil)
		var pRecom []interface{}
		if err == nil {
			for _, v := range rec {
				pRecom = append(r["hits"].(map[string]interface{})["hits"].([]interface{}), v)
			}
		}
		c.JSON(200, &api.ReturnJson{
			Code:    200,
			Data:    pRecom,
			Message: "success",
		})
	}

}

//func SearchService(c *gin.Context) {
//
//	var body *RequestBody
//
//	err := c.ShouldBindJSON(&body)
//
//	if err == nil {
//
//		request := &pb.SearchPolicyRequest{
//			Data:     &pb.UserData{Fields: body.Data},
//			Keywords: body.Keywords,
//		}
//
//		etcdClient := etcd_client.NewGrpcClient("121.37.119.47", 20000, "/search_policy")
//		ip := etcdClient.GetGrpcservIp()
//		log.Printf("get service from %s\n", ip)
//
//		conn, err := grpc.Dial(ip, grpc.WithTransportCredentials(insecure.NewCredentials()))
//		if err != nil {
//			log.Fatalf("cannot get service %v", err)
//		}
//
//		searchClient := pb.NewSearchPolicyServiceClient(conn)
//		stream, err := searchClient.SearchPolicy(context.Background(), request)
//		if err != nil {
//			c.JSON(500, &api.ReturnJson{
//				Code:    500,
//				Data:    err,
//				Message: "search err occurred",
//			})
//		}
//		for {
//
//			res, err := stream.Recv()
//
//			if err == io.EOF {
//				break
//			}
//
//			if err != nil {
//				c.JSON(500, &api.ReturnJson{
//					Code:    500,
//					Data:    err,
//					Message: "cannot read from result stream",
//				})
//			}
//
//			polices := res.GetResults()
//			c.JSON(200, &api.ReturnJson{
//				Code:    200,
//				Data:    polices,
//				Message: "search successfully",
//			})
//		}
//	} else {
//		c.JSON(500, &api.ReturnJson{
//			Code:    500,
//			Data:    body,
//			Message: "cannot bind json",
//		})
//	}
//}
