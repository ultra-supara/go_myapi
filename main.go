package main

import (
	"log"
	"net/http"

	"github.com/yourname/reponame/handlers"
)

func main() {
	//定義したハンドラをサーバーで使うことができるように登録する
	http.HandleFunc("/hello", handlers.HelloHandler)
	http.HandleFunc("/article", handlers.PostArticleHandler)
	http.HandleFunc("/article/list", handlers.ArticleListHandler)
	http.HandleFunc("/article/1", handlers.ArticleNumberHandler)
	http.HandleFunc("/article/nice", handlers.PostNiceHandler)
	http.HandleFunc("/comment", handlers.CommentHandler)

	//サーバー起動時のログを出力する
	log.Println("server start at port 8080")

	//ListenAndServeによってサーバーを起動する
	log.Fatal(http.ListenAndServe(":8080", nil))
}
