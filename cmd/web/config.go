package main

import (
	"html/template"
	"log"

	"github.com/plamenpentchev/snippetbox/pkg/models/mysql"
	"github.com/golangcollege/sessions"
)

// Application ...
type Application struct {
	InfoLogger    *log.Logger
	ErrorLogger   *log.Logger
	Session *sessions.Session
	SnippetModel  *mysql.SnippetModel
	TemplateCache map[string]*template.Template
}

//Env ...
type Env struct {
	InfoLog  *log.Logger
	ErrorLog *log.Logger
}
