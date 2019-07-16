//passing data through the url

package main

import (
	"io"
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

	//GET method posts watever you type in the input box both on the URL n the request body
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `
	<form method="GET">
	 <input type="text" name="q">
	 <input type="submit">
	</form>
	<br>`+v)
}

//visit this page
//// http://localhost:8080

//POST --> body
//GET --> URL
