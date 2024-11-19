package domain

type Product struct {
	ID          uint    `gorm:"primaryKey;autoIncrement"`
	Name        string  `gorm:"size:255"`
	Description string  `gorm:"size:120"`
	Price       float64 `gorm:"size:255"`
	CreatedAt   int64   `gorm:"autoCreateTime"` // save in epoch time
	UpdatedAt   int64   `gorm:"autoUpdateTime"`

	CategoryID uint
	Category   Category `gorm:"foreignKey:CategoryID"`
}
