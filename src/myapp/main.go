package main

import (
	"github.com/gorilla/mux"
	"log"
	"myapp/controllers"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	http.Handle("/javascript/", http.FileServer(http.Dir("public")))
	http.Handle("/css/", http.FileServer(http.Dir("public")))
	http.Handle("/images/", http.FileServer(http.Dir("public")))
	http.Handle("/vendor/", http.FileServer(http.Dir("public")))

	r.HandleFunc("/", controllers.GetHome)
	r.HandleFunc("/home", controllers.GetHome)
	r.HandleFunc("/getArticle", controllers.GetArticle)

	http.Handle("/", r)

	if err := http.ListenAndServe("localhost:9090", nil); err != nil {
		log.Fatalln("ListenAndServe: ", err)
	}
}

// package main

// import (
// 	// "fmt"
// 	r "gopkg.in/gorethink/gorethink.v3"
// 	"log"
// )

// func main() {

// 	session, err := r.Connect(r.ConnectOpts{
// 		Address: "localhost:28015",
// 	})
// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// 	err = r.DB("test").Table("test1").Insert(map[string]string{
// 		"id":    "123",
// 		"title": "world",
// 	}).Exec(session)

// }
