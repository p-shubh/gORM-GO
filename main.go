package main

import (
	connectdatabase "GORM-GO/ConnectDatabase"
	"GORM-GO/Controllers"

	"github.com/gin-gonic/gin"
	// "github.com/rs/cors/wrapper/gin"
)

func main() {

	gin.SetMode(gin.ReleaseMode)

	router := gin.New()

	// r.GET("/", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{"data": "hello world"})
	// })

	setupRoutes(router)

	router.SetTrustedProxies([]string{"192.168.1.2"})

	go connectdatabase.Connectdatabase()

	router.Run(":8080")

	// go fmt.Println("finally the test is working")

}

func setupRoutes(g *gin.Engine) {

	g.GET("/", Controllers.JustPrint)
	g.GET("/books", Controllers.FindBooks)
	g.POST("/entry", Controllers.CreateBooks)
	g.GET("/book/:id", Controllers.FindBook)
}
