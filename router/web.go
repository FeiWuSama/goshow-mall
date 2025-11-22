package router

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"workspace-goshow-mall/utils/logger"
)

type App struct {
	server  *gin.Engine
	address string
}

func NewApp(port int, router IRouter) *App {
	engine := gin.New()
	//注册路由
	router.Register(engine)
	return &App{
		server:  engine,
		address: ":" + strconv.Itoa(port),
	}
}

func (a *App) Run() {
	srv := &http.Server{
		Addr:    a.address,
		Handler: a.server,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen err: %v", err)
		}
	}()

	closeChanel := make(chan os.Signal, 1)
	signal.Notify(closeChanel, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	msg := <-closeChanel
	logger.Warn("server close", zap.String("msg", msg.String()))
}
