package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

type applications struct {
	infoLog  *log.Logger
	errorLog *log.Logger
}

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")

	flag.Parse()

	// leveled logging

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := applications{
		infoLog:  infoLog,
		errorLog: errorLog,
	}

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux := http.NewServeMux()
	mux.HandleFunc("/", app.handlerHome)
	mux.HandleFunc("/snippet/view", app.snippetView)
	mux.HandleFunc("/snippet/create", app.snippetCreate)
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  mux,
	}

	// log.Printf("starting on %s", *addr)
	app.infoLog.Printf("starting on %s", *addr)

	err := srv.ListenAndServe()
	// log.Fatal(err)
	app.errorLog.Fatal(err)
}
