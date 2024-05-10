package service

import (
	"go-blog/config"
	"go-blog/dao"
	"go-blog/models"
	"html/template"
)

func GetHomeInfo(page, pageSize int) (*models.HomeRes, error) {

	//页面上涉及到的所有的数据，必须有定义
	category, err := dao.QueryAllCategory()
	if err != nil {
		return nil, err
	}
	posts, err := dao.GetPost(page, pageSize)
	if err != nil {
		return nil, err
	}

	var postMores []models.PostMore
	for _, post := range posts {
		categoryName := dao.GetCategoryNameById(post.CategoryId)
		userName := dao.GetUsernameById(post.UserId)
		content := []rune(post.Content)
		if len(content) > 100 {
			content = content[:200]
		}
		postMore := models.PostMore{
			Pid:          post.Pid,
			Title:        post.Title,
			Slug:         post.Slug,
			Content:      template.HTML(content),
			CategoryId:   post.CategoryId,
			CategoryName: categoryName,
			UserId:       post.UserId,
			UserName:     userName,
			ViewCount:    post.ViewCount,
			Type:         post.Type,
			CreateAt:     models.DateDay(post.CreateAt),
			UpdateAt:     models.DateDay(post.UpdateAt),
		}
		postMores = append(postMores, postMore)
	}

	total := dao.GetPostCount()
	pageCount := (total-1)/10 + 1
	var pages []int
	for i := 0; i < pageCount; i++ {
		pages = append(pages, i+1)
	}

	homeRes := &models.HomeRes{
		Viewer:    config.Cfg.Viewer,
		Categorys: category,
		Posts:     postMores,
		Total:     total,
		Page:      page,
		Pages:     pages,
		PageEnd:   page != pageCount,
	}

	return homeRes, nil
}
