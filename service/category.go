package service

import (
	"go-blog/config"
	"go-blog/dao"
	"go-blog/models"
	"html/template"
)

func GetCategoryInfo(cid int, page int, pageSize int) (*models.CategoryResponse, error) {

	category, err := dao.QueryAllCategory()
	if err != nil {
		return nil, err
	}
	categoryName := dao.GetCategoryNameById(cid)

	posts, err := dao.GetPostsByCategoryId(cid, page, pageSize)
	if err != nil {
		return nil, err
	}

	var postMores []models.PostMore
	for _, post := range posts {
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

	total := dao.GetPostCountByCategoryId(cid)
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

	categoryResponse := &models.CategoryResponse{
		HomeRes:      *homeRes,
		CategoryName: categoryName,
	}

	return categoryResponse, nil
}
