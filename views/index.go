package views

import (
	"go-blog/config"
	"go-blog/models"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"
)

var (
	path, _ = os.Getwd()
	PREFIX  = path + "/template"
)

func IsODD(num int) bool {
	return num%2 == 0
}
func GetNextName(strs []string, index int) string {
	return strs[index+1]
}
func Date(layout string) string {
	return time.Now().Format(layout)
}

func (*HtmlApi) Index(w http.ResponseWriter, r *http.Request) {

	t := template.New("index.html")

	index := PREFIX + "/index.html"
	header := PREFIX + "/layout/header.html"
	footer := PREFIX + "/layout/footer.html"
	personal := PREFIX + "/layout/personal.html"
	post_list := PREFIX + "/layout/post-list.html"
	pagination := PREFIX + "/layout/pagination.html"
	home := PREFIX + "/home.html"

	t.Funcs(template.FuncMap{"isODD": IsODD, "getNextName": GetNextName, "date": Date})
	t, err := t.ParseFiles(index, header, footer, personal, post_list, pagination, home)
	if err != nil {
		log.Println(" T -------------")
		log.Fatalln(err)
	}
	//页面上涉及到的所有的数据，必须有定义
	var categorys = []models.Category{
		{
			Cid:  1,
			Name: "go",
		},
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
		Categorys: categorys,
		Posts:     posts,
		Total:     1,
		Page:      1,
		Pages:     []int{1},
		PageEnd:   true,
	}

	t.Execute(w, *homeRes)
}
