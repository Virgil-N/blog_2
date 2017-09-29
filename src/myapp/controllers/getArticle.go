package controllers

import (
	"html/template"
	// "myapp/models"
	"fmt"
	rt "gopkg.in/gorethink/gorethink.v3"
	"net/http"
)

func GetArticle(w http.ResponseWriter, r *http.Request) {

	articleId := r.FormValue("articleId")

	session, err := rt.Connect(rt.ConnectOpts{
		Address: "localhost:28015",
	})
	if err != nil {
		panic(err)
		return
	}
	defer session.Close()

	res, err := rt.DB("blog").Table("article").Get(articleId).Run(session)
	if err != nil {
		panic(err)
		return
	}
	defer res.Close()

	var articleResult ArticleResult
	err = res.One(&articleResult)
	if err != nil {
		panic(err)
		return
	}
	fmt.Printf("articleResult: %v", articleResult)

	t := template.Must(template.ParseFiles("templates/article.html", "templates/common/header.html", "templates/common/footer.html", "templates/common/sidebar.html"))
	t.ExecuteTemplate(w, "article", articleResult)
}
