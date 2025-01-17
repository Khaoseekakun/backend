package employee

import (
	"net/http"

	"backend/api/db"

	"strconv"

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
	Emp_id         int     `json:"emp_id"`
	Emp_firstname  string  `json:"emp_firstname" binding:"required"`
	Emp_lastname   string  `json:"emp_lastname" binding:"required"`
	Emp_department string  `json:"emp_department" binding:"required"`
	Emp_salary     float64 `json:"emp_salary" binding:"required"`
}

type EmployeeDelete struct {
	Emp_id int `json:"emp_id"`
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
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Invalid JSON"})
		return
	}

	if json.Emp_id == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Emp_id is required"})
		return
	}

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

func PutEmployee(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "PUT METHOD EMPLOYEE",
		"status":  "ok",
	})
}

func PutEmployeeDB(c *gin.Context) {
	empID := c.Param("id")

	var json EmployeeBody

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	if empID == "" {
		if json.Emp_id == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "emp_id is required add to path or body", "sample": "{emp_id:1} or employeedb/1"})
			return
		}

		empID = strconv.Itoa(json.Emp_id)
	} else {
		id, err := strconv.Atoi(empID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Emp_id"})
			return
		}
		json.Emp_id = id
	}

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

	tbl_employee := Tbl_employee(json)
	if err := db.Db.Where("emp_id = ?", empID).Model(&employee).Updates(tbl_employee).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Failed to update employee",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":   "ok",
		"message":  "Employee updated",
		"employee": tbl_employee,
	})
}

func DeleteEmployee(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "DELETE METHOD EMPLOYEE",
		"status":  "ok",
	})
}

func DeleteEmployeeDB(c *gin.Context) {
	empID := c.Param("id")
	var json EmployeeDelete
	if empID == "" {
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Emp_id is required add to path or body", "sample": "{Emp_id:1} or employeedb/1"})
			return
		}
		empID = strconv.Itoa(json.Emp_id)
	}

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

	db.Db.Where("emp_id = ?", empID).Delete(&employee)
	c.JSON(http.StatusOK, gin.H{
		"status":   "ok",
		"message":  "Employee deleted",
		"employee": employee,
	})
}
