package views

import (
	"errors"
	"go-blog/common"
	"go-blog/service"
	"log"
	"net/http"
	"strconv"
)

func (*HtmlApi) Index(w http.ResponseWriter, r *http.Request) {
	_index := common.Template.Index

	err := r.ParseForm()
	if err != nil {
		log.Println("表单获取失败")
	}

	page := 1
	pageStr := r.Form.Get("page")
	if pageStr != "" {
		_page, err := strconv.Atoi(pageStr)
		if err != nil {
			log.Println(err.Error())
		}
		page = _page
	}

	pageSize := 10

	res, err := service.GetHomeInfo(page, pageSize)
	if err != nil {
		log.Println("error:", err.Error())
		_index.WriteError(w, errors.New("系统错误"))
	}

	_index.WriteData(w, res)
}
