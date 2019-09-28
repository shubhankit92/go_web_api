package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Candidate struct {
	Candidate_Id     string `json:"candidate_id"`
	Candidate_Name     string `json:"candidate_name"`
	Candidate_Phonenumber string `json:"candidate_phonenumber"`
	Status string `json:"status"`
}


func getCandidatesHandler(w http.ResponseWriter, r *http.Request) {
	candidates, err := store.GetCandidates()
	candidatesListBytes, err := json.Marshal(candidates)

	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(candidatesListBytes)
}

func createCandidateHandler(w http.ResponseWriter, r *http.Request) {
	candidate := Candidate{}
	decoder:= json.NewDecoder(r.Body)
    err := decoder.Decode(&candidate)
    if err != nil {
        panic(err)
    }
	err = store.CreateCandidates(&candidate)
	if err != nil {
		fmt.Println(err)
	}

	http.Redirect(w, r, "/hello", http.StatusFound)
}
func updateCandidateHandler(w http.ResponseWriter, r *http.Request) {
	candidate := Candidate{}
	decoder:= json.NewDecoder(r.Body)
    err := decoder.Decode(&candidate)
    if err != nil {
        panic(err)
    }
	err = store.UpdateCandidates(&candidate)
	if err != nil {
		fmt.Println(err)
	}

	http.Redirect(w, r, "/hello", http.StatusFound)
}
func deleteCandidateHandler(w http.ResponseWriter, r *http.Request) {
	candidate := Candidate{}
	decoder:= json.NewDecoder(r.Body)
    err := decoder.Decode(&candidate)
    if err != nil {
        panic(err)
    }
	err = store.DeleteCandidates(&candidate)
	if err != nil {
		fmt.Println(err)
	}

	http.Redirect(w, r, "/hello", http.StatusFound)
}
