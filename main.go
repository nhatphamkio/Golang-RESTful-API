package main
import (
	"fmt"
	"log"
	"net/http"
)
type 
func homePage(w http.ResponseWriter,r *http.Request){
	fmt.Fprintf(w,"Homepage Endpoint Hit")
}
func handleRequest(){
	http.HandleFunc("/",homePage)
	log.Fatal(http.ListenAndServe(":8001",nil))

}
func main(){
	handleRequest()
}