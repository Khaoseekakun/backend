package main

import (
	EmployeeController "backend/api/controller/employee"
	"fmt"

	"backend/api/db"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	//get InitDB fuction
	db.InitDB()

	router := gin.Default()
	router.GET("/", EmployeeController.GetMain)
	router.GET("/employeedb", EmployeeController.GetEmployeeDB)
	router.GET("/employeedb/:id", EmployeeController.GetEmployeeByID)
	router.POST("/employee", EmployeeController.PostEmployee)
	router.POST("/employeedb", EmployeeController.PostEmployeeDB) //POST
	router.PUT("/employee", EmployeeController.PutEmployee)
	router.PUT("/employeedb/:id", EmployeeController.PutEmployeeDB)
	router.PUT("/employeedb", EmployeeController.PutEmployeeDB)
	router.DELETE("/employee", EmployeeController.DeleteEmployee)
	router.DELETE("/employeedb/:id", EmployeeController.DeleteEmployeeDB)
	router.DELETE("/employeedb", EmployeeController.DeleteEmployeeDB)

	router.Run(":8082")
}
