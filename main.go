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
	router.GET("/employee", EmployeeController.GetEmployeeDB)
	router.GET("/employee/:id", EmployeeController.GetEmployeeByID)
	router.POST("/employee", EmployeeController.PostEmployee)
	router.PUT("/employee", EmployeeController.PutEmployee)
	router.DELETE("/employee", EmployeeController.DeleteEmployee)

	router.Run(":8082")
}
