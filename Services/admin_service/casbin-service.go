package admin_service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"police_search/Controllers"
	"police_search/api"
)

func CreateCasbin(c *gin.Context) {

	var createCasbin *Controllers.CasbinCreateRequest

	if c.ShouldBind(&createCasbin) == nil {
		err := Controllers.CreateCasbin(createCasbin)
		if err != nil {
			c.JSON(500, &api.ReturnJson{
				Code:    500,
				Data:    nil,
				Message: "cannot create Auth",
			})
			//c.JSON(500, gin.H{"status": -1, "msg": "cannot create auth"})
		}
		c.JSON(200, &api.ReturnJson{
			Code:    200,
			Data:    createCasbin,
			Message: fmt.Sprintf("create Auth: %s %s", createCasbin.RoleId, createCasbin.CasbinInfos),
		})
		//c.JSON(200, gin.H{"status": 1, "msg": createCasbin.RoleId})
	}

}

func GetCasbinList(c *gin.Context) {
	roleID := c.Request.PostForm.Get("roleID")
	auths := Controllers.CasbinList(roleID)
	c.JSON(200, gin.H{"status": 1, "data": auths})
}
