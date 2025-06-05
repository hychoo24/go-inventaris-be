package book

import (
	"net/http"
	"strconv"
	"tes1/model"
	"tes1/varglobal"

	"github.com/gin-gonic/gin"
)

func CreateBook(c *gin.Context) {
	var newBook model.Book
	if err := c.ShouldBindJSON(&newBook); err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid request payload",
		})
		return
	}

	if err := varglobal.DB.Create(&newBook).Error; err != nil {
		c.JSON(500, gin.H{
			"error": "Failed to create book",
		})
		return
	}

	c.JSON(201, gin.H{
		"message": "Book created successfully",
		"book":    newBook,
	})
}

func UpdateBook(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID harus berupa angka"})
		return
	}

	var updatedBook model.Book
	if err := c.ShouldBindJSON(&updatedBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var existingBook model.Book
	if err := varglobal.DB.First(&existingBook, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Buku tidak ditemukan"})
		return
	}

	// Perbarui field buku yang ditemukan dengan data baru
	existingBook.Title = updatedBook.Title
	existingBook.Author = updatedBook.Author
	existingBook.Year = updatedBook.Year
	existingBook.CategoryID = updatedBook.CategoryID
	existingBook.PublishedYear = updatedBook.PublishedYear
	// Tambahkan field lain sesuai model kamu

	if err := varglobal.DB.Save(&existingBook).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal memperbarui buku"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Buku berhasil diperbarui", "book": existingBook})

}

func GetBooks(c *gin.Context) {
	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "10")
	search := c.Query("search")
	category := c.Query("category")

	page, _ := strconv.Atoi(pageStr)
	limit, _ := strconv.Atoi(limitStr)
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}
	offset := (page - 1) * limit

	// Mulai query
	query := varglobal.DB.Preload("Category").Model(&model.Book{})

	// Tambahkan filter search
	if search != "" {
		query = query.Where("title LIKE ?", "%"+search+"%")
	}

	// Tambahkan filter kategori
	if category != "" {
		query = query.Where("category_id LIKE ?", "%"+category+"%")
	}

	// Hitung total (sebelum pagination)
	var total int64
	if err := query.Count(&total).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menghitung total buku"})
		return
	}

	// Ambil data dengan pagination
	var books []model.Book
	if err := query.Limit(limit).Offset(offset).Find(&books).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data buku"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"page":    page,
		"limit":   limit,
		"total":   total,
		"results": books,
	})
}

func DeleteBook(c *gin.Context) {
	id := c.Param("id")
	var book model.Book

	if err := varglobal.DB.First(&book, id).Error; err != nil {
		c.JSON(404, gin.H{
			"error": "Book not found",
		})
		return
	}

	if err := varglobal.DB.Delete(&book).Error; err != nil {
		c.JSON(500, gin.H{
			"error": "Failed to delete book",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Book deleted successfully",
	})

}

func GetBookByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID tidak valid"})
		return
	}

	var book model.Book
	if err := varglobal.DB.First(&book, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Buku tidak ditemukan"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"book": book})

}
