package admin_service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"police_search/Controllers"
	"police_search/Middlewares"
	"police_search/Models"
	"police_search/api"
)

//AdminLogin 管理员登录
func AdminLogin(c *gin.Context) {
	var admin *Models.Admin

	if c.ShouldBind(&admin) == nil {
		a := Models.GetAdminByName(admin.Name)
		if admin.Password == a.Password {

			newJwt := Middlewares.NewJWT()
			token, err := newJwt.CreateToken(&Middlewares.CustomClaims{
				ID:   a.ID,
				Name: a.Name,
			})
			if err != nil {
				c.JSON(500, &api.ReturnJson{
					Code:    500,
					Data:    nil,
					Message: "internal admin login err",
				})
			}

			c.JSON(200, &api.ReturnJson{
				Code:    200,
				Data:    token,
				Message: "admin login successfully",
			})
			//c.JSON(200, gin.H{"status": 1, "msg": "you are logged in", "token": token, "data": admin})
		} else {
			c.JSON(401, &api.ReturnJson{
				Code:    401,
				Data:    nil,
				Message: "unauthorized",
			})
			//c.JSON(401, gin.H{"status": 1, "msg": "unauthorized"})
		}
	}
}

func GetAllAdmins(c *gin.Context) {
	admins := Controllers.GetAllAdmins()
	if admins == nil {
		c.JSON(500, &api.ReturnJson{
			Code:    500,
			Data:    nil,
			Message: "cannot find users",
		})
	}
	c.JSON(200, &api.ReturnJson{
		Code:    200,
		Data:    *admins,
		Message: "find all users",
	})
}

func CreateAdmin(c *gin.Context) {
	var a *Models.Admin

	if c.ShouldBind(&a) == nil {

		err := Controllers.CreateAdmin(a)

		if err != nil {
			c.JSON(500, &api.ReturnJson{
				Code:    500,
				Data:    err,
				Message: "cannot create admin",
			})
		}

		c.JSON(200, &api.ReturnJson{
			Code:    200,
			Data:    nil,
			Message: "successfully create admin",
		})
	} else {
		c.JSON(500, &api.ReturnJson{
			Code:    500,
			Data:    nil,
			Message: "cannot bind admin",
		})
	}
}

func DeleteAdmin(c *gin.Context) {
	name := c.Param("name")
	err := Controllers.DeleteAdmin(name)
	if err != nil {
		c.JSON(500, &api.ReturnJson{
			Code:    500,
			Data:    nil,
			Message: "cannot delete users",
		})
	} else {
		c.JSON(200, &api.ReturnJson{
			Code:    200,
			Data:    nil,
			Message: "remove admin successfully",
		})
	}
}

func GetAllUsers(c *gin.Context) {
	users := Controllers.GetAllUsers()
	if users == nil {
		c.JSON(500, &api.ReturnJson{
			Code:    500,
			Data:    nil,
			Message: "cannot find users",
		})
	}
	c.JSON(200, &api.ReturnJson{
		Code:    200,
		Data:    *users,
		Message: "find all users",
	})
	//c.JSON(200, gin.H{"users:": *users})
}

func CreateUser(c *gin.Context) {

	var user *Models.User

	if c.ShouldBind(&user) == nil {

		_, err := Controllers.CreateNewUser(user)

		if err != nil {
			c.JSON(500, &api.ReturnJson{
				Code:    500,
				Data:    err,
				Message: "cannot create user",
			})
			//c.JSON(500, gin.H{"status": -1, "msg": "cannot create user"})
		}

		//if err != nil {
		//	c.JSON(500, gin.H{"status": -1, "msg": "cannot create token"})
		//}莫名其妙的代码

		c.JSON(200, &api.ReturnJson{
			Code:    200,
			Data:    nil,
			Message: "successfully create",
		})
		//c.JSON(200, gin.H{"status:": 1, "msg": "successfully create"})
	} else {
		c.JSON(500, &api.ReturnJson{
			Code:    500,
			Data:    nil,
			Message: "cannot bind user",
		})
	}

}

//GetLogs 获取所有操作日志
func GetLogs(c *gin.Context) {
	logs := Controllers.GetAllLogs()
	if logs == nil {
		c.JSON(500, &api.ReturnJson{
			Code:    500,
			Data:    nil,
			Message: "cannot get logs",
		})
	}
	c.JSON(200, &api.ReturnJson{
		Code:    200,
		Data:    logs,
		Message: "get logs successfully",
	})
}

func FindLogsByName(c *gin.Context) {
	manager := c.Param("managerName")
	logs := Controllers.FindLogsByName(manager)
	c.JSON(200, &api.ReturnJson{
		Code:    200,
		Data:    logs,
		Message: fmt.Sprintf("find %s's logs", manager),
	})
}
