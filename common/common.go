package common

import (
	"fmt"
	"go-blog/config"
	"go-blog/models"
)

var Template models.HtmlTemplate

func LoadTemplate() {
	var err error
	Template, err = models.InitTemplate(config.Cfg.System.CurrentDir + "/template/")
	if err != nil {
		fmt.Println("Load Error: ", err)
	}
}
