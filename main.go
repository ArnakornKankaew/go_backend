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

	authorized := router.Group("/api", middleware.JwtAuthen())      //ทำการจัดกลุ่ม path ที่ต้องการล๊อค api
	authorized.GET("/employeedb", EmployeeController.GetEmployeeDB) //ล๊อค api โดยต้องแนบ token ก่อนถึงใช้งานได้

	//Employee API Method
	authorized.GET("/employee", EmployeeController.GetEmployee)         //GET
	authorized.GET("/employee/:id", EmployeeController.GetEmployeeByID) //GET BY ID

	authorized.POST("/employee", EmployeeController.PostEmployee)     //POST
	authorized.POST("/employeedb", EmployeeController.PostEmployeeDB) //POST TO DB
	router.POST("/register", AdminController.PostAdmin)               //POST TO DB

	authorized.PUT("/employee", EmployeeController.PutEmployee)     //PUT
	authorized.PUT("/employeedb", EmployeeController.PutEmployeeDB) //PUT DB
	router.POST("/login", AuthController.Login)                     //POST LOGIN

	authorized.DELETE("/employee", EmployeeController.DeleteEmployee)         //DELETE
	authorized.DELETE("/employeedb/:id", EmployeeController.DeleteEmployeeDB) //DELETE DB

	//Customer API Method

	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
