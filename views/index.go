package views

import (
	"errors"
	"go-blog/common"
	"go-blog/service"
	"log"
	"net/http"
)

func (*HtmlApi) Index(w http.ResponseWriter, r *http.Request) {
	_index := common.Template.Index

	res, err := service.GetHomeInfo()
	if err != nil {
		log.Println("error:", err.Error())
		_index.WriteError(w, errors.New("系统错误"))
	}

	_index.WriteData(w, res)
}
