//passing data through the url

package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	//req.FormValue passes the data throgh the URL
	v := req.FormValue("q")
	fmt.Fprintln(w, "What's uppp ::"+v)
}

//visit this page
//// http://localhost:8080/?q=dog
