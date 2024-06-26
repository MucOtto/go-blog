package router

import (
	"go-blog/api"
	"go-blog/views"
	"net/http"
)

func Router() {
	http.HandleFunc("/", views.HTML.Index)
	http.HandleFunc("/c/", views.HTML.Category)
	http.HandleFunc("/p/", views.HTML.Detail)
	http.HandleFunc("/login", views.HTML.Login)
	http.HandleFunc("/writing", views.HTML.Writing)
	http.HandleFunc("/api/v1/login", api.Api.Login)
	http.HandleFunc("/api/v1/post", api.Api.SavePost)
	http.HandleFunc("/api/v1/post/", api.Api.GetPost)
	http.Handle("/resource/", http.StripPrefix("/resource/", http.FileServer(http.Dir("public/resource/"))))
}
