package Middlewares

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"police_search/Models"
)

func CasbinHandler() gin.HandlerFunc {
	return func(c *gin.Context) {

		// 获取请求的URI
		obj := c.Request.URL.RequestURI()
		// 获取请求方法
		act := c.Request.Method
		// 获取用户的角色

		jsonClaims := c.GetString("claims")
		var claims CustomClaims
		err := json.Unmarshal([]byte(jsonClaims), &claims)
		if err != nil {
			c.JSON(500, gin.H{"status": -1, "msg": "cannot get claims"})
			c.Abort()
			return
		}
		sub := claims.Name
		e := Models.Casbin()
		fmt.Println(obj, act, sub)
		// 判断策略中是否存在
		success, err := e.Enforce(sub, obj, act)
		if success {
			log.Println("恭喜您,权限验证通过")
			c.Next()
		} else {
			log.Printf("e.Enforce err: %v", err)
			c.JSON(403, gin.H{"code": -1, "msg": "unauthorized"})
			c.Abort()
			return
		}
	}
}
