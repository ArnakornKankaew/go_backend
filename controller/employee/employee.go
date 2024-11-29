package employee

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GET
func GetEmployee(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Employee GET ALL Method!",
	})
}

// GET BY ID
func GetEmployeeByID(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"message": id,
	})
}

// POST
func PostEmployee(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Employee POST Method!",
	})
}

// PUT
func PutEmployee(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Employee PUT Method!",
	})
}

// DELETE
func DeleteEmployee(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Employee DELETE Method!",
	})
}
