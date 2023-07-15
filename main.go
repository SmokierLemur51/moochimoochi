package main

import (
	"net/http"
	
	"./routes"

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
	
	// routes
	r.GET("/about", rotues.aboutHandler)
	r.GET("/contact", routes.contactHandler)
	r.GET("/testing", routes.testHandler)


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


