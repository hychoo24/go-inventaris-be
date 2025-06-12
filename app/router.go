package app

import (
	"tes1/controller/book"
	"tes1/controller/category"
	"tes1/controller/inventory"

	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) {

	//Router for books
	r.POST("/books", book.CreateBook)
	r.GET("/books", book.GetBooks)
	r.PUT("/books/:id", book.UpdateBook)
	r.DELETE("/books/:id", book.DeleteBook)
	r.GET("/books/:id", book.GetBookByID)

	//Router for inventory
	r.GET("/inventory", inventory.GetInventory)
	r.POST("/inventory", inventory.CreateInventoryItem)
	r.PUT("/inventory/:id", inventory.UpdateInventoryItem)
	r.DELETE("/inventory/:id", inventory.DeleteInventoryItem)
	r.GET("/inventory/:id", inventory.GetInventoryItemByID)

	// Additional routes can be added here
	r.GET("/category", category.GetCategory)

}
