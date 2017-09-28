package controllers

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"html/template"
	// "myapp/models"
	"net/http"
)

// type Result struct {
// 	Title      string `json:"title"`
// 	Category   string `json:"category"`
// 	Banner_url string `json:"banner_url"`
// 	Content    string `json:"content"`
// 	Created    string `json:"created"`
// 	AuthorName string `json:"authorName"`
// }

func GetArticle(w http.ResponseWriter, r *http.Request) {

	articleId := r.FormValue("articleId")

	db, err := sql.Open("mysql", "root:123456@/blog_1")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// 内连接
	row := db.QueryRow("select article.title, article.category, article.banner_url, article.content, article.created, author.name from article inner join author on article.author_id = author.id and article.id = ?;", articleId)

	var result Result
	err = row.Scan(&result.Title, &result.Category, &result.Banner_url, &result.Content, &result.Created, &result.AuthorName)
	if err != nil {
		panic(err)
	}

	t := template.Must(template.ParseFiles("templates/article.html", "templates/common/header.html", "templates/common/footer.html", "templates/common/sidebar.html"))
	t.ExecuteTemplate(w, "article", result)
}
