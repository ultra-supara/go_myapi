package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ultra-supara/myapi/handlers"
)

func main() {
	//Router Rの明示的宣言
	r := mux.NewRouter()

	//定義したハンドラをサーバーで使うことができるように登録する
	r.HandleFunc("/hello", handlers.HelloHandler).Methods(http.MethodGet)

	r.HandleFunc("/article", handlers.PostArticleHandler).Methods(http.MethodPost)
	r.HandleFunc("/article/list", handlers.ArticleListHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/{id:[0-9]+}", handlers.ArticleNumberHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/nice", handlers.PostNiceHandler).Methods(http.MethodPost)

	r.HandleFunc("/comment", handlers.PostCommentHandler).Methods(http.MethodPost)

	//サーバー起動時のログを出力する
	log.Println("server start at port 8080")

	//ListenAndServeによってサーバーを起動する
	log.Fatal(http.ListenAndServe(":8080", r))
}
