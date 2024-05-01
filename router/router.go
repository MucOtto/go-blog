package router

import (
	"go-blog/views"
	"net/http"
)

func Router() {
	http.HandleFunc("/", views.HTML.Index)
	http.Handle("/resource/", http.StripPrefix("/resource/", http.FileServer(http.Dir("public/resource/"))))
}
