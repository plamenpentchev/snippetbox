package main

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/justinas/alice"
)

//Routes_1 ...
//Version 1:  func (app *Application) Routes() *http.ServeMux {
func (app *Application) Routes_1() http.Handler {

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

//Routes ...
func (app *Application) Routes() http.Handler {

	standardMiddleware := alice.New(app.RecoverPanic, app.LogRequest, SecureHeaders)
	// standardMiddleware = standardMiddleware.Append(SecureHeaders)

	dynamicMiddleware := alice.New(app.Session.Enable)

	mux := pat.New()

	mux.Get("/", dynamicMiddleware.ThenFunc(HomeWithClosure(app)))

	mux.Get("/snippet/create", dynamicMiddleware.ThenFunc(CreateSnippetFormWithClosure(app)))

	mux.Post("/snippet/create", dynamicMiddleware.ThenFunc(CreateSnippetWithClosure(app)))

	mux.Get("/snippet/:id", dynamicMiddleware.ThenFunc(ShowSnippetWithClosure(app)))

	// leave the static file without session cookie information(dynamicMiddleware)
	mux.Get("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./ui/static"))))

	return standardMiddleware.Then(mux)

}
