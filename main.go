package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)
func newRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/hello", handler).Methods("GET")
	r.HandleFunc("/records", getCandidatesHandler).Methods("GET")
	r.HandleFunc("/create", createCandidateHandler).Methods("POST")
	r.HandleFunc("/update", updateCandidateHandler).Methods("PATCH")
	r.HandleFunc("/delete", deleteCandidateHandler).Methods("DELETE")
	return r
}
const (
  host     = "localhost"
  port     = 5432
  user     = "postgres"
  password = "postgres"
  dbname   = "startetelelogic_interview"
)

func main() {
	fmt.Println("Starting server...")
	connString := fmt.Sprintf("host=%s port=%d user=%s "+
    "password=%s dbname=%s sslmode=disable",
    host, port, user, password, dbname)
	db, err := sql.Open("postgres", connString)

	if err != nil {
		panic(err)
	}
	err = db.Ping()

	if err != nil {
		panic(err)
	}

	InitStore(&dbStore{db: db})
	r := newRouter()
	fmt.Println("Serving on port 8080")
	http.ListenAndServe(":8080", r)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}
