package employee

import (
	"net/http"

	"backend/api/db"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Tbl_employee struct {
	Emp_id         int
	Emp_firstname  string
	Emp_lastname   string
	Emp_department string
	Emp_salary     float64
}

type EmployeeBody struct {
	Emp_id         int     `json:"emp_id" binding:"required"`
	Emp_firstname  string  `json:"emp_firstname" binding:"required"`
	Emp_lastname   string  `json:"emp_lastname" binding:"required"`
	Emp_department string  `json:"emp_department" binding:"required"`
	Emp_salary     float64 `json:"emp_salary" binding:"required"`
}

func GetMain(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"message": "Please use path /employeedb",
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

func PostEmployeeDB(c *gin.Context) {
	var json EmployeeBody
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//check body request

	tbl_employee := Tbl_employee(json)
	db.Db.Create(&tbl_employee)
	if tbl_employee.Emp_firstname != "" {
		c.JSON(http.StatusOK, gin.H{"status": "ok", "message": "Employee Created", "tbl_fund": tbl_employee})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": "error", "message": "User Failed", "tbl_fund": tbl_employee})
	}

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

func PutEmployeeDB(c *gin.Context) {
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

	var json EmployeeBody
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tbl_employee := Tbl_employee(json)
	db.Db.Where("emp_id = ?", empID).Model(&employee).Updates(tbl_employee)
	c.JSON(http.StatusOK, gin.H{
		"status":   "ok",
		"message":  "Employee updated",
		"employee": employee,
	})
}

// DELETE Method
func DeleteEmployee(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "DELETE METHOD EMPLOYEE",
		"status":  "ok",
	})
}

func DeleteEmployeeDB(c *gin.Context) {
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
	} else {
		//delete and check if record exist
		db.Db.Where("emp_id = ?", empID).Delete(&employee)
		c.JSON(http.StatusOK, gin.H{
			"status":   "ok",
			"message":  "Employee deleted",
			"employee": employee,
		})

	}

}
