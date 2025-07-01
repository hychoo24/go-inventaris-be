package homepage

import (
	"net/http"
	"tes1/model"
	"tes1/varglobal"

	"github.com/gin-gonic/gin"
)

func getStatsByCategory(table string) []model.CategoryStat {
	type Result struct {
		CategoryID   uint
		CategoryName string
		Total        int
	}

	var results []Result

	query := `
		SELECT c.id AS category_id, c.name AS category_name, COUNT(t.id) AS total
		FROM categories c
		LEFT JOIN ` + table + ` t ON t.category_id = c.id
		GROUP BY c.id, c.name
		ORDER BY total DESC
	`

	varglobal.DB.Raw(query).Scan(&results)

	stats := make([]model.CategoryStat, 0, len(results))
	for _, r := range results {
		stats = append(stats, model.CategoryStat{
			CategoryID:   r.CategoryID,
			CategoryName: r.CategoryName,
			Total:        r.Total,
		})
	}

	return stats
}

func GetHomepageStats(c *gin.Context) {
	var countBook, countInventory, countCategory int64

	varglobal.DB.Model(&model.Book{}).Count(&countBook)
	varglobal.DB.Model(&model.InventoryItem{}).Count(&countInventory)
	varglobal.DB.Model(&model.Category{}).Count(&countCategory)

	booksPerCategory := getStatsByCategory("books")
	inventoryPerCategory := getStatsByCategory("inventory_items")

	result := model.StatsAll{
		TtlBook:       countBook,
		TtlInventory:  countInventory,
		TtlCategory:   countCategory,
		DataBook:      booksPerCategory,
		DataInventory: inventoryPerCategory,
	}

	c.JSON(http.StatusOK, result)
}
