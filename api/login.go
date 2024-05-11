package api

import (
	"go-blog/common"
	"go-blog/service"
	"net/http"
)

func (*API) Login(w http.ResponseWriter, r *http.Request) {

	param := common.GetRequestJsonParam(r)
	username := param["username"].(string)
	password := param["passwd"].(string)
	loginRes, err := service.GetLoginInfo(username, password)
	if err != nil {
		common.Error(w, err)
	}

	common.Success(w, loginRes)
}
