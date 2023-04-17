package admin_service

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"police_search/api"
	"police_search/es"
	"strconv"
)

func AllPolicy(c *gin.Context) {
	q := map[string]string{}
	from, _ := strconv.Atoi(c.Param("from"))
	size, _ := strconv.Atoi(c.Param("size"))
	res, err := es.Search("policy", -1, from, size, "PUB_TIME:desc", q)
	if err != nil {
		c.JSON(500, &api.ReturnJson{
			Code:    500,
			Data:    nil,
			Message: "failed",
		})
	} else {
		c.JSON(200, &api.ReturnJson{
			Code:    200,
			Data:    res,
			Message: "success",
		})
	}
}

func AddPolicy(c *gin.Context) {

	var b = make([]byte, 4096)
	l, err := c.Request.Body.Read(b)
	id := c.Param("id")
	if err != io.EOF {
		log.Printf("cannot get request body : %v \n", err)
		c.JSON(500, &api.ReturnJson{
			Code:    500,
			Data:    nil,
			Message: "failed to get request body",
		})
	}

	err = nil
	var buf = &bytes.Buffer{}
	buf.Write(b)
	buf.Truncate(l)
	//id, err := es.GetSize("policy")
	//if err != nil {
	//	log.Printf("cannot get policy size: %v \n", err)
	//	c.JSON(500, &api.ReturnJson{
	//		Code:    500,
	//		Data:    nil,
	//		Message: "failed to get policy size",
	//	})
	//}

	//idNum, err := strconv.Atoi(id)
	//if err != nil {
	//	log.Printf("cannot convert to string: %v \n", err)
	//	c.JSON(500, &api.ReturnJson{
	//		Code:    500,
	//		Data:    nil,
	//		Message: "failed to convert to string",
	//	})
	//}
	err = es.Add("policy", id, buf)
	if err != nil {
		log.Printf("cannot add policy")
		c.JSON(500, &api.ReturnJson{
			Code:    500,
			Data:    err,
			Message: "failed to add policy",
		})
	} else {
		c.JSON(200, &api.ReturnJson{
			Code:    200,
			Data:    id,
			Message: "success",
		})
	}
}

func DeletePolicy(c *gin.Context) {
	id := c.Param("id")
	err := es.Delete("policy", id)
	if err != nil {
		c.JSON(500, &api.ReturnJson{
			Code:    500,
			Data:    nil,
			Message: "cannot delete policy",
		})
	} else {
		c.JSON(200, &api.ReturnJson{
			Code:    200,
			Data:    nil,
			Message: "success",
		})
	}

}
