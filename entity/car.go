package entity

type Car struct {
	ID          int
	Brand       int
	BrandStruct Brand `gorm:"foreignKey:Brand"`
	Model       string
}
