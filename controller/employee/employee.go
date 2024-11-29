package employee

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type employee struct {
	Emp_id         string  `json:"emp_id"`
	Emp_firstname  string  `json:"emp_firstname"`
	Emp_lastname   string  `json:"emp_lastname"`
	Emp_department string  `json:"emp_department"`
	Emp_salary     float64 `json:"emp_salary"`
}

var employees = []employee{
	{Emp_id: "1", Emp_firstname: "Blue", Emp_lastname: "Coltrane", Emp_department: "sales", Emp_salary: 50000.00},
	{Emp_id: "2", Emp_firstname: "Jeru", Emp_lastname: "Mulligan", Emp_department: "marketing", Emp_salary: 55000.00},
	{Emp_id: "3", Emp_firstname: "Sarah", Emp_lastname: "Vaughan", Emp_department: "IT", Emp_salary: 80000.00},
}

func GetMain(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"message": "Please use path /employee",
	})
}

func GetEmployee(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{
		"data":   employees,
		"status": "ok",
	})
}

func GetEmployeeByID(c *gin.Context) {
	id := c.Param("id")

	for _, emp := range employees {
		if emp.Emp_id == id {
			c.IndentedJSON(http.StatusOK, gin.H{
				"data":   emp,
				"status": "ok",
			})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{
		"status":  "not_found",
		"message": "Employee not found",
	})
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
