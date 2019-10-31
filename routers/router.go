package routers

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github/ToPeas/go-curd-templatepkg/setting"
	"github/ToPeas/go-curd-templaterouters/api/auth"
	"github/ToPeas/go-curd-templaterouters/api/web"
)

func InitRouter() *gin.Engine {
	if setting.Config.App.Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()

	r.Use(gin.Recovery())

	// CORS
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	//config.AllowCredentials = true
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	//config.AllowOrigins = []string{"http://localhost:8080"}
	r.Use(cors.New(config))

	// 首页
	r.GET("/", func(c *gin.Context) {
		c.String(200, "it works")
		return
	})

	// 健康检测
	r.HEAD("/", func(c *gin.Context) {
		c.Status(200)
		return
	})

	r.Use(gin.Logger())

	// Web后台  获取token
	webLoginNotRequired := r.Group("/admin/api")
	{
		webLoginNotRequired.POST("/auth", auth.Login)
	}

	webLoginRequired := r.Group("/admin/api")

	//webLoginRequired.Use(jwt.JWT())

	{
		// 获取links
		webLoginRequired.GET("/links", web.GetAllLinks)
		//// 获取link
		webLoginRequired.GET("/link/:id", web.GetLink)
		// 增加link
		webLoginRequired.POST("/link", web.AddLink)
		//// 删除link
		webLoginRequired.DELETE("/link/:id", web.DeleteLink)
		//// 修改link
		webLoginRequired.PATCH("/link/:id", web.PatchLink)
	}

	return r
}
