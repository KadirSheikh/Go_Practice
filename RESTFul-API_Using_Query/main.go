package main

import (
	"log"
	"net/http"

	"employee-rest-api/controller"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	log.Println("Server started on: http://localhost:8070")
	http.HandleFunc("/get", controller.GetAllEmployees)
	http.HandleFunc("/getById", controller.GetEmployeeById)
	http.HandleFunc("/add", controller.AddEmployee)
	http.HandleFunc("/update", controller.UpdateEmployee)
	http.HandleFunc("/delete", controller.DeleteEmployee)

	http.ListenAndServe(":8070", nil)
}
