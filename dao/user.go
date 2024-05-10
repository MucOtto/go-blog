package dao

import "log"

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
