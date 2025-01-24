package auth

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"backend/api/db"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

var hmacSampleSecret []byte

// Model
type Tbl_admin struct {
	Id        int `gorm:"primaryKey"`
	Firstname string
	Lastname  string
	Username  string
	Password  string
}

// Biding from Login JSON
type LoginBody struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// Login Function
func Login(c *gin.Context) {

	var json LoginBody
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//Check User Exists
	var adminExist Tbl_admin
	db.Db.Where("username = ?", json.Username).First(&adminExist)
	if adminExist.Id == 0 {
		c.JSON(http.StatusOK, gin.H{"status": "error", "message": "Admin Does Not Exists"})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(adminExist.Password), []byte(json.Password))
	if err == nil {

		hmacSampleSecret = []byte(os.Getenv("JWT_SECRET_KEY"))
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"Id":        adminExist.Id,
			"Firstname": adminExist.Firstname,
			"Lastname":  adminExist.Lastname,
			"Username":  adminExist.Username,
			"Password":  adminExist.Password,
			"exp":       time.Now().Add(time.Minute * 24).Unix(),
		})
		tokenString, err := token.SignedString(hmacSampleSecret)
		fmt.Println(tokenString, err)

		c.JSON(http.StatusOK, gin.H{"status": "ok", "message": "Login Success", "token": tokenString})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": "error", "message": "Login Failed"})
	}
}

func Auth(c *gin.Context) {
	hmacSampleSecret := []byte(os.Getenv("JWT_SECRET_KEY"))
	header := c.Request.Header.Get("Authorization")
	tokenString := strings.Replace(header, "Bearer ", "", 1)

	// Parse the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Check if the signing method used in the token is what we expect
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return hmacSampleSecret, nil
	})

	// Check if the token is valid
	if err == nil && token.Valid {
		log.Println("Token is Success")
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			c.Set("Id", claims["Id"])
			c.Set("Firstname", claims["Firstname"])
			c.Set("Lastname", claims["Lastname"])
			c.Set("Username", claims["Username"])
			c.Set("Password", claims["Password"])
			c.JSON(http.StatusOK, gin.H{"status": "ok", "message": "Token Success", "Id": claims["Id"], "Firstname": claims["Firstname"], "Lastname": claims["Lastname"], "Username": claims["Username"], "Password": claims["Password"]})
		}
	} else {
		log.Println("Token is invalid:", err)
		c.JSON(http.StatusOK, gin.H{"status": "error", "message": "Token Failed"})
	}

}
