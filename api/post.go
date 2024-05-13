package api

import (
	"go-blog/common"
	"go-blog/models"
	"go-blog/service"
	"go-blog/utils"
	"net/http"
	"strconv"
	"time"
)

func (*API) SavePost(w http.ResponseWriter, r *http.Request) {

	token := r.Header.Get("Authorization")
	_, claims, err := utils.ParseToken(token)
	if err != nil {
		common.Error(w, err)
	}
	uid := claims.Uid
	method := r.Method
	switch method {
	case http.MethodPost:
		param := common.GetRequestJsonParam(r)
		categoryId := param["categoryId"].(string)
		cid, _ := strconv.Atoi(categoryId)
		content := param["content"].(string)
		markdown := param["markdown"].(string)
		slug := param["slug"].(string)
		//_type := param["type"].(int)
		title := param["title"].(string)

		post := &models.Post{
			Pid:        -1,
			Title:      title,
			Slug:       slug,
			Content:    content,
			Markdown:   markdown,
			CategoryId: cid,
			UserId:     uid,
			ViewCount:  0,
			Type:       0,
			CreateAt:   time.Now(),
			UpdateAt:   time.Now(),
		}

		service.SavePost(post)
		common.Success(w, post)

	case http.MethodPut:
	}
}
