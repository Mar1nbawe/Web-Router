package main

import (
	"github.com/go-kit/kit/log"
	stdlog "log"
	"net/http"
	"os"

	router "Web-Router/Router"
)

func main() {
	var logger log.Logger
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	stdlog.SetOutput(log.NewStdlibAdapter(logger))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC, "loc", log.DefaultCaller)

	r := &router.Router{}

	r.Use(router.LoggingMiddleware(logger), router.CorsMiddleware())
	r.Route(http.MethodGet, "/home", homeHandler, 1)
	r.Route(http.MethodGet, "/about", aboutHandler, 0)
	http.ListenAndServe("localhost:8000", r)

}
