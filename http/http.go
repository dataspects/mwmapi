package http

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// Serve responses to all incoming API requests
func Serve(p int, ao []string) error {
	r := mux.NewRouter().StrictSlash(true)

	r.HandleFunc("/indexing/facetmanager", api.Info).Methods("POST")

	log.Println("Web Service listening on: http://localhost:" + strconv.Itoa(p))
	c := cors.New(cors.Options{
		AllowedOrigins:   ao,
		AllowedHeaders:   []string{"Content-Type", "Content-Length", "Accept-Encoding", "Authorization"},
		AllowedMethods:   []string{"GET", "PATCH", "POST", "PUT", "OPTIONS", "DELETE"},
		AllowCredentials: true,
	})

	handler := c.Handler(r)
	err := http.ListenAndServe(":"+strconv.Itoa(p), handler)
	return err
}
