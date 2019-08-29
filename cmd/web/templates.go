package main

import (
	"html/template"
	"path/filepath"
	"time"

	"github.com/plamenpentchev/snippetbox/pkg/forms"
	"github.com/plamenpentchev/snippetbox/pkg/models"
)

type templateData struct {
	Form        *forms.Form
	Flash       string
	Snippet     *models.Snippet
	Snippets    []*models.Snippet
	CurrentYear int
}

var templateFunctions = template.FuncMap{
	"humanDate": func(t time.Time) string {
		return t.Format("02 Jan 2006 at 15:04")
	},
}

//NewTemplateCache caches all th templates at aplicatopm start time
func NewTemplateCache(directory string) (map[string]*template.Template, error) {
	//Initialize a new map to act as a cache
	cache := make(map[string]*template.Template)

	pages, err := filepath.Glob(filepath.Join(directory, "*.page.tmpl"))
	if err != nil {
		return nil, err
	}
	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).Funcs(templateFunctions).ParseFiles(page)
		if err != nil {
			return nil, err
		}
		ts, err = ts.ParseGlob(filepath.Join(directory, "*.layout.tmpl"))
		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseGlob(filepath.Join(directory, "*.partial.tmpl"))
		if err != nil {
			return nil, err
		}
		cache[name] = ts
	}
	return cache, nil
}
