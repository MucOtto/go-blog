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

func GetCategoryById(cid int) []models.Category {
	var category models.Category
	// 直接执行 Scan，在 Scan 中处理可能的错误
	err := DB.QueryRow("SELECT * FROM blog_category WHERE cid = ?", cid).Scan(
		&category.Cid,
		&category.Name,
		&category.CreateAt,
		&category.UpdateAt,
	)
	if err != nil {
		log.Println(err.Error())
		return nil // 或返回一个空切片，根据你的错误处理策略
	}

	// 直接在返回语句中使用构造的切片
	return []models.Category{category}
}
