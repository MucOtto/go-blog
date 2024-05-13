package dao

import (
	"go-blog/models"
	"log"
)

func GetPost(page, pageSize int) ([]models.Post, error) {
	start := (page - 1) * pageSize
	rows, err := DB.Query("select * from blog_post limit ?,?", start, pageSize)
	if err != nil {
		return nil, err
	}
	var posts []models.Post
	for rows.Next() {
		var post models.Post
		err := rows.Scan(&post.Pid,
			&post.Title,
			&post.Content,
			&post.Markdown,
			&post.CategoryId,
			&post.UserId,
			&post.ViewCount,
			&post.Type,
			&post.Slug,
			&post.CreateAt,
			&post.UpdateAt,
		)
		if err != nil {
			log.Println("post数据库读取失败", err.Error())
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func GetPostCount() (count int) {
	row := DB.QueryRow("select count(1) from blog_post")
	if row.Err() != nil {
		log.Println(row.Err().Error())
	}
	err := row.Scan(&count)
	if err != nil {
		return 0
	}
	return count
}

func GetPostsByCategoryId(cid int, page int, pageSize int) ([]models.Post, error) {
	start := (page - 1) * pageSize

	rows, err := DB.Query("select * from blog_post where category_id = ? limit ?,?", cid, start, pageSize)
	if err != nil {
		log.Println("根据类别id查询文章失败L", err.Error())
	}

	var posts []models.Post
	for rows.Next() {
		var post models.Post
		err := rows.Scan(&post.Pid,
			&post.Title,
			&post.Content,
			&post.Markdown,
			&post.CategoryId,
			&post.UserId,
			&post.ViewCount,
			&post.Type,
			&post.Slug,
			&post.CreateAt,
			&post.UpdateAt,
		)
		if err != nil {
			log.Println("post数据库读取失败", err.Error())
		}
		posts = append(posts, post)
	}

	return posts, nil
}

func GetPostCountByCategoryId(cid int) (count int) {
	err := DB.QueryRow("select count(1) from blog_post where category_id = ?", cid).Scan(&count)
	if err != nil {
		return 0
	}
	return count
}

func GetPostDetailById(pid int) (*models.Post, error) {
	post := &models.Post{}
	err := DB.QueryRow("select * from blog_post where pid = ?", pid).Scan(
		&post.Pid,
		&post.Title,
		&post.Content,
		&post.Markdown,
		&post.CategoryId,
		&post.UserId,
		&post.ViewCount,
		&post.Type,
		&post.Slug,
		&post.CreateAt,
		&post.UpdateAt,
	)

	if err != nil {
		return nil, err
	}

	return post, nil
}
