package main

import (
	"net/http"

	router "Web-Router/Router"
)

func main() {

	r := &router.Router{}

	r.Route(http.MethodGet, "/about", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("The Best Router!"))
	})
	http.ListenAndServe(":8000", r)
}
