package Middlewares

//
//import (
//	"fmt"
//	"github.com/gin-gonic/gin"
//	"net/http"
//	"police_search/Models"
//	"police_search/api"
//	"strings"
//)
//
//var AuthCheckMiddleware = authCheck()
//
//func authCheck() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		if admin, _ := c.Get("admin"); admin != nil {
//
//			method := c.Request.Method
//			url := c.Request.URL.Path
//
//			roleId := admin
//			isSuper, _ := c.Get("isSuper")
//			//adminInfo := admin.(**service.RunningClaims)
//			//isSuper := (*adminInfo).IsSuper //是否是超管
//			//roleId := (*adminInfo).RoleId
//
//			if isSuper != 1 {
//				fmt.Println("method:", method)
//				fmt.Println("url:", url)
//				permissionFunc := strings.ToLower(fmt.Sprintf("%s_%s", method, url))
//				haveAuth := Models.CheckRolePermission(uint(roleId), permissionFunc)
//				fmt.Println("haveAuth  ", haveAuth)
//				if !haveAuth {
//					c.JSON(http.StatusOK, api.ReturnJson{Code: http.StatusForbidden, Data: "", Message: "无权访问"})
//					c.Abort()
//				}
//			}
//			c.Next()
//		}
//	}
//}
