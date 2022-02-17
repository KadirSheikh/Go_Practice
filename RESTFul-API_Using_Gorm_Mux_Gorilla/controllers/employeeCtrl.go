package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"rest-api-go-mysql.com/database"
	"rest-api-go-mysql.com/entity"

	"github.com/gorilla/mux"
)

//GetAllEmployee get all employee data
func GetAllEmployee(w http.ResponseWriter, r *http.Request) {

	//Creating array of employees struct
	var employees []entity.Employee

	//Get All Employee data using find
	database.Connector.Find(&employees)

	//Set Header
	w.Header().Set("Content-Type", "application/json")

	//Check status
	w.WriteHeader(http.StatusOK)

	//Get data json format
	json.NewEncoder(w).Encode(employees)
}

//GetEmployeeByID returns employee with specific ID
func GetEmployeeByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	var employee entity.Employee
	database.Connector.First(&employee, key)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(employee)
}

//CreateEmployee creates employee
func CreateEmployee(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	var employee entity.Employee
	json.Unmarshal(requestBody, &employee)

	database.Connector.Create(employee)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(employee)
}

//UpdateEmployeeByID updates employee with respective ID
func UpdateEmployeeByID(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	key := vars["id"]

	// requestBody, _ := ioutil.ReadAll(r.Body)
	// var employee entity.Employee
	// json.Unmarshal(requestBody, &employee)

	id, _ := strconv.ParseInt(key, 10, 64)
	// database.Connector.Where("id = ?", id).Save(&employee)

	var employee entity.Employee
	json.NewDecoder(r.Body).Decode(&employee)
	database.Connector.Where("id = ?", id).Save(&employee)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(employee)
}

//DeletEmployeeByID delete's employee with specific ID
func DeletEmployeeByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	var employee entity.Employee
	id, _ := strconv.ParseInt(key, 10, 64)
	database.Connector.Where("id = ?", id).Delete(&employee)
	w.WriteHeader(http.StatusNoContent)
}
