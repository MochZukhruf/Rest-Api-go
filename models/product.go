package models

type Product struct {
	Id          int64  `gorm:"primaryKey" json:"id"`
	Nama_Produk string `gorm:"type:varchar(300)" json:"nama_produk"`
	Deskripsi   string `gorm:"type:text" json:"deskripsi"`
	// Stock     int64   `gorm:"type:smallint(6)" json:"stock"`
	// Category  Category `gorm:"foreignKey:CategoryID" json:"category"`
	// CategoryID int64   `gorm:"not null" json:"category_id"`
}

// type Category struct {
// 	Id   int64  `gorm:"primaryKey" json:"id"`
// 	Name string `gorm:"type:varchar(100)" json:"name"`
// }
