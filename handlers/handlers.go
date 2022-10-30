package handlers

import (
	"fmt"
	"io"
	"net/http"
)

// todo: helloのハンドラ
func HelloHandler(w http.ResponseWriter, req *http.Request) {
	// todo: reqの中のMethodフィールドがGETだったら
	if req.Method == http.MethodGet {
		io.WriteString(w, "Hello, world!\n")
		// reqの中のMethodフィールドがGETでなかったら
	} else {
		//Invalid method というレスポンスを405ステータスコードとともに返す
		http.Error(w, "Invalid method!", http.StatusMethodNotAllowed)
	}
}

// todo: articleのハンドラ
func PostArticleHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		io.WriteString(w, "Posting Article...\n")
	} else {
		http.Error(w, "Invalid method!", http.StatusMethodNotAllowed)
	}
}

func ArticleListHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		io.WriteString(w, "Article List\n")
	} else {
		http.Error(w, "Invalid method!", http.StatusMethodNotAllowed)
	}
}

func ArticleNumberHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		articleID := 1
		resString := fmt.Sprintf("Article No.%d\n", articleID)
		io.WriteString(w, resString)
	} else {
		http.Error(w, "Invalid method!", http.StatusMethodNotAllowed)
	}
}

func PostNiceHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		io.WriteString(w, "Posting Nice...\n")
	} else {
		http.Error(w, "Invalid method!", http.StatusMethodNotAllowed)
	}
}

func CommentHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		io.WriteString(w, "Posting Comment...\n")
	} else {
		http.Error(w, "Invalid method!", http.StatusMethodNotAllowed)
	}
}
