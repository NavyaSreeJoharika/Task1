package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Article struct {
	Title   string `json:"Title"`
	Desc    string `json:"Desc"`
	Content string `json:"Content"`
}

type Articles []Article

var articles = Articles{
	Article{Title: "Article 1", Desc: "Test Description", Content: "Hello World"},
	Article{Title: "Article 2", Desc: "Plants", Content: "Global Warming "},
}

func allArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: All Articles Endpoint")
	json.NewEncoder(w).Encode(articles)
}

func testPostArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Test PATCH Endpoint worked")
}

func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage Endpoint Hit")
}

func specificArticles(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for _, item := range articles {
		if item.Title == params["Title"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Article{})
}

func DeleteArticles(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	title := vars["Title"]

	for index, article := range articles {
		if article.Title == title {
			articles = append(articles[:index], articles[index+1:]...)
			break
		}
	}
}

func CreateArticle(w http.ResponseWriter, req *http.Request) {
	var article Article
	_ = json.NewDecoder(req.Body).Decode(&article)
	articles = append(articles, article)
	fmt.Println("Endpoint Hit: Create Article Endpoint")
	json.NewEncoder(w).Encode(article)
}

func handleRequests() {

	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", homepage)
	myRouter.HandleFunc("/articles", allArticles).Methods("GET")
	myRouter.HandleFunc("/articles/{Title}", specificArticles).Methods("GET")
	myRouter.HandleFunc("/articles", testPostArticles).Methods("PATCH")
	myRouter.HandleFunc("/articles", CreateArticle).Methods("POST")
	myRouter.HandleFunc("/articles/{Title}", DeleteArticles).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8081", myRouter))

}

func main() {
	handleRequests()
}
