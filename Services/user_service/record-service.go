package user_service

import (
	"github.com/gin-gonic/gin"
	"police_search/Redis"
	"police_search/api"
)

func UpdateRec(c *gin.Context) {
	UserID := c.Param("UserId")
	policyID := c.Param("PolicyId")
	//var buf *bytes.Buffer
	//_, _ = buf.WriteString(policyID)
	//search, err := es.EsCli.Search(
	//	es.EsCli.Search.WithIndex("policy"),
	//	es.EsCli.Search.WithBody(buf))
	//if err != nil {
	//	return
	//}
	err := Redis.UpDateRec(UserID, policyID)
	if err != nil {
		c.JSON(500, &api.ReturnJson{
			Code:    500,
			Data:    nil,
			Message: "cannot record",
		})
	} else {
		c.JSON(200, &api.ReturnJson{
			Code:    200,
			Data:    nil,
			Message: "success",
		})
	}
}

//func AddRecord(c *gin.Context) {
//	var record *Models.Footprint
//	if c.ShouldBind(&record) == nil {
//		record.IP = api.GetClientIP(c)
//		recId, err := Controllers.AddRecord(record)
//		if err != nil {
//			c.JSON(500, &api.ReturnJson{
//				Code:    500,
//				Data:    nil,
//				Message: "cannot add record",
//			})
//		} else {
//			c.JSON(200, &api.ReturnJson{
//				Code:    200,
//				Data:    recId,
//				Message: "success",
//			})
//		}
//	} else {
//		c.JSON(500, &api.ReturnJson{
//			Code:    500,
//			Data:    nil,
//			Message: "cannot bind json",
//		})
//	}
//}
//
//func UpdateDuration(c *gin.Context) {
//	var body struct {
//		RecordId int `json:"recordId"`
//		Duration int `json:"duration"`
//	}
//	if c.ShouldBind(&body) == nil {
//		err := Controllers.UpdateRecordDuration(body.RecordId, body.Duration)
//		if err != nil {
//			c.JSON(500, &api.ReturnJson{
//				Code:    500,
//				Data:    nil,
//				Message: "cannot update duration",
//			})
//		} else {
//			c.JSON(200, &api.ReturnJson{
//				Code:    200,
//				Data:    nil,
//				Message: "success",
//			})
//		}
//
//	} else {
//		c.JSON(500, &api.ReturnJson{
//			Code:    500,
//			Data:    nil,
//			Message: "cannot bind json",
//		})
//	}
//}
