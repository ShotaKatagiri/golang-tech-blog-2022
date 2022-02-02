package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go-tech-blog/handler"
	"go-tech-blog/repository"
)

type articles struct {
	ID      uint `gorm:"primary_key"`
	Article string
	A       int
	V       int
}

var e = createMux()

func main() {
	db := gormConnect()
	repository.SetDB(db)
	// 構造体のインスタンス化
	//articleEx := article{}
	//articleEx.Id = 0
	//articleEx.Article = "test"
	//articleEx.A = 22
	//articleEx.V = 3
	// INSERTを実行
	//result := db.Find(&articleEx)

	//db.Create(&articleEx)
	//repository.SetDB(db)

	e.GET("/", handler.ArticleIndex)
	e.GET("/new", handler.ArticleNew)
	e.GET("/:id", handler.ArticleShow)
	e.GET("/:id/edit", handler.ArticleEdit)

	e.Logger.Fatal(e.Start("localhost:8080"))

}

func createMux() *echo.Echo {
	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.Gzip())

	e.Static("/css", "src/css")
	e.Static("/js", "src/js")
	return e
}

func gormConnect() *gorm.DB {

	DBMS := "mysql"
	USER := "root"
	PASS := ""
	PROTOCOL := "tcp(127.0.0.1:3306)"
	DBNAME := "golang"

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME
	db, err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		panic(err.Error())
	}

	return db
}
