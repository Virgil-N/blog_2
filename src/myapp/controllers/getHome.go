package controllers

import (
	rt "gopkg.in/gorethink/gorethink.v3"
	"html/template"
	// "myapp/models"
	"fmt"
	"log"
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

	session, err := rt.Connect(rt.ConnectOpts{
		Address: "localhost:28015",
	})
	if err != nil {
		log.Fatalln(err)
	}
	defer session.Close()

	// 明天把author表移到blog数据库中
	res, err := rt.DB("test").Table("author").Filter(map[string]interface{}{
		"name": "宋江",
	}).Run(session)
	if err != nil {
		panic(err)
		return
	}
	defer res.Close()

	var results []interface{}
	err = res.All(&results)
	if err != nil {
		panic(err)
		return
	}
	fmt.Printf("%d results", len(results))

	t := template.Must(template.ParseFiles("templates/home.html", "templates/common/header.html", "templates/common/footer.html", "templates/common/sidebar.html"))
	t.ExecuteTemplate(w, "home", nil)
}
