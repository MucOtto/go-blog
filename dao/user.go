package dao

import (
	"go-blog/models"
	"log"
)

func GetUsernameById(id int) string {
	row := DB.QueryRow("select user_name from blog_user where uid = ?", id)
	if row.Err() != nil {
		log.Println(row.Err().Error())
	}

	var name string
	err := row.Scan(&name)
	if err != nil {
		return "匿名"
	}
	return name
}

func GetUser(username string, password string) *models.User {
	user := &models.User{}

	err := DB.QueryRow("select * from blog_user where user_name = ? and passwd = ?", username, password).Scan(
		&user.Uid,
		&user.UserName,
		&user.Passwd,
		&user.Avatar,
		&user.CreateAt,
		&user.UpdateAt,
	)
	if err != nil {
		return nil
	}

	return user
}
