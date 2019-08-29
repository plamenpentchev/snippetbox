package main

import (
	"fmt"
	"net/http"
	"runtime/debug"
	"time"
)

func (app *Application) addDefaultData(td *templateData, r *http.Request) *templateData {

	if nil == td {
		td = &templateData{}
	}
	//... set current year
	td.CurrentYear = time.Time.Year(time.Now())
	//... retireve information from the current user session
	if nil != app.Session && app.Session.Exists(r, "flash") {
		td.Flash = app.Session.PopString(r, "flash")
	}

	return td
}

func (app *Application) render(w http.ResponseWriter, r *http.Request, name string, td *templateData) {
	ts, ok := app.TemplateCache[name]
	if !ok {
		app.ServerError(w, fmt.Errorf("The template %s does not exist", name))
		return
	}

	//execute the template set, passing in any data
	err := ts.Execute(w, app.addDefaultData(td, r))
	// err := ts.Execute(w, td)
	if err != nil {
		app.ServerError(w, err)
		return
	}
}

//ServerError ...
func (app *Application) ServerError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n\n%s", err.Error(), debug.Stack())
	app.ErrorLogger.Output(2, trace)
	http.Error(w, fmt.Sprintf("%s\n%s", http.StatusText(http.StatusInternalServerError), trace), http.StatusInternalServerError)
}

//ClientError ...
func (app *Application) ClientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

//NotFound ...
func (app *Application) NotFound(w http.ResponseWriter) {
	app.ClientError(w, http.StatusNotFound)
}
