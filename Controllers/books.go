package Controllers

import (
	Connectdatabase "GORM-GO/ConnectDatabase"
	models "GORM-GO/model"

	"net/http"

	"github.com/gin-gonic/gin"
)

// var books models.Books

func JustPrint(g *gin.Context) {
	res := gin.H{
		"result": "true",
	}
	g.JSON(http.StatusOK, res)
}

func CreateBooks(c *gin.Context) {
	// validate input
	var input models.CreateBooksInput

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// create books
	book := models.Books{Title: input.Title, Author: input.Author}
	Connectdatabase.DB.Create(&book)

	c.JSON(http.StatusOK, gin.H{
		"data": book,
	})
}

func FindBook(c *gin.Context) {

	var books models.Books

	if err := Connectdatabase.DB.Where("id = ?", c.Param("id")).First(&books).Error; err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": books,
	})
}

func FindBooks(c *gin.Context) {

	var Books []models.Books

	// models.DB.Find(&Books)
	Connectdatabase.DB.Find(&Books)

	res := gin.H{
		"data": Books,
	}

	c.JSON(http.StatusOK, res)
}
