package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/kult0922/idea-note/api"
)

func loadEnv() {
	// ここで.envファイル全体を読み込みます。
	// この読み込み処理がないと、個々の環境変数が取得出来ません。
	// 読み込めなかったら err にエラーが入ります。
	err := godotenv.Load(".env")

	// もし err がnilではないなら、"読み込み出来ませんでした"が出力されます。

	if err != nil {
		fmt.Printf("読み込み出来ませんでした: %v", err)
	}
}

var ()

func main() {
	loadEnv()
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbDatabase := os.Getenv("DB_NAME")
	locale, _ := time.LoadLocation("Asia/Tokyo")
	c := mysql.Config{
		DBName:               dbDatabase,
		User:                 dbUser,
		Passwd:               dbPassword,
		Addr:                 "127.0.0.1:3306",
		Net:                  "tcp",
		Collation:            "utf8mb4_bin",
		ParseTime:            true,
		Loc:                  locale,
		AllowNativePasswords: true,
	}
	dsn := c.FormatDSN()
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Println("fail to connect DB")
		return
	}

	r := api.NewRouter(db)

	log.Println("server start at port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
