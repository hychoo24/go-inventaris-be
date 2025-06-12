package category

import (
	"net/http"
	"tes1/model"

	"github.com/gin-gonic/gin"
)

// controller get category
var categories = []model.Category{}

func GetCategory(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"data": categories,
	})
}
