package main

import (
	"database/sql"
	"fmt"
)
type Store interface {
	CreateCandidates(candidate *Candidate) error
	UpdateCandidates(candidate *Candidate) error
	DeleteCandidates(candidate *Candidate) error
	GetCandidates() ([]*Candidate, error)
}
type dbStore struct {
	db *sql.DB
}

func (store *dbStore) CreateCandidates(candidate *Candidate) error {
	_, err := store.db.Query("INSERT INTO candidates(candidate_name, candidate_phonenumber, status) VALUES ($1, $2, $3)",
		candidate.Candidate_Name, candidate.Candidate_Phonenumber, candidate.Status);
	return err
}

func (store *dbStore) UpdateCandidates(candidate *Candidate) error {
    fmt.Println("Update.............")
	stmt, err := store.db.Prepare("update candidates set candidate_name=$1, candidate_phonenumber=$2, status=$3 where candidate_id=$4")
    _, err = stmt.Exec(candidate.Candidate_Name, candidate.Candidate_Phonenumber, candidate.Status, candidate.Candidate_Id)
    return err
}
func (store *dbStore) DeleteCandidates(candidate *Candidate) error {
    fmt.Println("Delete.............")
	stmt, err := store.db.Prepare("delete from candidates where candidate_id=$1")
    _, err = stmt.Exec(candidate.Candidate_Id)
    return err
}

func (store *dbStore) GetCandidates() ([] *Candidate, error) {
	rows, err := store.db.Query("SELECT candidate_name, candidate_phonenumber, status from candidates")
	if err != nil {
		return nil, err
	}
	fmt.Println(rows);

	defer rows.Close()
	list := [] *Candidate{}
	for rows.Next() {
		object := &Candidate{}
		if err := rows.Scan(&object.Candidate_Name,&object.Candidate_Phonenumber,&object.Status); err != nil {
			return nil, err
		}
		list = append(list, object)
	}
	return list, nil
}
var store Store
func InitStore(s Store) {
	store = s
}
