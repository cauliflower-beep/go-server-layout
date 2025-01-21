/**
 * @Author: LiuShuXin
 * @Description:
 * @File:  route
 * Software: Goland
 * @Date: 2025/1/21 9:58
 */

package router

import (
	"app-server/middlewares"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"net/http"

	"github.com/gin-contrib/pprof"
)

func SetupRouter(mode string) *gin.Engine {
	gin.SetMode(mode)

	r := gin.New()
	// 使用自定义logger、recovery中间件取代gin默认的
	r.Use(middlewares.Logger())
	//r.Use(logger.GinLogger(), logger.GinRecovery(true))

	// 本地测试页面 线上需要关闭
	r.LoadHTMLGlob("templates/*") // 加载模板
	r.Static("/assets", "static") // 设置静态文件路径

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	pprof.Register(r) // 注册pprof相关路由

	// 自定义路由

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})

	return r

}
