package main

import (
	"context"
	"github.com/dsxg666/notebook/global"
	"github.com/dsxg666/notebook/internal/db"
	"github.com/dsxg666/notebook/internal/routers"
	"github.com/dsxg666/notebook/pkg/log"
	"github.com/dsxg666/notebook/pkg/setting"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func init() {
	log.InitLogger()
	err := setupSetting()
	if err != nil {
		global.SugarLogger.Error("setupSetting err:", err)
	}
	global.MySqlConn = db.InitDB(global.DatabaseSetting)
}

func main() {
	gin.SetMode(global.ServerSetting.RunMode)
	router := newRouter()
	s := &http.Server{
		Addr:           ":" + global.ServerSetting.HttpPort,
		Handler:        router,
		ReadTimeout:    global.ServerSetting.ReadTimeout,
		WriteTimeout:   global.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	go func() {
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			global.SugarLogger.Error("ListenAndServer err:", err)
		}
	}()
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	global.SugarLogger.Fatal("Shuting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		global.SugarLogger.Fatal("Server forced to shutdown:", err)
	}
	global.SugarLogger.Info("Server Exiting")
}

func setupSetting() error {
	s, err := setting.NewSetting()
	if err != nil {
		return err
	}
	err = s.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}
	err = s.ReadSection("Email", &global.EmailSetting)
	if err != nil {
		return err
	}
	err = s.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}
	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second
	return nil
}

func newRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.LoadHTMLGlob("templates/**/*")
	r.Static("/static", "./static")
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))
	routers.ShowRouterInit(r)
	routers.BackRouterInit(r)
	return r
}
