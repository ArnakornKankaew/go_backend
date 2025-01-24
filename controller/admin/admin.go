package admin

import (
	"net/http"

	"backend/api/db"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt" //import ไรบารี่ bcrypt
)

var hmacSampleSecret []byte //ประกาศตัวแปรเพื่อเข้ารหัส byte code

// Model
type Tbl_admin struct {
	Id        int `gorm:"primaryKey"`
	Firstname string
	Lastname  string
	Username  string
	Password  string
}

// Biding from Register JSON
type AdminBody struct {
	Id        int    `json:"id" binding:"required"`
	Firstname string `json:"firstname" binding:"required"`
	Lastname  string `json:"lastname" binding:"required"`
	Username  string `json:"username" binding:"required"`
	Password  string `json:"password" binding:"required"`
}

// GetAdmin
func GetAdmin(c *gin.Context) {
	var admins []Tbl_admin
	db.Db.Find(&admins)
	c.JSON(http.StatusOK, gin.H{"status": "ok", "message": "Admin Read Success", "Admins": admins})
}

// POST Admin to Database
func PostAdmin(c *gin.Context) {
	var json AdminBody
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	encryptedPassword, _ := bcrypt.GenerateFromPassword([]byte(json.Password), 10) //เอารหัสผ่านมาเข้ารหัส has
	tbl_admin := Tbl_admin{Id: json.Id, Firstname: json.Firstname, Lastname: json.Lastname, Username: json.Username, Password: string(encryptedPassword)}
	db.Db.Create(&tbl_admin)
	if tbl_admin.Firstname != "" {
		c.JSON(http.StatusOK, gin.H{"status": "ok", "message": "Admin Created", "tbl_success": tbl_admin})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": "error", "message": "Admin Failed", "tbl_fund": tbl_admin})
	}

}

