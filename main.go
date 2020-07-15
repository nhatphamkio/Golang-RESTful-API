package main
import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
)
type Article struct{
	Title string `json:"Title"`
	Desc string `json:"dess"`
	Content string `json:"content"`
}
type Articles []Article
func allArticles(w http.ResponseWriter,r *http.Request){
	articles := Articles{
		Article{Title:"Test Tittle",Desc:"Test Description",Content:"Hello"},
	}
	fmt.Println("Endpoint Hit:All Articles Endpoint")
	json.NewEncoder(w).Encode(articles)
}
func homePage(w http.ResponseWriter,r *http.Request){
	fmt.Fprintf(w,"Homepage Endpoint Hit")
}
func testPostArticles(w http.ResponseWriter,r *http.Request){
	fmt.Fprintf(w,"Test POST endpoint worker")
}
func handleRequest(){
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/",homePage)
	myRouter.HandleFunc("/articles",allArticles).Methods("GET")
	myRouter.HandleFunc("/articles",testPostArticles).Methods("POST")
	log.Fatal(http.ListenAndServe(":8001",myRouter))

}
func main(){
	handleRequest()
}