package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// defined type Employee
type Employee struct {
	Id   string `json:"Id"`
	Name string `json:"name"`
}

// declared a global array
var Employees []Employee

// to get root page (/)
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Root page")
	fmt.Println("Root page")
}

// to return all employee records (/emplpoyees)
func returnAllEmployees(w http.ResponseWriter, r *http.Request) {
	//fmt.Println("Endpoint Hit: returnAllEmployees")
	json.NewEncoder(w).Encode(Employees)

}

// to return single employee records (/emplpoyees/{empid})
func returnSingleEmp(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	fmt.Fprintf(w, "Key: "+key)

	for _, emp := range Employees {
		if emp.Id == key {
			json.NewEncoder(w).Encode(emp)
		}
	}
}

//post
func createNewEmp(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hi from create")
	reqBody, _ := ioutil.ReadAll(r.Body)
	//fmt.Fprintf(w, "%+v", string(reqBody))
	var emp Employee
	json.Unmarshal(reqBody, &emp)
	Employees = append(Employees, emp)
	json.NewEncoder(w).Encode(emp)
}

//delete
func deleteEmp(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hi from delete")
	vars := mux.Vars(r)
	id := vars["id"]
	for index, emp := range Employees {
		if emp.Id == id {

			Employees = append(Employees[:index], Employees[index+1:]...)
		}
	}
}

//update PUT
func updateEmp(w http.ResponseWriter, r *http.Request) {
	fmt.Println("from update")
	EmployeeID := mux.Vars(r)["id"]
	var emp Employee

	reqBody, _ := ioutil.ReadAll(r.Body)

	json.Unmarshal(reqBody, &emp)

	for i, looper := range Employees {
		if looper.Id == EmployeeID {
			looper.Id = emp.Id
			looper.Name = emp.Name
			Employees = append(Employees[:i], looper)
			json.NewEncoder(w).Encode(looper)
		}
	}
}

// handlers
func handleRequests() {
	// http.HandleFunc("/employees", returnAllEmployees)
	// http.HandleFunc("/", homePage)
	// log.Fatal(http.ListenAndServe(":8080", nil))
	//using  mux for better api usage
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/employees", returnAllEmployees)
	myRouter.HandleFunc("/employee", createNewEmp).Methods("POST")
	myRouter.HandleFunc("/employee/{id}", deleteEmp).Methods("DELETE")
	myRouter.HandleFunc("/employee/{id}", updateEmp).Methods("PUT")
	myRouter.HandleFunc("/employee/{id}", returnSingleEmp)
	log.Fatal(http.ListenAndServe(":8089", myRouter))
}

func main() {
	fmt.Println("Hi")

	Employees = []Employee{
		Employee{Id: "1032", Name: "Vishnu"},
		Employee{Id: "1033", Name: "Abhi"},
	}
	handleRequests()
}
