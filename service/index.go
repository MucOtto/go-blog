package service

import (
	"go-blog/config"
	"go-blog/dao"
	"go-blog/models"
)

func GetHomeInfo() (*models.HomeRes, error) {

	//页面上涉及到的所有的数据，必须有定义
	category, err := dao.QueryAllCategory()
	if err != nil {
		return nil, err
	}
	var posts = []models.PostMore{
		{
			Pid:          1,
			Title:        "go博客",
			Content:      "内容",
			UserName:     "张三",
			ViewCount:    123,
			CreateAt:     "2022-02-20",
			CategoryId:   1,
			CategoryName: "go",
			Type:         0,
		},
	}

	homeRes := &models.HomeRes{
		Viewer:    config.Cfg.Viewer,
		Categorys: category,
		Posts:     posts,
		Total:     1,
		Page:      1,
		Pages:     []int{1},
		PageEnd:   true,
	}

	return homeRes, nil
}
