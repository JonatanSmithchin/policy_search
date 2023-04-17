package user_service

import (
	"github.com/gin-gonic/gin"
	"police_search/Controllers"
	"police_search/Middlewares"
	"police_search/Models"
	"police_search/api"
)

//Register 新用户注册
func Register(c *gin.Context) {

	var user *Models.User

	if c.ShouldBind(&user) == nil {
		id, err := Controllers.CreateNewUser(user)
		if err != nil {
			c.JSON(500, &api.ReturnJson{
				Code:    500,
				Data:    err.Error(),
				Message: "cannot create new user",
			})
		} else {

			newJwt := Middlewares.NewJWT()
			token, err := newJwt.CreateToken(&Middlewares.CustomClaims{
				ID:       id,
				Password: user.UserPwd,
				Name:     user.UserName,
			})
			if err != nil {
				c.JSON(500, &api.ReturnJson{
					Code:    500,
					Data:    err,
					Message: "cannot create token",
				})
				//c.JSON(500, gin.H{"status": -1, "msg": "cannot create token"})
			}
			res := make(map[string]interface{}, 3)
			res["token"] = token
			res["UserId"] = id
			res["UserName"] = user.UserName
			c.JSON(200, &api.ReturnJson{
				Code:    200,
				Data:    res,
				Message: "registered successfully",
			})
		}
		//c.JSON(200, gin.H{"status:": 1, "msg": "successfully registered", "token": token})
	} else {
		c.JSON(500, &api.ReturnJson{
			Code:    500,
			Data:    nil,
			Message: "cannot bind user",
		})
		//c.JSON(500, "register failed")
	}
}

//LogIn 用户登录
func LogIn(c *gin.Context) {

	var user *Models.User

	if c.ShouldBind(&user) == nil {
		u := Controllers.GetUser(user.UserName)

		if u.UserPwd == user.UserPwd {
			//签发token
			newJwt := Middlewares.NewJWT()
			token, err := newJwt.CreateToken(&Middlewares.CustomClaims{
				ID:       u.UserID,
				Name:     u.UserName,
				Password: u.UserPwd,
			})

			if err != nil {
				c.JSON(500, &api.ReturnJson{
					Code:    500,
					Data:    err,
					Message: "cannot create token",
				})
				//c.JSON(500, gin.H{"status": -1, "msg": "cannot create token"})
			}
			res := make(map[string]interface{}, 4)
			res["token"] = token
			res["UserId"] = u.UserID
			res["UserName"] = u.UserName
			res["email"] = u.Email
			res["tendency"], _ = Models.FindTendency(u.UserID)
			c.JSON(200, &api.ReturnJson{
				Code:    200,
				Data:    res,
				Message: "log in successfully",
			})
			//c.JSON(200, gin.H{"status": 1, "msg": "you are logged in", "token": token})
		} else {
			c.JSON(401, &api.ReturnJson{
				Code:    401,
				Data:    nil,
				Message: "incorrect password",
			})
			//c.JSON(401, gin.H{"status": 1, "msg": "unauthorized"})
			//	if form.UserName == "usr" && form.UserPwd == "pwd" {
			//		c.JSON(200, gin.H{"status": "you are logged in"})
			//	} else {
			//		c.JSON(401, gin.H{"status": "unauthorized"})
			//	}
			//}
		}
	}
}

//ChangePassword 用户修改密码
func ChangePassword(c *gin.Context) {
	var body struct {
		UserName    string `json:"userName"`
		Password    string `json:"password"`
		NewPassword string `json:"newPassword"`
	}
	if c.ShouldBind(&body) == nil {
		user := Controllers.GetUser(body.UserName)
		if user.UserPwd == body.Password {
			err := Controllers.ChangePassword(user, body.NewPassword)
			if err != nil {
				c.JSON(500, &api.ReturnJson{
					Code:    500,
					Data:    nil,
					Message: "cannot change password",
				})
			} else {
				c.JSON(200, &api.ReturnJson{
					Code:    200,
					Data:    nil,
					Message: "success",
				})
			}
		} else {
			c.JSON(401, &api.ReturnJson{
				Code:    401,
				Data:    nil,
				Message: "incorrect password",
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

func SetEmail(c *gin.Context) {
	var body struct {
		UserName string `json:"userName"`
		Email    string `json:"email"`
	}
	if c.ShouldBind(&body) == nil {
		user := Controllers.GetUser(body.UserName)
		if user != nil {
			err := user.ChangeEmail(body.Email)
			if err != nil {
				c.JSON(500, &api.ReturnJson{
					Code:    500,
					Data:    nil,
					Message: "cannot change email",
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
