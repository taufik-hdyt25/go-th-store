package models

import "time"

type Product struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	CategoryID uint      `json:"category_id"`
	Name       string    `json:"name"`
	Stok       int       `json:"stok"`
	Harga      float64   `json:"harga"`
	Category   Category  `gorm:"foreignKey:CategoryID" json:"category"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
