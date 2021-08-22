/**
 * @Author:      zhangyuyao
 * @Description: TODO
 * @File:        main.go
 * @Version:     1.0.0
 * @Date:        2021/8/22 12:06 下午
 */

package main

import (
	"blog-service/internal/routers"
	"net/http"
	"time"
)

func main() {
	router := routers.NewRouter()
	s := &http.Server{
		Addr:           ":8081",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
