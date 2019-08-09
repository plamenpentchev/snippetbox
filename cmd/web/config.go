package main

import (
	"html/template"
	"log"

	"github.com/plamenpentchev/snippetbox/pkg/models/mysql"
)

// Application ...
type Application struct {
	InfoLogger    *log.Logger
	ErrorLogger   *log.Logger
	SnippetModel  *mysql.SnippetModel
	TemplateCache map[string]*template.Template
}

//Env ...
type Env struct {
	InfoLog  *log.Logger
	ErrorLog *log.Logger
}
