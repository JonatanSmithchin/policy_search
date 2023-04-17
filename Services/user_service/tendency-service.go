package user_service

import (
	"github.com/gin-gonic/gin"
	"log"
	"police_search/Controllers"
	"police_search/api"
)

func SetTendency(c *gin.Context) {

	var body struct {
		UserID     int      `json:"UserID"`
		Tendencies []string `json:"Tendencies"`
	}

	if c.ShouldBind(&body) == nil {
		err := Controllers.SetTendency(body.UserID, body.Tendencies)
		if err != nil {
			log.Print(err)
			c.JSON(500, &api.ReturnJson{
				Code:    500,
				Data:    nil,
				Message: "cannot set tendencies",
			})
		} else {
			c.JSON(200, &api.ReturnJson{
				Code:    200,
				Data:    nil,
				Message: "success",
			})
		}
	} else {
		c.JSON(500, &api.ReturnJson{
			Code:    500,
			Data:    nil,
			Message: "cannot bind json",
		})
	}
}

func ResetTendency(c *gin.Context) {

	var body struct {
		UserID  int      `json:"UserID"`
		Removed []string `json:"Removed"`
		Added   []string `json:"Added"`
	}
	if c.ShouldBind(&body) == nil {
		err := Controllers.RemoveTendency(body.UserID, body.Removed) //首先删除要删除的倾向
		if err != nil {
			c.JSON(500, &api.ReturnJson{
				Code:    500,
				Data:    nil,
				Message: "cannot remove old tendency",
			})
		} else {
			err = Controllers.SetTendency(body.UserID, body.Added) //然后添加新倾向
			if err != nil {
				c.JSON(500, &api.ReturnJson{
					Code:    500,
					Data:    nil,
					Message: "cannot add tendency",
				})
			} else {
				c.JSON(200, &api.ReturnJson{
					Code:    200,
					Data:    nil,
					Message: "success",
				})
			}
		}
	} else {
		c.JSON(500, &api.ReturnJson{
			Code:    500,
			Data:    nil,
			Message: "cannot bind json",
		})
	}
}
