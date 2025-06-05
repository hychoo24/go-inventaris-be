package inventory

import (
	"net/http"
	"strconv"
	"tes1/model"
	"tes1/varglobal"

	"github.com/gin-gonic/gin"
)

func GetInventory(c *gin.Context) {
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
	query := varglobal.DB.Preload("Category").Model(&model.InventoryItem{})

	// Tambahkan filter search
	if search != "" {
		query = query.Where("name LIKE ?", "%"+search+"%")
	}

	// Tambahkan filter kategori
	if category != "" {
		query = query.Where("category_id LIKE ?", "%"+category+"%")
	}

	// Hitung total (sebelum pagination)
	var total int64
	if err := query.Count(&total).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menghitung total inventory item"})
		return
	}

	// Ambil data dengan pagination
	var inventory []model.InventoryItem
	if err := query.Limit(limit).Offset(offset).Find(&inventory).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data inventory item"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"page":    page,
		"limit":   limit,
		"total":   total,
		"results": inventory,
	})
}

func CreateInventoryItem(c *gin.Context) {
	var newItem model.InventoryItem
	if err := c.ShouldBindJSON(&newItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := varglobal.DB.Create(&newItem).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create inventory item",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message":   "Inventory item created successfully",
		"inventory": newItem,
	})

}

func UpdateInventoryItem(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID tidak valid"})
		return
	}

	var updatedItem model.InventoryItem
	if err := c.ShouldBindJSON(&updatedItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var existingItem model.InventoryItem
	if err := varglobal.DB.First(&existingItem, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item inventaris tidak ditemukan"})
		return
	}

	// Perbarui field yang diizinkan
	existingItem.Name = updatedItem.Name
	existingItem.CategoryID = updatedItem.CategoryID
	existingItem.Quantity = updatedItem.Quantity
	// Tambahkan field lain sesuai dengan model kamu

	if err := varglobal.DB.Save(&existingItem).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal memperbarui item inventaris"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Item inventaris berhasil diperbarui", "item": existingItem})

}

func DeleteInventoryItem(c *gin.Context) {
	id := c.Param("id")
	var inventory model.InventoryItem

	if err := varglobal.DB.First(&inventory, id).Error; err != nil {
		c.JSON(404, gin.H{
			"error": "Inventory item not found",
		})
		return
	}

	if err := varglobal.DB.Delete(&inventory).Error; err != nil {
		c.JSON(500, gin.H{
			"error": "Failed to delete Inventory item",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Inventory tem deleted successfully",
	})
}

func GetInventoryItemByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID tidak valid"})
		return
	}

	var item model.InventoryItem
	if err := varglobal.DB.First(&item, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item inventaris tidak ditemukan"})
		return
	}

	c.JSON(http.StatusOK, item)

}
