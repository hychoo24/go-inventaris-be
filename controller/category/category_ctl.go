package category

import (
	"net/http"
	"strconv"
	"tes1/model"
	"tes1/varglobal"

	"github.com/gin-gonic/gin"
)

func GetCategory(c *gin.Context) {
	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "10")
	search := c.Query("search")

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
	query := varglobal.DB.Model(&model.Category{})

	// Tambahkan filter search
	if search != "" {
		query = query.Where("name LIKE ?", "%"+search+"%")
	}

	// Hitung total (sebelum pagination)
	var total int64
	if err := query.Count(&total).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menghitung total kategori"})
		return
	}

	// Ambil data dengan pagination
	var category []model.Category
	if err := query.Limit(limit).Offset(offset).Find(&category).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data kategori"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"page":    page,
		"limit":   limit,
		"total":   total,
		"results": category,
	})
}

// func untuk menambahkan kategori
func AddCategory(c *gin.Context) {
	var category model.Category
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data kategori tidak valid"})
		return
	}

	if err := varglobal.DB.Create(&category).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menambahkan kategori"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Kategori berhasil ditambahkan", "category": category})
}

// func untuk mengupdate kategori
func UpdateCategory(c *gin.Context) {
	id := c.Param("id")
	var category model.Category

	if err := varglobal.DB.First(&category, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Kategori tidak ditemukan"})
		return
	}

	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data kategori tidak valid"})
		return
	}

	if err := varglobal.DB.Save(&category).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengupdate kategori"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Kategori berhasil diupdate", "category": category})
}

// func untuk menghapus kategori
func DeleteCategory(c *gin.Context) {
	id := c.Param("id")
	var category model.Category

	if err := varglobal.DB.First(&category, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Kategori tidak ditemukan"})
		return
	}

	if err := varglobal.DB.Delete(&category).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menghapus kategori"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Kategori berhasil dihapus"})
}
