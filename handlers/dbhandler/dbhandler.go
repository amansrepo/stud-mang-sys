package dbhandler

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "Aman123"
	dbname   = "studentinfo"
)

func DbConnection() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
}

func SqlInsert(id int, name string, subject string, address string){

	sqlStatement := `INSERT INTO details (id, name, subject, address)
VALUES (1, 'Aman Singh', 'Mathematics', 'Bangalore')`
	_, err := db.Exec(sqlStatement)
	if err != nil {
		panic(err)
	}
}

func SqlSelect(id int){
	sqlStatement := `SELECT id, name, subject, address FROM details WHERE id=$1;`
	row := db.QueryRow(sqlStatement, id)
	
	switch err := row.Scan(&id, &name,$subject, &address); err {
	case sql.ErrNoRows:
	  fmt.Println("No rows were returned!")
	case nil:
	  fmt.Println(id, email)
	default:
	  panic(err)
	}
}

func SqlUpdate(id int){
	sqlStatement := `UPDATE users SET address = Mumbai WHERE id = $1;`
	row := db.QueryRow(sqlStatement, id)
	
	switch err := row.Scan(&id, &name,$subject, &address); err {
	case sql.ErrNoRows:
	  fmt.Println("No rows were returned!")
	case nil:
	  fmt.Println(id, email)
	default:
	  panic(err)
	}
}


