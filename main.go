package main

import (
	EmployeeController "backend/api/controller/employee"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/employee", EmployeeController.GetEmployee)
	router.GET("/", EmployeeController.GetMain)
	router.GET("/employee/:id", EmployeeController.GetEmployeeByID)
	router.POST("/employee", EmployeeController.PostEmployee)
	router.PUT("/employee", EmployeeController.PutEmployee)
	router.DELETE("/employee", EmployeeController.DeleteEmployee)

	router.Run(":8082")
}
