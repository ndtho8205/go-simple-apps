package middlewares

import (
	"log"
	"net/http"
	"time"
)

type Middleware func(handlerFunc http.HandlerFunc) http.HandlerFunc

func LoggingMiddleware() Middleware {
	return func(handlerFunc http.HandlerFunc) http.HandlerFunc {
		return func(writer http.ResponseWriter, request *http.Request) {
			start := time.Now()
			defer func() {
				log.Println(request.URL.Path, time.Since(start))
			}()

			handlerFunc(writer, request)
		}
	}
}

func MethodMiddleware(method string) Middleware {
	return func(handlerFunc http.HandlerFunc) http.HandlerFunc {
		return func(writer http.ResponseWriter, request *http.Request) {
			if request.Method != method {
				http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
				return
			}

			handlerFunc(writer, request)
		}
	}
}

func Chain(handlerFunc http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, m := range middlewares {
		handlerFunc = m(handlerFunc)
	}

	return handlerFunc
}
