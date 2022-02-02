package repository

import (
	"github.com/jinzhu/gorm"
	"go-tech-blog/model"
)

// ArticleList ...
func ArticleList() ([]*model.Article, *gorm.DB) {
	// 実行する SQL

	//query := `SELECT * FROM articles;`

	// データベースから取得した値を格納する変数を宣言
	var articles []*model.Article

	// Query を実行して、取得した値を変数に格納
	db.Find(&articles)
	//if err := db.Select(&articles, query); err != nil {
	//	fmt.Printf("azzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzz")
	//	return nil, nil
	//}

	return articles, nil
}
