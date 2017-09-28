package controllers

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"html/template"
	// "myapp/models"
	// "fmt"
	"net/http"
)

type Result struct {
	Title      string `json:"title"`
	Category   string `json:"category"`
	Banner_url string `json:"banner_url"`
	Content    string `json:"content"`
	Created    string `json:"created"`
	AuthorName string `json:"authorName"`
}

func GetHome(w http.ResponseWriter, r *http.Request) {

	db, err := sql.Open("mysql", "root:123456@/blog_1")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// stmt, err := db.Prepare("insert article (title, category, banner_url, content, author_position, athor_id) values (?, ?, ?, ?, ?, ?)")
	// if err != nil {
	// 	panic(err)
	// }

	// res, err := stmt.Exec("金银岛", "文学", "/images/home/article_1.jpg", "文章占座...", "浙江省杭州市拱墅区灯彩街22号", 1)
	// if err != nil {
	// 	panic(err)
	// }

	// 内连接
	rows, err := db.Query("select article.id, article.title, article.category, article.banner_url, article.content, article.created, author.name from article inner join author on article.author_id = author.id;")
	if err != nil {
		panic(err)
	}

	// 这种写法太长了
	// for rows.Next() {
	// 	var result models.Article
	// 	err := rows.Scan(&result.Id, &result.Title, &result.Category, &result.Banner_url, &result.Content, &result.Created, &result.Author_position, &result.Author_id)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	fmt.Println(result)
	// }

	// 这个比较好
	columns, _ := rows.Columns()
	scanArgs := make([]interface{}, len(columns))
	values := make([]interface{}, len(columns))
	var result []interface{}

	for i := range values {
		scanArgs[i] = &values[i]
	}

	for rows.Next() {
		err := rows.Scan(scanArgs...)
		if err != nil {
			panic(err)
		}
		record := make(map[string]string)
		for k, v := range values {
			if v != nil {
				record[columns[k]] = string(v.([]byte))
			}
		}
		result = append(result, record)
	}

	t := template.Must(template.ParseFiles("templates/home.html", "templates/common/header.html", "templates/common/footer.html", "templates/common/sidebar.html"))
	t.ExecuteTemplate(w, "home", result)
}
