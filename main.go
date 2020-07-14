package main
import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
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
func handleRequest(){
	http.HandleFunc("/",homePage)
	http.HandleFunc("/articles",allArticles)
	log.Fatal(http.ListenAndServe(":8001",nil))

}
func main(){
	handleRequest()
}