package router

import (
	"context"
	"github.com/gin-gonic/gin"
	"strconv"
	"workspace-goshow-mall/adaptor"
	"workspace-goshow-mall/adaptor/repo/vo"
	"workspace-goshow-mall/api/admin"
	"workspace-goshow-mall/api/user"
	"workspace-goshow-mall/config"
	"workspace-goshow-mall/utils/logger"
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
	adaptor   *adaptor.Adaptor
}

func NewRouter(adaptor *adaptor.Adaptor, config config.Config, checkFunc func() error) *Router {
	return &Router{
		FullPPROF: config.Server.EnablePprof,
		rootPath:  "/api",
		config:    config,
		user:      user.NewCtrl(adaptor),
		admin:     admin.NewCtrl(adaptor),
		checkFunc: checkFunc,
		adaptor:   adaptor,
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
	// 注册gin的错误处理中间件
	engine.Use(gin.Recovery())
	// 日志中间件
	engine.Use(AccessLogMiddleware(r.AccessRecordFilter))
	if r.config.Server.EnablePprof {
		//todo 增加链路追踪的逻辑
	}
	root := engine.Group(r.rootPath)
	engine.Any("/check", r.checkServer())
	r.route(root)
	err := engine.Run(":" + strconv.Itoa(r.config.Server.Port))
	if err != nil {
		logger.Debug("启动服务失败")
	}
}

func (r *Router) SpanFilter(c *gin.Context) bool {
	if c.Request.URL.Path == "/api/admin/login" ||
		c.Request.URL.Path == "/api/user/login" ||
		c.Request.URL.Path == "/api/user/register" {
		return false
	}
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
	r.adminRoute(root)
}

func (r *Router) adminRoute(root *gin.RouterGroup) {
	// 鉴权中间件
	adminRoute := root.Group("/admin", AdminAuthMiddleware(r.SpanFilter, func(c context.Context, token string) (*vo.UserVo, error) {
		return &vo.UserVo{}, nil
	}, r.adaptor))
	{
		adminRoute.POST("/create", r.admin.CreateUser)
	}
}
