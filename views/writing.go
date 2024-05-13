package views

import (
	"go-blog/common"
	"go-blog/service"
	"net/http"
)

func (*HtmlApi) Writing(w http.ResponseWriter, r *http.Request) {
	writing := common.Template.Writing
	wr := service.Writing()
	writing.WriteData(w, wr)
}
