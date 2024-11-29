package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// employee represents data about a record employee.
type employee struct {
	Emp_id         string  `json:"emp_id"`
	Emp_firstname  string  `json:"emp_firstname"`
	Emp_lastname   string  `json:"emp_lastname"`
	Emp_department string  `json:"emp_department"`
	Emp_salary     float64 `json:"emp_salary"`
}

// employee slice to seed record employee data.
var employees = []employee{
	{Emp_id: "1", Emp_firstname: "Blue", Emp_lastname: "Coltrane", Emp_department: "sales", Emp_salary: 50000.00},
	{Emp_id: "2", Emp_firstname: "Jeru", Emp_lastname: "Mulligan", Emp_department: "marketing", Emp_salary: 55000.00},
	{Emp_id: "3", Emp_firstname: "Sarah", Emp_lastname: "Vaughan", Emp_department: "IT", Emp_salary: 80000.00},
}

func main() {
	router := gin.Default()
	router.GET("/employee", getEmployee)

	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

// getEmployee responds with the list of all employee as JSON.
func getEmployee(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, employees)
}
