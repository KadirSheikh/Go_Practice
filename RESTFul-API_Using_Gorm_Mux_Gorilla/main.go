package main

import (
	"log"
	"net/http"

	"rest-api-go-mysql.com/controllers"
	"rest-api-go-mysql.com/database"
	"rest-api-go-mysql.com/entity"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql" //Required for MySQL dialect
)

func main() {
	initDB()
	log.Println("Starting the HTTP server on port 8090")

	router := mux.NewRouter().StrictSlash(true)
	initaliseHandlers(router)
	log.Fatal(http.ListenAndServe(":8090", router))
}

func initaliseHandlers(router *mux.Router) {
	router.HandleFunc("/create", controllers.CreateEmployee).Methods("POST")
	router.HandleFunc("/get", controllers.GetAllEmployee).Methods("GET")
	router.HandleFunc("/get/{id}", controllers.GetEmployeeByID).Methods("GET")
	router.HandleFunc("/update/{id}", controllers.UpdateEmployeeByID).Methods("PUT")
	router.HandleFunc("/delete/{id}", controllers.DeletEmployeeByID).Methods("DELETE")
}

func initDB() {
	config :=
		database.Config{
			ServerName: "localhost",
			User:       "root",
			Password:   "",
			DB:         "learning",
		}

	connectionString := database.GetConnectionString(config)
	err := database.Connect(connectionString)
	if err != nil {
		panic(err.Error())
	}
	database.Migrate(&entity.Employee{})
}
