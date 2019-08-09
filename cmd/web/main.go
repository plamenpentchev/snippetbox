package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/plamenpentchev/snippetbox/pkg/models/mysql"
)

// const servAddr = ":4000"
var serverAddr *string
var dsn *string
var filePathOnError, logMicrosec *bool
var infoLog, errorLog *log.Logger
var logFlags int
var env *Env

func init() {
	serverAddr = flag.String("addr", ":4000", "HTTP network address")

	dsn = flag.String("dsn", "web:sn1pp3tb0x@tcp(192.168.99.101:3306)/snippetbox?parseTime=True", "mysql connection string")
	// dsn = flag.String("dsn", "root:root@tcp(192.168.99.100:3306)/snippetbox", "mysql connection string")
	filePathOnError = flag.Bool("logFilePathOnError", false, "log full file path on error")
	logMicrosec = flag.Bool("logMicroSec", false, "log with microseconds")
	flag.Parse()

	logFlags = log.Ldate | log.Ltime
	if *logMicrosec {
		logFlags |= log.Lmicroseconds
	}
	infoLog = log.New(os.Stdout, "INFO\t", logFlags)
	if *filePathOnError {
		logFlags |= log.Llongfile
	} else {
		logFlags |= log.Lshortfile
	}
	errorLog = log.New(os.Stderr, "ERROR\t", logFlags)
	env = &Env{
		InfoLog:  infoLog,
		ErrorLog: errorLog,
	}
}

func main() {

	app := &Application{
		InfoLogger:  infoLog,
		ErrorLogger: errorLog,
	}

	app.InfoLogger.Printf("connecting ... [%s]", *dsn)
	db, err := OpenDB(*dsn)
	if err != nil {
		app.ErrorLogger.Fatal(err)
	} else {
		app.InfoLogger.Println("DB connection pool established")
	}

	defer db.Close()

	templateCache, err := NewTemplateCache("./ui/html/")
	if err != nil {
		app.ErrorLogger.Fatal(err)
	}

	// inject the data base model over the application
	app.SnippetModel = &mysql.SnippetModel{DB: db}
	//set the cached templates
	app.TemplateCache = templateCache

	// create new server to log its issues in our error log file
	srv := http.Server{
		Addr:     *serverAddr,
		ErrorLog: errorLog,
		Handler:  app.Routes(),
	}

	infoLog.Printf("Listening on  '%s'", *serverAddr)
	// err := http.ListenAndServe(*serverAddr, mux)
	err = srv.ListenAndServe()
	if err != nil {
		errorLog.Fatal(err)
	}
}

//OpenDB ...
func OpenDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
