package router

import (
	"net/http"
	"runtime/debug"
	"time"

	log "github.com/go-kit/kit/log"
)

type Middleware func(http.HandlerFunc) http.HandlerFunc

func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next http.HandlerFunc) http.HandlerFunc {
		fn := func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if err := recover(); err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					logger.Log(
						"err", err,
						"trace", debug.Stack(),
					)
				}
			}()

			start := time.Now()
			wrapped := wrapResponseWriter(w)
			next.ServeHTTP(wrapped, r)
			logger.Log("status", wrapped.status,
				"method", r.Method,
				"path", r.URL.EscapedPath(),
				"duration", time.Since(start),
			)
		}

		return http.HandlerFunc(fn)
	}
}

//func LoggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
//
//	return func(w http.ResponseWriter, r *http.Request) {
//		fmt.Println("Request received:", r.Method, "in link:", r.URL.Path, "\n from IP:", r.RemoteAddr)
//		next.ServeHTTP(w, r)
//	}
//
//}

//func MethodNotAllowedMiddleware(next http.HandlerFunc, allowedMethods string) http.HandlerFunc {
//	return func(w http.ResponseWriter, r *http.Request) {
//		methodallowed := false
//		for _, method := range allowedMethods {
//			if r.Method == method {
//				methodallowed = true
//				break
//			}
//		}
//		if !methodallowed {
//			http.Error(w, "Method not Allowed", http.StatusMethodNotAllowed)
//		}
//		next.ServeHTTP(w, r)
//	}
//}
