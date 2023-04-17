package Router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"police_search/Middlewares"
	"police_search/Services"
	"police_search/Services/admin_service"
	"police_search/Services/user_service"
	"time"
)

func logFormat(param gin.LogFormatterParams) string {
	// your custom format
	return fmt.Sprintf("%s - [%s] %d \"%s %s %s %d %s \"%s\" %s\"\n",
		param.ClientIP,
		param.TimeStamp.Format(time.RFC3339Nano),
		param.BodySize,
		param.Method,
		param.Path,
		param.Request.Proto,
		param.StatusCode,
		param.Latency,
		param.Request.UserAgent(),
		param.ErrorMessage,
	)
}

func InitRouter() {
	r := gin.Default()

	r.Use(gin.LoggerWithFormatter(logFormat))
	r.Use(Middlewares.Cors())

	r.POST("/register", user_service.Register)
	r.POST("/login", user_service.LogIn)

	adminRouter := r.Group("/admin")
	userRouter := r.Group("/user")
	searchRouter := r.Group("/search")

	{
		userRouter.Use(Middlewares.JWTAuth())
		userRouter.GET("/query", func(c *gin.Context) {
			c.JSON(200, gin.H{"status": "ok"})
		})
		profileRouter := userRouter.Group("/profile")
		{
			profileRouter.POST("/changePassword", user_service.ChangePassword)
			profileRouter.POST("/setEmail", user_service.SetEmail)
		}
		userRouter.POST("/setTendency", user_service.SetTendency)
		userRouter.POST("/resetTendency", user_service.ResetTendency)
		userRouter.PUT("/updateRecord/:UserId/:PolicyId", user_service.UpdateRec)
		//userRouter.POST("/addRecord", user_service.AddRecord)
		//userRouter.PUT("/updateDuration", user_service.UpdateDuration)
	}

	{
		searchRouter.GET("/:UserId", Services.AutoRecom)
		searchRouter.POST("/policy/:id", Services.SearchService)
		searchRouter.POST("/policy/range", Services.RangeSearch)
		searchRouter.POST("test", Services.TestBind)
	}

	{
		adminRouter.POST("/login", admin_service.AdminLogin)
		adminRouter.Use(Middlewares.JWTAuth())
		adminRouter.Use(Middlewares.CasbinHandler())
		adminRouter.Use(Middlewares.AdminLogs)
		adminRouter.POST("/createUser", admin_service.CreateUser)
		adminRouter.GET("/users", admin_service.GetAllUsers)
		adminRouter.GET("/admins", admin_service.GetAllAdmins)
		adminRouter.POST("/createAdmin", admin_service.CreateAdmin)
		adminRouter.DELETE("/deleteAdmin/:name", admin_service.DeleteAdmin)
		adminRouter.POST("/createAuth", admin_service.CreateCasbin)
		adminRouter.GET("/getAuths", admin_service.GetCasbinList)
		adminRouter.POST("/addPolicy/:id", admin_service.AddPolicy)
		adminRouter.DELETE("/deletePolicy/:id", admin_service.DeletePolicy)
		adminRouter.GET("/policy/:from/:size", admin_service.AllPolicy)
		logsRouter := adminRouter.Group("/logs")
		{
			logsRouter.GET("/getLogs", admin_service.GetLogs)
			logsRouter.GET("/:managerName", admin_service.FindLogsByName)
		}
		dataRouter := adminRouter.Group("/data")
		{
			dataRouter.POST("/addData", admin_service.AddPolicy)
		}
	}

	r.Run(":8080")
}
