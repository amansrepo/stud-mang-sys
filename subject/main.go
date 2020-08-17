package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type stud struct {
	num  byte   `json : "id"`
	name string `json : "name"`
	sub  string `json : "sub"`
	addr string `json : "addr"`
}

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "Aman123"
	dbname   = "studentinfo"
)

func main() {

	http.HandleFunc("/student_mgmt_info/student/1", func(rw http.ResponseWriter, r *http.Request) {

	}
	log.Fatal(http.ListenAndServe(":6060", nil))
}
