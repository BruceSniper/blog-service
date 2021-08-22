/**
 * @Author:      zhangyuyao
 * @Description: TODO
 * @File:        main.go
 * @Version:     1.0.0
 * @Date:        2021/8/22 12:06 下午
 */

package main

import (
	"blog-service/global"
	"blog-service/internal/routers"
	"blog-service/pkg/setting"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func init() {
	err := setupSetting()
	if err != nil {
		log.Fatalf("init setupSetting err: %v", err)
	}
}

func setupSetting() error {
	setting2, err := setting.NewSetting()
	if err != nil {
		return err
	}

	err = setting2.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}

	err = setting2.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}

	err = setting2.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}

	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second
	return nil
}

func main() {
	gin.SetMode(global.ServerSetting.RunMode)
	router := routers.NewRouter()
	s := &http.Server{
		Addr:           ":" + global.ServerSetting.HttpPort,
		Handler:        router,
		ReadTimeout:    global.ServerSetting.ReadTimeout,
		WriteTimeout:   global.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
