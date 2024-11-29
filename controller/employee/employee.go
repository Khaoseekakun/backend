package employee

import (
	"net/http"

	"backend/api/db"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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

func GetEmployeeByID(c *gin.Context) {
	empID := c.Param("id")

	var employee Tbl_employee

	if err := db.Db.Where("emp_id = ?", empID).First(&employee).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{
				"status":  "error",
				"message": "Employee not found",
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  "error",
				"message": "Failed to retrieve employee",
			})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":   "ok",
		"message":  "Employee found",
		"employee": employee,
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
