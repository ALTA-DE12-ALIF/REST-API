package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DaftarUser = []User{}

type User struct {
	Name   string `json:"name" form:"name"`
	HP     string `json:"hp" form:"hp"`
	Alamat string `json:"alamat" form:"alamat"`
	Email  string `json:"email" form:"email"`
}

func connectDB() *gorm.DB {
	dsn := "root:@tcp(localhost:3306)/logistic?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return db
}

// Login
func GetLogin(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		name := c.Param("name")

		var resQry User

		if err := db.First(&resQry, "name = ?", name).Error; err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": "cannot select data",
			})
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "succes Login user",
			"data":    resQry,
		})
	}
}

// Register
func PostRegister(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var regis User
		if err := c.Bind(&regis); err != nil {
			log.Error(err)
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "cannot select data",
			})
		}
		if err := db.Create(&regis).Error; err != nil {
			log.Error(err)
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": "cannot select data",
			})
		}
		return c.JSON(http.StatusCreated, map[string]interface{}{
			"massage": "succes Register new user",
			"data":    regis,
		})
	}
}

// Get All Data Vendor
func AllVendor(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var vendor []User
		if err := db.Find(&vendor).Error; err != nil {
			log.Error(err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": "error on database",
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "succes get all data",
			"data":    vendor,
		})
	}
}

// Data Vendor
func DateVendor(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var Dvendor []User
		if err := db.Find(&Dvendor).Error; err != nil {
			log.Error(err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": "error on database",
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "succes get data Vendor",
			"data":    Dvendor,
		})
	}
}

// Tambah Data Vendor (POST)
func CreateVendor(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var CreateVendor User
		if err := c.Bind(&CreateVendor); err != nil {
			log.Error(err)
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "cannot select data",
			})
		}
		if err := db.Create(&CreateVendor).Error; err != nil {
			log.Error(err)
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": "cannot select data",
			})
		}
		return c.JSON(http.StatusCreated, map[string]interface{}{
			"massage": "succes Create Data Vendor",
			"data":    CreateVendor,
		})
	}
}

func main() {
	e := echo.New()
	db := connectDB()
	db.AutoMigrate(&User{})
	e.Use(middleware.Logger())

	o := e.Group("/orm")
	o.GET("/login/:name", GetLogin(db))
	o.GET("/vendor", AllVendor(db))
	o.POST("/register", PostRegister(db))
	o.GET("/dataVendor", DateVendor(db))
	o.POST("/createVendor", CreateVendor(db))
	e.Start(":8000")
}
