package common

import (
	"encoding/json"
	"go-blog/config"
	"go-blog/models"
	"io"
	"log"
	"net/http"
	"sync"
)

var Template models.HtmlTemplate

func LoadTemplate() {
	w := sync.WaitGroup{}
	w.Add(1)
	go func() {
		//耗时
		var err error
		Template, err = models.InitTemplate(config.Cfg.System.CurrentDir + "/template/")
		if err != nil {
			panic(err)
		}
		w.Done()
	}()
	w.Wait()
}

func Success(w http.ResponseWriter, data interface{}) {
	var result models.Result
	result.Code = 200
	result.Data = data
	result.Error = ""
	w.Header().Set("Content-Type", "application/json")
	resJson, _ := json.Marshal(result)
	_, err := w.Write(resJson)
	if err != nil {
		log.Println(err.Error())
	}
}

func Error(w http.ResponseWriter, err error) {
	var result models.Result
	result.Code = 400
	result.Data = nil
	result.Error = err.Error()
	w.Header().Set("Content-Type", "application/json")
	resJson, _ := json.Marshal(result)
	_, err = w.Write(resJson)
	if err != nil {
		log.Println(err.Error())
	}
}

func GetRequestJsonParam(r *http.Request) map[string]interface{} {
	var params map[string]interface{}
	body, _ := io.ReadAll(r.Body)
	_ = json.Unmarshal(body, &params)
	return params
}
