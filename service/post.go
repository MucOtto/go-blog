package service

import (
	"go-blog/config"
	"go-blog/dao"
	"go-blog/models"
	"html/template"
	"log"
)

func GetPostDetailById(pid int) (*models.PostRes, error) {
	post, err := dao.GetPostDetailById(pid)
	if err != nil {
		return nil, err
	}

	categoryName := dao.GetCategoryNameById(post.CategoryId)
	username := dao.GetUsernameById(post.UserId)

	return &models.PostRes{
		Viewer:       config.Cfg.Viewer,
		SystemConfig: config.Cfg.System,
		Article: models.PostMore{
			Pid:          pid,
			Title:        post.Title,
			Slug:         post.Slug,
			Content:      template.HTML(post.Content),
			CategoryId:   post.CategoryId,
			CategoryName: categoryName,
			UserId:       post.UserId,
			UserName:     username,
			ViewCount:    post.ViewCount,
			Type:         post.Type,
			CreateAt:     models.DateDay(post.CreateAt),
			UpdateAt:     models.DateDay(post.UpdateAt),
		},
	}, nil
}

func Writing() (wr models.WritingRes) {
	wr.Title = config.Cfg.Viewer.Title
	wr.CdnURL = config.Cfg.System.CdnURL
	category, err := dao.QueryAllCategory()
	if err != nil {
		log.Println(err)
		return
	}
	wr.Categorys = category
	return
}

func SavePost(post *models.Post) {
	dao.SavePost(post)
}
