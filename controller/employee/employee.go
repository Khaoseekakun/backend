package employee

import (
	"net/http"

	"backend/api/db"

	"github.com/gin-gonic/gin"
)

type Tbl_employee struct {
	Emp_id        int
	Emp_firstname string
	Emp_lastname  string
}

func GetMain(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"message": "Please use path /employee",
	})
}

func GetEmployeeDB(c *gin.Context) {
	var employee []Tbl_employee
	db.Db.Find(&employee)
	c.JSON(http.StatusOK, gin.H{"status": "ok", "message": "Employee Read Success", "employees": employee})
}

func PostEmployee(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "POST METHOD EMPLOYEE",
		"status":  "ok",
	})
}

// PUT Method
func PutEmployee(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "PUT METHOD EMPLOYEE",
		"status":  "ok",
	})
}

// DELETE Method
func DeleteEmployee(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "DELETE METHOD EMPLOYEE",
		"status":  "ok",
	})
}
