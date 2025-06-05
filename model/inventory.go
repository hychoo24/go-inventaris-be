package model

type InventoryItem struct {
	ID         int      `json:"id"`
	Name       string   `json:"name"`
	Quantity   int      `json:"quantity"`
	CategoryID int      `json:"category_id"`           // Jika menggunakan relasi ke model Category
	Category   Category `gorm:"foreignKey:CategoryID"` // Relasi ke model Category
}
