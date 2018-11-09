package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.DisableConsoleColor()
	// router := gin.Default()
	router := gin.New()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// parameters in path
	router.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})

	router.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		message := name + " is " + action
		c.String(http.StatusOK, message)
	})

	// querystring parameters
	router.GET("/welcome", func(c *gin.Context) {
		firstname := c.DefaultQuery("firstname", "guest")
		lastname := c.Query("lastname")
		c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
	})

	// multipart/urlencoded form
	router.POST("/form_post", func(c *gin.Context) {
		message := c.PostForm("message")
		nick := c.DefaultPostForm("nick", "anonymous")
		c.JSON(200, gin.H{
			"status":  "posted",
			"message": message,
			"nick":    nick,
		})
	})

	// map as querystring or postform parameters
	router.POST("/post", func(c *gin.Context) {
		ids := c.QueryMap("ids")
		names := c.PostFormMap("names")
		fmt.Printf("ids: %v; names: %v", ids, names)
	})

	// upload file
	router.POST("/upload", func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
			log.Fatalln(err)
		}
		log.Println(file.Filename)

		c.SaveUploadedFile(file, "testfile.zip")
		c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
	})

	// Middlewares
	authorized := router.Group("/api", gin.Logger())
	{
		authorized.POST("/read", readEndpoint)
	}

	router.Run() // listen and serve on 0.0.0.0:8080
}

func readEndpoint(c *gin.Context) {
	c.String(http.StatusOK, "Hello read")
}
