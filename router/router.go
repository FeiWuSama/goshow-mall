package router

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"workspace-goshow-mall/adaptor"
	"workspace-goshow-mall/api/admin"
	"workspace-goshow-mall/api/user"
	"workspace-goshow-mall/config"
)

type IRouter interface {
	Register(engine *gin.Engine)
	SpanFilter(c *gin.Context) bool
	AccessRecordFilter(c *gin.Context) bool
}

type Router struct {
	FullPPROF bool
	rootPath  string
	config    config.Config
	checkFunc func() error
	user      *user.Ctrl
	admin     *admin.Ctrl
}

func NewRouter(adaptor *adaptor.Adaptor, config config.Config, checkFunc func() error) *Router {
	return &Router{
		FullPPROF: config.Server.EnablePprof,
		rootPath:  "/api",
		config:    config,
		user:      user.NewCtrl(adaptor),
		admin:     admin.NewCtrl(adaptor),
		checkFunc: checkFunc,
	}
}

func (r *Router) checkServer() func(c *gin.Context) {
	return func(c *gin.Context) {
		err := r.checkFunc()
		if err != nil {
			c.JSON(500, gin.H{
				"message": err.Error(),
			})
			c.Abort()
			return
		}
		c.JSON(200, gin.H{
			"message": "check success",
		})
	}
}

func (r *Router) Register(engine *gin.Engine) {
	// 默认gin的日志级别是release级别
	gin.SetMode(gin.DebugMode)
	// 添加日志中间件
	engine.Use(gin.Logger())
	//注册gin的错误处理中间件
	engine.Use(gin.Recovery())
	//todo 日志中间件

	if r.config.Server.EnablePprof {
		//todo 增加链路追踪的逻辑
	}
	root := engine.Group(r.rootPath)
	engine.Any("/check", r.checkServer())
	r.route(root)
	engine.Run(":" + strconv.Itoa(r.config.Server.Port))
}

func (r *Router) SpanFilter(c *gin.Context) bool {
	return true
}

func (r *Router) AccessRecordFilter(c *gin.Context) bool {
	return true
}

func (r *Router) route(root *gin.RouterGroup) {
	root.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ok",
		})
	})
	root.GET("hello", r.admin.HelloWorld)
}
