package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB

// Person :
type Person struct {
	gorm.Model
	FirstName string
	LastName  string
}

func main() {
	var err error
	// open db connection
	db, err = gorm.Open("postgres", "host=localhost port=5432 user=nilsonmolina dbname=demo-people password=nils0n sslmode=disable")
	if err != nil {
		panic(err)
	}
	db.Close()

	db.AutoMigrate(&Person{})
	// setup router and handlers
	router := gin.Default()

	router.GET("/people", getPeople)
	router.GET("/people/:id", getPerson)
	router.POST("/people", createPerson)
	router.PUT("/people/:id", updatePerson)
	router.DELETE("/people/:id", deletePerson)
	router.Run(":8080")
}

func getPeople(c *gin.Context) {
	var people []Person
	if err := db.Find(&people).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
		return
	}
	c.JSON(200, people)
}

func getPerson(c *gin.Context) {
	fmt.Fprint(c.Writer, "GET Person")

}

func createPerson(c *gin.Context) {
	var person Person
	c.BindJSON(&person)

	db.Create(&person)
	c.JSON(200, person)
}

func updatePerson(c *gin.Context) {

}

func deletePerson(c *gin.Context) {

}
