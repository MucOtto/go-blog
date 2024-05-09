package dao

import (
	"go-blog/models"
	"log"
)

func queryAllCategory() ([]models.Category, error) {
	rows, err := DB.Query("select * from blog_category")
	if err != nil {
		log.Panicln("queryAllCategory 失败", err.Error())
		return nil, err
	}
	var categories []models.Category
	for rows.Next() {
		var category models.Category
		err = rows.Scan(&category.Cid, &category.Name, &category.CreateAt, &category.UpdateAt)
		if err != nil {
			log.Println("category数据读取出错", err.Error())
			return nil, err
		}
		categories = append(categories, category)
	}
	return categories, nil
}
