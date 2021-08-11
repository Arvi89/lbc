package entity

type Category struct {
	ID   int
	Name string
	Ads  []Category `gorm:"foreignKey:Category"`
}
