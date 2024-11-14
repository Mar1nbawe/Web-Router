package main

import (
	"fmt"
	"net/http"

	router "Web-Router/Router"
)

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("test"))
}

func main() {

	r := &router.Router{}

	r.Use(router.LoggingMiddleware)
	r.Route(http.MethodGet, "/about", aboutHandler)
	http.ListenAndServe(":8000", r)
	fmt.Println("Listening to inputs!")
}
