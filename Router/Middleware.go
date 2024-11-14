package router

import (
	"fmt"
	"net/http"
)

func LoggingMiddleware(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Request received:", r.Method, "in link:", r.URL.Path, "\n from IP:", r.RemoteAddr)
		next.ServeHTTP(w, r)
	}

}
