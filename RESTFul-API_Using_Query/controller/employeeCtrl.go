package controller

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"employee-rest-api/connection"
	"employee-rest-api/entity"
)

//Get All Employee
func GetAllEmployees(w http.ResponseWriter, r *http.Request) {
	db := connection.DbConn()

	if r.Method == "GET" {
		selDB, err := db.Query("SELECT * FROM Employee ORDER BY id DESC")
		if err != nil {
			panic(err.Error())
		}
		emp := entity.Employee{}
		res := []entity.Employee{}
		for selDB.Next() {
			var id int
			var name, city string
			err = selDB.Scan(&id, &name, &city)
			if err != nil {
				panic(err.Error())
			}
			emp.Id = id
			emp.Name = name
			emp.City = city
			res = append(res, emp)
		}
		// tmpl.ExecuteTemplate(w, "Index", res)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(res)
	}
	defer db.Close()
}

func GetEmployeeById(w http.ResponseWriter, r *http.Request) {
	db := connection.DbConn()
	if r.Method == "GET" {

		id, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil || id < 1 {
			http.NotFound(w, r)
			return
		}
		selDB, err := db.Query("SELECT * FROM Employee WHERE id=?", id)
		if err != nil {
			panic(err.Error())
		}
		emp := entity.Employee{}
		for selDB.Next() {
			var id int
			var name, city string
			err = selDB.Scan(&id, &name, &city)
			if err != nil {
				panic(err.Error())
			}
			emp.Id = id
			emp.Name = name
			emp.City = city
		}
		// tmpl.ExecuteTemplate(w, "Show", emp)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(emp)

	}
	defer db.Close()
}

// //Create Employee

func AddEmployee(w http.ResponseWriter, r *http.Request) {
	db := connection.DbConn()
	requestBody, _ := ioutil.ReadAll(r.Body)
	var employee entity.Employee
	json.Unmarshal(requestBody, &employee)
	if r.Method == "POST" {

		insForm, err := db.Prepare("INSERT INTO Employee(name, city) VALUES(?,?)")
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(employee.Name, employee.City)
		log.Println("INSERT: Name: " + employee.Name + " | City: " + employee.City)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(employee)
	defer db.Close()
	// http.Redirect(w, r, "/", 301)
}

func UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	db := connection.DbConn()
	requestBody, _ := ioutil.ReadAll(r.Body)
	var employee entity.Employee
	json.Unmarshal(requestBody, &employee)
	if r.Method == "PUT" {

		id, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil || id < 1 {
			http.NotFound(w, r)
			return
		}

		insForm, err := db.Prepare("UPDATE Employee SET name=?, city=? WHERE id=?")
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(employee.Name, employee.City, id)
		log.Println("UPDATE: Name: " + employee.Name + " | City: " + employee.City)

	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(employee)
	defer db.Close()

}

func DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	db := connection.DbConn()
	if r.Method == "DELETE" {
		id, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil || id < 1 {
			http.NotFound(w, r)
			return
		}
		delForm, err := db.Prepare("DELETE FROM Employee WHERE id=?")
		if err != nil {
			panic(err.Error())
		}
		delForm.Exec(id)
	}

	defer db.Close()

}
