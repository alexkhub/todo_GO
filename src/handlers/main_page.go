package handlers

import ("fmt"
	"net/http"
)

func MainPageHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "Main page")
}