package main

import (
	"fmt"
	"net/http"

	v1 "github.com/yu1er/gin-blog/api/v1"
	"github.com/yu1er/gin-blog/config"
	"github.com/yu1er/gin-blog/router"
	"github.com/yu1er/gin-blog/service"
)

func init() {
	config.InitConfig()
	v1.InitApi()
	service.InitDB()
}

func main() {

	e := router.SetupRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", config.HTTPPort),
		Handler:        e,
		ReadTimeout:    config.ReadTimeout,
		WriteTimeout:   config.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}
