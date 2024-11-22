package employee

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GET Method
func GetEmployee(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "GET METHOD EMPLOYEE",
	})
}

// GET Method by ID
func GetEmployeeByID(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"message": "GET METHOD EMPLOYEE BY ID",
		"id":      id,
	})
}

// POST Method
func PostEmployee(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "POST METHOD EMPLOYEE",
	})
}

// PUT Method
func PutEmployee(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "PUT METHOD EMPLOYEE",
	})
}

// DELETE Method
func DeleteEmployee(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "DELETE METHOD EMPLOYEE",
	})
}
