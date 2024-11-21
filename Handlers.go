package main

import (
	router "Web-Router/Router"
	"net/http"
)

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	message := "Default message"
	if code, ok := r.Context().Value("code").(int); ok {
		message = router.MessageCodes[code]
	}
	w.Write([]byte(message))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	message := "This is the default page"

	w.Write([]byte(message))
}
