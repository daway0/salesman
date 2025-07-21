package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"salesman/db"
	"salesman/handlers"

	// "html/template"
	// "path/filepath"

	"github.com/gin-gonic/gin"
)

func CustomHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Hello, World!"})
}

func main() {
	DB_HOST := os.Getenv("DB_HOST")
	DB_PORT := os.Getenv("DB_PORT")
	DB_USER := os.Getenv("DB_USER")
	DB_PASSWORD := os.Getenv("DB_PASSWORD")
	DB_NAME := os.Getenv("DB_NAME")
	DB_SSLMODE := os.Getenv("DB_SSLMODE")

	dataSourceName := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", DB_USER, DB_PASSWORD, DB_HOST, DB_PORT, DB_NAME, DB_SSLMODE)
	dbConn, err := db.InitDB(dataSourceName)
	if err != nil {
		log.Fatal(err)
	}
	defer dbConn.Close()

	r := gin.Default()
	r.Static("/vendors", "./vendors")
	r.Static("/build", "./build")

	// t, err := template.ParseFiles(
	// 	filepath.Join("templates", "layouts", "base.html"),
	// 	filepath.Join("templates", "partials", "menu_profile.html"),
	// 	filepath.Join("templates", "partials", "menu_sidebar.html"),
	// 	filepath.Join("templates", "partials", "top_navigation.html"),
	// 	filepath.Join("templates", "user_form.html"),
	// )

	// t, err := template.ParseFiles(
	// 	filepath.Join("templates", "base.html"),
	// 	filepath.Join("templates", "child1.html"),
	// 	filepath.Join("templates", "child2.html"),
	// )
	// if err != nil {
	// 	panic(err)
	// }
	// r.SetHTMLTemplate(t)

	r.Use(func(c *gin.Context) {
		c.Set("db", dbConn)
		c.Next()
	})

	handlers.SetupRoutes(r, dbConn)
	// r.GET("/xx", func(c *gin.Context) {
	// 	c.HTML(http.StatusOK, "child1.html", gin.H{})
	// })

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
