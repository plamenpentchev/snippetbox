package main

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
	"strconv"
	"time"

	"github.com/plamenpentchev/snippetbox/pkg/forms"
	"github.com/plamenpentchev/snippetbox/pkg/models"
)

//HomeWithClosure ...
func HomeWithClosure(app *Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		app.InfoLogger.Println("HomeWithClosure ...")
		app.InfoLogger.Printf("Path :'%s'", r.URL.String())
		// if r.URL.String() != "/" {
		// 	app.NotFound(w)
		// 	return
		// }

		//panic("ooops something bad has happened") //Deliberate panic

		snippets, err := app.SnippetModel.Latest()
		if err != nil {
			app.ServerError(w, err)
			return
		}
		app.render(w, r, "home.page.tmpl", &templateData{
			Snippets:    snippets,
			CurrentYear: time.Time.Year(time.Now()),
		})
	}
}

//ShowSnippetWithClosure ...
func ShowSnippetWithClosure(app *Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		app.InfoLogger.Println("ShowSnippetWithClosure ...")
		id := r.URL.Query().Get(":id")
		validID, err := strconv.Atoi(id)
		if err == nil && validID >= 0 {
			app.InfoLogger.Printf("snippet id [%d]", validID)
			snip, err := app.SnippetModel.Get(validID)
			if err != nil {
				if err == models.ErrNoRecord {
					app.NotFound(w)
					return
				} else {
					app.ErrorLogger.Printf("No snippet to this ID: %d", validID)
				}

			} else {
				app.render(w, r, "show.page.tmpl", &templateData{
					Snippet:     snip,
					CurrentYear: time.Time.Year(time.Now()),
				})
			}

		} else {
			app.NotFound(w)
			http.Error(w, fmt.Sprintf("Incorrect ID [%s] should be numerical and positiv or 0", id), 405)
		}
	}
}

//CreateSnippetWithClosure ...
func CreateSnippetWithClosure(app *Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		app.InfoLogger.Println("CreateSnippetWitchClosure ...")
		app.InfoLogger.Printf("Method: %s ", r.Method)

		err := r.ParseForm()
		if err != nil {
			app.ClientError(w, http.StatusBadRequest)
			return
		}

		// title := r.PostForm.Get("title")
		// content := r.PostForm.Get("content")
		// expires := r.PostForm.Get("expires")

		//... validation issues
		// errors := make(map[string]string)
		// //... for title field
		// if strings.TrimSpace(title) == "" {
		// 	errors["title"] = "this field can not be empty"
		// } else {
		// 	if utf8.RuneCountInString(title) > 20 {
		// 		errors["title"] = "the character count in this field must not exceed 20"
		// 	}
		// }
		// //... for content field
		// if strings.TrimSpace(content) == "" {
		// 	errors["content"] = "this field can not be empty"
		// }
		// //... for expires field
		// expTmp := strings.TrimSpace(expires)
		// if expTmp == "" {
		// 	errors["expires"] = "this field can not be empty"
		// } else {
		// 	if expTmp != "365" && expTmp != "7" && expTmp != "1" {
		// 		errors["expires"] = fmt.Sprintf("inacceptable value in this field found - %s", expTmp)
		// 	}
		// }
		// if len(errors) > 0 {
		// 	fmt.Fprint(w, errors)
		// 	return
		// }
		// id, err := app.SnippetModel.Insert(title, content, expires)
		// if err != nil {
		// 	app.ServerError(w, err)
		// 	return
		// }

		//... if there were validation errors re-display the create template with the previously submitted values
		//... and errors

		form := forms.New(r.PostForm)
		form.Required("title", "content", "expires")
		form.MaxLength("title", 100)
		form.PermittedValues("expires", "365", "7", "1")

		if !form.Valid() {
			app.render(w, r, "create.page.tmpl", &templateData{
				Form: form,
			})
			return
		}
		//... end validation issues

		id, err := app.SnippetModel.Insert(form.Get("title"), form.Get("content"), form.Get("expires"))
		if err != nil {
			app.ServerError(w, err)
			return
		}
		app.InfoLogger.Printf("new snippet added [%d]", id)
		http.Redirect(w, r, fmt.Sprintf("/snippet/%d", id), http.StatusSeeOther)
	}
}

//CreateSnippetFormWithClosure ...
func CreateSnippetFormWithClosure(app *Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		app.render(w, r, "create.page.tmpl", &templateData{Form: &forms.Form{}})
	}
}

//DownLoadLogoHandler ....
func DownLoadLogoHandler(app *Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//ServeFIle doesnt sanitize the file path(warning dicrectory traversal attacks)
		//put file through filepath.Clean to sanitize path
		//"./ui/static/./../static/img/logo.png" == "./ui/static/img/logo.png"
		http.ServeFile(w, r, filepath.Clean("./ui/static/./../static/img/logo.png"))
	}
}

func (env *Env) home(w http.ResponseWriter, r *http.Request) {
	env.InfoLog.Printf("Path :'%s'", r.URL.String())
	if r.URL.String() != "/" {
		http.NotFound(w, r)
		return
	}

	templateFiles := []string{
		"./ui/html/home.page.tmpl",
		"./ui/html/base.layout.tmpl",
		"./ui/html/footer.partial.tmpl",
		// "./ui/html/sidebar.partial.tmpl",
	}

	// w.Write([]byte("Displays the home page"))
	ts, err := template.ParseFiles(templateFiles...)
	if err != nil {
		env.ErrorLog.Fatal(err.Error())
		http.Error(w, "Internal server error", 500)
	}
	err = ts.Execute(w, nil)
	if err != nil {
		env.ErrorLog.Fatal(err.Error())
		http.Error(w, "Internal server error", 500)
	}
}

func (env *Env) showSnippet(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")
	validID, err := strconv.Atoi(id)
	if err == nil && validID >= 0 {
		env.InfoLog.Printf("snippet id [%d]", validID)
		// w.Write([]byte("Display a specific snippet"))
		fmt.Fprintf(w, "Display a specific snippet-ID[%d]", validID)
	} else {
		http.Error(w, fmt.Sprintf("Incorrect ID [%s] should be numerical and positiv or 0", id), 405)
	}

}

func (env *Env) createSnippet(w http.ResponseWriter, r *http.Request) {
	env.InfoLog.Printf("Method: %s ", r.Method)
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		// w.WriteHeader(405)
		// w.Write([]byte(fmt.Sprintf("Method [%s] not allowed", r.Method)))
		http.Error(w, fmt.Sprintf("Method [%s] not allowed", r.Method), 405)
		env.ErrorLog.Printf("Method [%s] not allowed", r.Method)
		return
	}

	w.Header().Set("Cache-Control", "max-age=31536000")
	w.Header().Add("Cache-Control", "public")
	env.InfoLog.Printf("Cache Control: %s", w.Header().Get("Cache-Control"))
	w.Header()["Date"] = nil
	// w.Header()["Content-Length"] = []string{"0"}

	w.Write(([]byte("Create a new snippet")))
}
