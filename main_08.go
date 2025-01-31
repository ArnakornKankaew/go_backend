package main

import (
	AdminController "backend/api/controller/admin"
	AuthController "backend/api/controller/auth"
	EmployeeController "backend/api/controller/employee"
	"backend/api/middleware"

	"backend/api/db"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	//เรียกใช้ไฟล์ที่อยู่ในห้อง middleware
)

func main() {
	//Get .env
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	//get InitDB fuction
	db.InitDB()

	router := gin.Default()

	authorized := router.Group("/api", middleware.JwtAuthen())
	authorized.GET("/employeedb", EmployeeController.GetEmployeeDB)
	authorized.GET("/employee", EmployeeController.GetEmployee)
	authorized.GET("/employee/:id", EmployeeController.GetEmployeeByID)
	authorized.POST("/employee", EmployeeController.PostEmployee)
	authorized.POST("/employeedb", EmployeeController.PostEmployeeDB)
	authorized.PUT("/employee", EmployeeController.PutEmployee)
	authorized.PUT("/employeedb", EmployeeController.PutEmployeeDB)
	authorized.DELETE("/employee", EmployeeController.DeleteEmployee)
	authorized.DELETE("/employeedb/:id", EmployeeController.DeleteEmployeeDB)

	router.POST("/register", AdminController.PostAdmin)
	router.POST("/login", AuthController.Login)

	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
