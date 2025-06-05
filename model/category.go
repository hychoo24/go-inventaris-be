package model

type Category struct {
	ID          int    `json:"id" gorm:"primaryKey"`
	Name        string `json:"name" gorm:"type:varchar(100);not null;unique"`
	Description string `json:"description" gorm:"type:text"`
}
