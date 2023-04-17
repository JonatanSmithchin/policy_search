package Middlewares

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"police_search/Models"
	"police_search/api"
	"strings"
)

type bodyLogWriter struct {
	gin.ResponseWriter
	bodyBuf *bytes.Buffer
}

func (w *bodyLogWriter) Write(b []byte) (int, error) {
	w.bodyBuf.Write(b)
	return w.ResponseWriter.Write(b)
}

var CommonLogInterceptor = commonLogInterceptor()

func commonLogInterceptor() gin.HandlerFunc {
	return func(c *gin.Context) {
		AdminLogs(c)
	}
}

func AdminLogs(c *gin.Context) {

	jsonClaims := c.GetString("claims")
	var claims CustomClaims
	err := json.Unmarshal([]byte(jsonClaims), &claims)
	if err != nil {
		c.JSON(500, gin.H{"status": -1, "msg": "cannot get claims"})
		c.Abort()
	}

	if claims.Name != "" {
		method := c.Request.Method
		url := c.Request.URL.Path

		strBody := ""
		var blw bodyLogWriter
		blw = bodyLogWriter{
			ResponseWriter: c.Writer,
			bodyBuf:        bytes.NewBufferString(strBody),
		}
		c.Writer = &blw
		c.Next()

		if method != "GET" {
			strBody = strings.Trim(blw.bodyBuf.String(), "\n")
			go func(strBody string) {
				var returnJson api.ReturnJson
				json.Unmarshal([]byte(strBody), &returnJson)
				message := fmt.Sprintf("%v", returnJson.Message)

				//adminInfo := admin.(**service.RunningClaims)
				//adminId := (*adminInfo).ID
				//adminName := (*adminInfo).Account

				var log = Models.AdminLog{
					AdminId:   uint(claims.ID),
					AdminName: claims.Name,
					Method:    method,
					Url:       url,
					Ip:        api.GetClientIP(c),
					Code:      returnJson.Code,
					Message:   message,
				}

				log.CreateLog(log)
			}(strBody)
		}

	}
}
