package controllers

import (
	rt "gopkg.in/gorethink/gorethink.v3"
	"html/template"
	// "myapp/models"
	// "fmt"
	"net/http"
)

type ArticleResult struct {
	Id             string `json:"id"`
	Title          string `json:"title"`
	Category       string `json:"category"`
	BannerUrl      string `json:"bannerUrl"`
	Content        string `json:"content"`
	Created        string `json:"created"`
	AuthorName     string `json:"authorName"`
	AuthorPosition string `json:"authorPosition"`
}

func GetHome(w http.ResponseWriter, r *http.Request) {

	session, err := rt.Connect(rt.ConnectOpts{
		Address: "localhost:28015",
	})
	if err != nil {
		panic(err)
		return
	}
	defer session.Close()

	res, err := rt.DB("blog").Table("article").Filter("").Run(session)
	if err != nil {
		panic(err)
		return
	}
	defer res.Close()

	var articleResults []ArticleResult
	err = res.All(&articleResults)
	if err != nil {
		panic(err)
		return
	}

	t := template.Must(template.ParseFiles("templates/home.html", "templates/common/header.html", "templates/common/footer.html", "templates/common/sidebar.html"))
	t.ExecuteTemplate(w, "home", articleResults)
}
