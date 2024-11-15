package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "GET METHOD",
		})
	})

	//---------------EMPLOYEE METHOD----------------
	r.GET("/employee", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "GET METHOD EMPLOYEE",
		})
	})

	r.POST("/employee", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "POST METHOD EMPLOYEE",
		})
	})

	r.PUT("/employee", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "PUT METHOD EMPLOYEE",
		})
	})

	r.DELETE("/employee", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "DELETE METHOD EMPLOYEE",
		})
	})

	//---------------CUSTOMER METHOD----------------
	r.GET("/customer", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "GET METHOD CUSTOMER",
		})
	})

	r.POST("/customer", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "POST METHOD CUSTOMER",
		})
	})

	r.PUT("/customer", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "PUT METHOD CUSTOMER",
		})
	})

	r.DELETE("/customer", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "DELETE METHOD CUSTOMER",
		})
	})

	//---------------CATEGORY METHOD----------------
	r.GET("/category", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "GET METHOD CATEGORY",
		})
	})

	r.POST("/category", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "POST METHOD CATEGORY",
		})
	})

	r.PUT("/category", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "PUT METHOD CATEGORY",
		})
	})

	r.DELETE("/category", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "DELETE METHOD CATEGORY",
		})
	})

	//---------------PRODUCT METHOD----------------
	r.GET("/product", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "GET METHOD PRODUCT",
		})
	})

	r.POST("/product", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "POST METHOD PRODUCT",
		})
	})

	r.PUT("/product", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "PUT METHOD PRODUCT",
		})
	})

	r.DELETE("/product", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "DELETE METHOD PRODUCT",
		})
	})

	//---------------ORDER METHOD----------------
	r.GET("/orders", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "GET METHOD ORDERS",
		})
	})

	r.POST("/orders", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "POST METHOD ORDERS",
		})
	})

	r.PUT("/orders", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "PUT METHOD ORDERS",
		})
	})

	r.DELETE("/orders", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "DELETE METHOD ORDERS",
		})
	})

	r.Run(":8001")
}
