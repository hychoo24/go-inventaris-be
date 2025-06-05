package model

type Book struct {
	ID            int      `json:"id"`
	Title         string   `json:"title"`
	Author        string   `json:"author"`
	Year          int      `json:"year"`
	CategoryID    int      `json:"category_id"`           // Jika menggunakan relasi ke model Category
	Category      Category `gorm:"foreignKey:CategoryID"` // Relasi ke model Category
	PublishedYear int      `json:"published_year"`
}
