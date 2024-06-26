package main

import (
	"fmt"
	"go-blog/common"
	"go-blog/router"
	"net/http"
)

func init() {
	common.LoadTemplate()
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	// 路由
	router.Router()

	if err := server.ListenAndServe(); err != nil {
		fmt.Println(err)
	}
}
