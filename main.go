package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type IndexData struct {
	Title string `json:"title"`
	Desc  string `json:"desc"`
}

func index(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	IndexData := IndexData{
		Title: "blog",
		Desc:  "hello world!",
	}
	data, _ := json.Marshal(IndexData)
	if _, err := w.Write(data); err != nil {
		fmt.Println(err)
	}

}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.HandleFunc("/", index)

	if err := server.ListenAndServe(); err != nil {
		fmt.Println(err)
	}
}
