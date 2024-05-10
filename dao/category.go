package dao

import (
	"go-blog/models"
	"log"
)

func QueryAllCategory() ([]models.Category, error) {
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

func GetCategoryNameById(id int) string {
	row := DB.QueryRow("select name from blog_category where cid = ?", id)
	if row.Err() != nil {
		log.Println(row.Err().Error())
	}
	var name string
	err := row.Scan(&name)
	if err != nil {
		return "未知"
	}
	return name
}
