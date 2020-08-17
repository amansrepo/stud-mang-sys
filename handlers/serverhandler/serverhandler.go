package serverhandler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type stud struct {
	Num  byte   `json : "id"`
	Name string `json : "name"`
	Sub  string `json : "sub"`
	Addr string `json : "addr"`
}

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "Aman123"
	dbname   = "studentinfo"
)

// GET Handler

func GetHandler(rw http.ResponseWriter, r *http.Request) {
	var id int
	if r.Method == "GET" {
		jsndata, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatal("Error reading the body", err)
		}

		err = json.Unmarshal(jsndata, &id)
		if err != nil {
			log.Fatal("Decoding error: ", err)
		}

		log.Printf("Received Data: %v\n", id)

		// Retrieving from Database

		psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
			"password=%s dbname=%s sslmode=disable",
			host, port, user, password, dbname)

		db, err := sql.Open("postgres", psqlInfo)
		if err != nil {
			panic(err)
		}

		err = db.Ping()
		if err != nil {
			panic(err)
		}

		fmt.Println("Successfully connected!")

		var user stud

		sqlStatement := `SELECT id, name, subject, address FROM details WHERE id=$1;`
		row := db.QueryRow(sqlStatement, id)

		err = row.Scan(&user.Num, &user.Name, &user.Sub, &user.Addr)
		db.Close()

		studJSON, err := json.Marshal(user)
		if err != nil {
			fmt.Fprintf(rw, "Error: %s", err)
		}
		fmt.Println(user)
		rw.Write(studJSON)
	} else {
		rw.Write([]byte("Expecting GET Requests only"))
	}
}

//POST Handler

func PostHandler(rw http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {

		studData := stud{}
		jsndata, err := ioutil.ReadAll(r.Body)
		err = json.Unmarshal(jsndata, &studData)
		if err != nil {
			log.Fatal("Decoding error: ", err)
		}

		log.Printf("Received Data: %v\n", studData)

		// Posting Data into Database

		psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
			"password=%s dbname=%s sslmode=disable",
			host, port, user, password, dbname)

		db, err := sql.Open("postgres", psqlInfo)
		if err != nil {
			panic(err)
		}

		err = db.Ping()
		if err != nil {
			panic(err)
		}

		fmt.Println("Successfully connected!")

		sqlStatement := `INSERT INTO details (id, name, subject, address)
		   		VALUES ($1, $2, $3, $4)`
		_, err = db.Exec(sqlStatement, studData.Num, studData.Name, studData.Sub, studData.Addr)
		if err != nil {
			panic(err)
		}
		db.Close()
		rw.Write([]byte("Data written into Database Successfully"))

	} else {
		rw.Write([]byte("Expecting POST Requests only"))
	}
}

// PUT Handler

func PutHandler(rw http.ResponseWriter, r *http.Request) {
	if r.Method == "PUT" {

		var address string
		jsndata, err := ioutil.ReadAll(r.Body)
		err = json.Unmarshal(jsndata, &address)
		if err != nil {
			log.Fatal("Decoding error: ", err)
		}

		log.Printf("Received Data: %v\n", address)

		psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
			"password=%s dbname=%s sslmode=disable",
			host, port, user, password, dbname)

		db, err := sql.Open("postgres", psqlInfo)
		if err != nil {
			panic(err)
		}

		err = db.Ping()
		if err != nil {
			panic(err)
		}

		fmt.Println("Successfully connected!")

		sqlStatement := `UPDATE details SET address = $2 WHERE id = $1;`
		_, err = db.Exec(sqlStatement, 1, address)
		if err != nil {
			panic(err)
		}
		db.Close()

		rw.Write([]byte("Data updated Successfully"))
	} else {
		rw.Write([]byte("Expecting PUT Requests only"))
	}
}

/*=========================================================================*/

type Student struct{}

type Address struct{}

type Subject struct{}

func StudServ() *Student {
	return &Student{}
}

func AddressServ() *Address {
	return &Address{}
}

func SubjectServ() *Subject {
	return &Subject{}
}

func (st *Student) ServeHTTP(rw http.ResponseWriter, r *http.Request) {

}

func (ad *Address) ServeHTTP(rw http.ResponseWriter, r *http.Request) {

}

func (su *Subject) ServeHTTP(rw http.ResponseWriter, r *http.Request) {

}
