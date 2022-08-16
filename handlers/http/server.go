package http

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

var router *mux.Router

func ListenAndServe(port string, build embed.FS) {
	// create router
	router = mux.NewRouter()

	// build FS
	buildFS, err := fs.Sub(build, "views")
	if err != nil {
		log.Fatal(err)
	}

	// define routes
	router.HandleFunc("/", http.FileServer(http.FS(buildFS)).ServeHTTP).Methods(http.MethodGet)
	router.HandleFunc("/url", GetAll).Methods(http.MethodGet)
	router.HandleFunc("/url", ShortenURL).Methods(http.MethodPost)

	router.HandleFunc("/{id}", Resolve).Methods(http.MethodGet)
	router.HandleFunc("/qr/{id}", GetQR).Methods(http.MethodGet)

	// create logging and recovery middlewares
	loggedRouter := handlers.LoggingHandler(os.Stdout, router)
	recoveryRouter := handlers.RecoveryHandler()(loggedRouter)

	// run the http server
	log.Printf("url service running on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), recoveryRouter))
}