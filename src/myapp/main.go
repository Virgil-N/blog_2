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
	r.HandleFunc("/getArticle", controllers.GetArticle)

	http.Handle("/", r)

	if err := http.ListenAndServe("localhost:9090", nil); err != nil {
		log.Fatalln("ListenAndServe: ", err)
	}
}
