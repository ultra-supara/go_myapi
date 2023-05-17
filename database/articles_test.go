//todo: 投稿IDを指定して、記事データを取得する関数 SelectArticleDetail が正しく動いているのかを確認するunit test
/*  1. すでにdatabaseに入っている記事のデータを1つ選択する
    2. 選択した記事のデータをSelectArticleDetailでdbから取得する
    3. 取得したデータと、選択した記事のデータが一致するかを確認する
*/

// Compare this snippet from go_myapi/database/articles_test.go:

package database_test

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/ultra-supara/go_myapi/database"
	"github.com/ultra-supara/go_myapi/models"

	_ "github.com/go-sql-driver/mysql"
)

func TestSelectArticleDetail(t *testing.T) {
	//todo: データベースに接続する
	dbuser := "docker"
	dbpassword := "docker"
	dbdatabase := "sampledb"
	dbConn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true", dbuser, dbpassword, dbdatabase)

	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		t.Fatal(err) //接続できなかったときFatalでtestを終了させる
	}
	defer db.Close()

	//todo: test結果として期待する値を定義する
	expected := models.Article{
		ID:        1,
		Title:     "first post",
		Contents:  "this is my first blog",
		UserName:  "ultra-supara",
		NiceNum:   2,
	}

	//todo: test対象となる関数を実行する , 結果は変数gotに格納
	got, err := database.SelectArticleDetail(db, expected.ID)
	if err != nil {
		//関数の実行そのものに失敗した場合はtestも失敗させる
		t.Fatal(err)
	}

	//todo: 期待する出力と実際の出力が一致するかを確認する
	if got.ID != expected.ID {
		// 一致しない場合はtest失敗
		t.Errorf("ID: get %d, want %d", got.ID, expected.ID)
	}

	if got.Title != expected.Title {
		t.Errorf("Title: get %s, want %s", got.Title, expected.Title)
	}

	if got.Contents != expected.Contents {
		t.Errorf("Contents: get %s, want %s", got.Contents, expected.Contents)
	}

	if got.UserName != expected.UserName {
		t.Errorf("UserName: get %s, want %s", got.UserName, expected.UserName)
	}

	if got.NiceNum != expected.NiceNum {
		t.Errorf("NiceNum: get %d, want %d", got.NiceNum, expected.NiceNum)
	}

	// t.Fatal , t.Errorf ともに実行されずに関数が終了した時はtestに成功したことになる
}
