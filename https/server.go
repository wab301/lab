package main

import (
	"fmt"
	"net/http"
)

/*
https://localhost:9090/
or curl -k https://localhost:9090/

*/

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi, This is an example of https service in golang! URL.Path = %q\n", r.URL.Path)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServeTLS(":9090", "server.crt", "server.key", nil)
}
