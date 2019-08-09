package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"
)

//MyMiddleware depicts the middlware pattern
func MyMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//... TODO execute middleware logic here, will be executed on the way down

		//next is in closure, call the ServeHTTP function of the next handler in the row.
		next.ServeHTTP(w, r)

		//... TODO execute middleware logic here, will be executed on the way back(defered function will do as well)
	})
}

//RecoverPanic recovers from panic, and gives a meaningful message to the client
//... covers panics in all subsequent middleware and handlers
func (app *Application) RecoverPanic(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//... will always be run in the event of a panic, as Go unwinds the stack
		defer func() {
			if err := recover(); err != nil {
				w.Header().Set("Connection", "close")
				app.ServerError(w, fmt.Errorf("%s", err))
			}
		}()
		next.ServeHTTP(w, r)
	})
}

//LogRequest logs each request on the server
func (app *Application) LogRequest(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		app.InfoLogger.Printf("\n%s\nNew %s-request [%s], proto [%s], comming in from [%s]...", strings.Repeat("=", 50), r.Method, r.URL.RequestURI(), r.Proto, r.RemoteAddr)
		startReq := time.Now()
		next.ServeHTTP(w, r)
		endReq := time.Now()
		app.InfoLogger.Printf("request  from '%s' has been processed. Elapsed: %s\n%s", r.RemoteAddr, endReq.Sub(startReq), strings.Repeat("=", 50))
	})
}

//SecureHeaders adds some secure headers to help against XSS and Clickjacking attacks
func SecureHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-Frame-Options", "deny")
		w.Header().Set("X-Xss-Protection", "1;mode=block")
		next.ServeHTTP(w, r)
	})
}
