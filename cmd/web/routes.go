package main

import (
	"net/http"
)

//Routes ...
//Version 1:  func (app *Application) Routes() *http.ServeMux {
func (app *Application) Routes() http.Handler {
	mux := http.NewServeMux()
	// mux.HandleFunc("/snippet/create", env.createSnippet)
	mux.HandleFunc("/snippet/create", CreateSnippetWithClosure(app))

	// mux.HandleFunc("/snippet", env.showSnippet)
	mux.HandleFunc("/snippet", ShowSnippetWithClosure(app))

	mux.HandleFunc("/logo", DownLoadLogoHandler(app))

	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// mux.HandleFunc("/", env.home)
	mux.HandleFunc("/", HomeWithClosure(app))

	return app.RecoverPanic(app.LogRequest(SecureHeaders(mux)))
}
