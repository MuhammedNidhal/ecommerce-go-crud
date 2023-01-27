package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Product struct { //product db model
	gorm.Model
	Name  string
	Value int
}

type User struct { //user db model
	gorm.Model
	Name string
	Password string
}

func db() *gorm.DB { //sqlite3 db connection
	db, err := gorm.Open("sqlite3", "test.db")
	if err !=nil {
		panic("couldn't connect to db!")
	}
	return db
}

func getAllUsers(c *gin.Context) { //getting all users
	var users []User
	db().Find(&users)
	c.JSON(200, users)
}

func getUser(c *gin.Context) { //getting one user
	var user User
	if err := db().Where("id = ?",c.Param("id")).First(&user).Error; err != nil {
		c.JSON(404, gin.H{"error": "user not found !"})
		return
	}
	c.JSON(200, user)
}

func register(c *gin.Context) { //creating a new user
	var user User
	c.BindJSON(&user)
	db().Create(&user)
	c.JSON(201,gin.H{"success": "User Created Successfully!"})
}

func delUser(c *gin.Context) { //deleting a user
	var user User
	if err := db().Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(404, gin.H{"error": "User Not Found!"})
		return
	}
	db().Delete(&user)
	c.JSON(200, gin.H{"success": "User Deleted Successfully!"})
}

func updateUser(c *gin.Context) { //updating a user
	var user User
	if err := db().Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(404, gin.H{"error": "User Not Found!"})
		return
	}
	c.BindJSON(&user)
	db().Save(&user)
	c.JSON(200, gin.H{"success":"User Updated Successfully!"})
}

func getAllProducts(c *gin.Context) { //getting all products
	var products []Product
	db().Find(&products)
	c.JSON(200, products)
}

func getOne(c *gin.Context) { //getting a single product
	var product Product
	if err := db().Where("id = ?", c.Param("id")).First(&product).Error; err != nil {
		c.JSON(404, gin.H{"error": "Product not found!"})
		return
	}
	c.JSON(200, product)
}

func createNewProduct(c *gin.Context) { //creating a new product
	var product Product
	c.BindJSON(&product)
	db().Create(&product)
	c.JSON(201, gin.H{"success": "Product Created Successfully!"})
}

func updateProduct(c *gin.Context) { //updating a product
	var product Product
	if err := db().Where("id = ?", c.Param("id")).First(&product).Error; err != nil {
		c.JSON(404, gin.H{"error": "Product not found"})
		return
	}
	c.BindJSON(&product)
	db().Save(&product)
	c.JSON(200, gin.H{"success":"Product Updated Successfully!"})
}

func delProduct(c *gin.Context) { //deleting a product
	var product Product
	if err := db().Where("id = ?", c.Param("id")).First(&product).Error; err != nil {
		c.JSON(404, gin.H{"error": "Product not found!"})
		return
	}
	db().Delete(&product)
	c.JSON(200, gin.H{"success": "Product Deleted Successfully!"})
}

func main() {
	db().AutoMigrate(&Product{},&User{}) //migrating any changes in models to db

	r := gin.Default() //setting the default API router

	//--------------------------all endpoints used-------------------------------
	r.GET("/products", getAllProducts) //fetches all products
	r.GET("/products/:id", getOne) //fetches product by id
	r.POST("/products", createNewProduct) //creates a new products
	r.PUT("/products/:id", updateProduct) //updates an existing product by id
	r.DELETE("/products/:id", delProduct) //deletes an existing product by id
	r.GET("users", getAllUsers) //fetches all users
	r.GET("users/:id", getUser) //fetches user by id
	r.POST("users", register) //creates a new user
	r.PUT("users/:id", updateUser) //updates user information by id
	r.DELETE("users/:id", delUser) //deletes a user by id
	//---------------------------------------------------------------------------

	r.Run("localhost:8080") //runs the server on localhost = 127.0.0.1:8080
}
