package main

import (
	"net/http"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

const (
	TITLE = "Greenleaf Cleaning"
)

func main(){
	r := gin.Default()

	r.LoadHTMLGlob("templates/*")

	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": TITLE,
		})
	})

	r.GET("/about", aboutHandler)

	r.GET("/contact", contactHandler)

	r.Use(static.Serve("/", static.LocalFile("./views", true)))

	// setup route group for the api
	api := r.Group("/api")
	{
		api.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H {
				"message": "pong",
			})
		})
	}

	r.Run(":5000")
}

type ContactInformation struct {
	Name			string	`json:"name"`
	Email			string	`json:"email"`
	PhoneNumber		string	`json:"phonenumber"`
}


func contactHandler(c *gin.Context) {
	var contact ContactInformation 

	if err := c.ShouldBindJSON(&contact); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Thank you, we will be in touch!",
		"contact": contact, 
	})
}


func aboutHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "about.html", gin.H{
		"page": "About",
		"title": TITLE,	
	})
}