package main

import (
	"log"
	"net/http"

	_ "github.com/lib/pq"
	"github.com/stud-mang-sys/handlers/serverhandler"
)

func main() {
	http.HandleFunc("/student_mgmt_info/student/1", serverhandler.GetHandler)
	http.HandleFunc("/student_mgmt_info/student", serverhandler.PostHandler)
	http.HandleFunc("/student_mgmt_info/student/1/address", serverhandler.PutHandler)
	log.Fatal(http.ListenAndServe(":7070", nil))
}
