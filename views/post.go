package views

import (
	"errors"
	"go-blog/common"
	"go-blog/service"
	"net/http"
	"strconv"
	"strings"
)

func (*HtmlApi) Detail(w http.ResponseWriter, r *http.Request) {
	detailTemplate := common.Template.Detail

	path := r.URL.Path
	if path == "" {
		detailTemplate.WriteError(w, errors.New("路径读取失败"))
	}

	pidStr := strings.TrimPrefix(path, "/p/")
	pidStr = strings.TrimSuffix(pidStr, ".html")
	pid, err := strconv.Atoi(pidStr)
	if err != nil {
		detailTemplate.WriteError(w, err)
	}

	data, err := service.GetPostDetailById(pid)
	if err != nil {
		detailTemplate.WriteError(w, err)
	}

	detailTemplate.WriteData(w, data)
}
