package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/ultra-supara/go_myapi/models"
)

// todo: helloのハンドラ
func HelloHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello World!\n")
}

// todo: articleのハンドラ
func PostArticleHandler(w http.ResponseWriter, req *http.Request) {
	//io.WriteString(w, "Posting article\n")
	article := models.Article1
	jsonData, err := json.Marshal(article)
	if err != nil {
		http.Error(w, "fail to encode json\n", http.StatusInternalServerError)
		return
	}
	w.Write(jsonData)
}

// todo: article/listのハンドラ , クエリパラメタ取得機能を実装する
func ArticleListHandler(w http.ResponseWriter, req *http.Request) {
	queryMap := req.URL.Query()

	var page int
	if p, ok := queryMap["page"]; ok && len(p) > 0 {
		var err error
		page, err = strconv.Atoi(p[0])
		if err != nil {
			http.Error(w, "invalid query parameter\n", http.StatusBadRequest)
			return
		}
	} else {
		page = 1
	}
	//resString := fmt.Sprintf("Article List (page No.)%d\n", page)
	//io.WriteString(w, resString)
	articleLList := []models.Article{models.Article1, models.Article2}
	jsonData, err := json.Marshal(articleLList)
	if err != nil {
		errMsg := fmt.Sprintf("fail to encode json (page No.%d)\n", page)
		http.Error(w, errMsg, http.StatusInternalServerError)
		return
	}
	w.Write(jsonData)
}

// todo: GET /article/{id} のハンドラ
func ArticleNumberHandler(w http.ResponseWriter, req *http.Request) {
	articleID, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		http.Error(w, "invalid query parameter\n", http.StatusBadRequest)
		return
	}
	article := models.Article1
	jsonData, err := json.Marshal(article)
	if err != nil {
		errMsg := fmt.Sprintf("fail to encode json (articleID No.%d)\n", articleID)
		http.Error(w, errMsg, http.StatusInternalServerError)
		return
	}
	//resString := fmt.Sprintf("Article No.%d\n", articleID)
	//io.WriteString(w, resString)
	w.Write(jsonData)
}

// todo: POST /article/nice のハンドラ
func PostNiceHandler(w http.ResponseWriter, req *http.Request) {
	article := models.Article1
	jsonData, err := json.Marshal(article)
	if err != nil {
		http.Error(w, "fail to encode json\n", http.StatusInternalServerError)
		return
	}
	w.Write(jsonData)
}

// todo: POST /comment のハンドラ
func CommentHandler(w http.ResponseWriter, req *http.Request) {
	comment := models.Comment1
	jsonData, err := json.Marshal(comment)
	if err != nil {
		http.Error(w, "fail to encode json\n", http.StatusInternalServerError)
		return
	}
	w.Write(jsonData)
}
