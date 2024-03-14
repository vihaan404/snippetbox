package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/vihaan404/snippetbox/internal/models"

	_ "github.com/go-sql-driver/mysql"
)

type applications struct {
	infoLog  *log.Logger
	errorLog *log.Logger
	snippets *models.SnippetModel
}

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	dsn := flag.String("dsn", "web:Rudy@123@/snippetbox?parseTime=true", "Mysql data source name")

	flag.Parse()

	// leveled logging

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	// Database
	db, err := openDB(*dsn)
	//
	if err != nil {
		errorLog.Fatal(err)
	}
	// closing before the main function exits
	defer db.Close()

	app := applications{
		infoLog:  infoLog,
		errorLog: errorLog,
		snippets: &models.SnippetModel{DB: db},
	}

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	// log.Printf("starting on %s", *addr)
	app.infoLog.Printf("starting on %s", *addr)

	err = srv.ListenAndServe()
	// log.Fatal(err)
	app.errorLog.Fatal(err)
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
