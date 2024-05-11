package views

import (
	"go-blog/common"
	"go-blog/service"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func (*HtmlApi) Category(w http.ResponseWriter, r *http.Request) {
	categoryTemplate := common.Template.Category

	path := r.URL.Path
	cid := strings.TrimPrefix(path, "/c/")
	cidInt, err := strconv.Atoi(cid)
	if err != nil {
		categoryTemplate.WriteError(w, err)
	}

	err = r.ParseForm()
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

	data, err := service.GetCategoryInfo(cidInt, page, pageSize)
	if err != nil {
		categoryTemplate.WriteError(w, err)
	}

	categoryTemplate.WriteData(w, data)
}
