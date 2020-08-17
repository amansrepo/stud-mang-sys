package main

import (
	"bytes"
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

func getReqHandler() {
	id := 1
	idJSON, _ := json.Marshal(id)
	req, err := http.NewRequest(http.MethodGet, "http://localhost:7070/student_mgmt_info/student/1", bytes.NewBuffer(idJSON))
	if err != nil {
		log.Fatal(err)
	}

	//req.Header.Set("Content-Type", "application/json; charset=utf-8")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	var dat = stud{}
	data, _ := ioutil.ReadAll(resp.Body)

	if err := json.Unmarshal(data, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)
	resp.Body.Close()
}

func postReqHandler() {

	studJson, err := json.Marshal(stud{Num: 2, Name: "Shweta Verma", Sub: "Physics", Addr: "New Delhi"})
	res, err := http.NewRequest(http.MethodPost, "http://localhost:7070/student_mgmt_info/student", bytes.NewBuffer(studJson))
	if err != nil {
		log.Fatal(err)
	}

	client := &http.Client{}
	resp, err := client.Do(res)
	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(resp.Body)

	fmt.Println("Response Data: ", string(body))
	res.Body.Close()
}

func putReqHandler() {
	newAddress := "Pune"
	addrJSON, err := json.Marshal(newAddress)
	req, err := http.NewRequest(http.MethodPut, "http://localhost:7070/student_mgmt_info/student/1/address", bytes.NewBuffer(addrJSON))
	if err != nil {
		panic(err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(resp.Body)

	fmt.Println("Response Data: ", string(body))
	resp.Body.Close()
}
func main() {

	//getReqHandler()

	//postReqHandler()

	putReqHandler()

}
