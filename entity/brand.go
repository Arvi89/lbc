package entity

type Brand struct {
	ID   int
	Name string
	Cars []Car `gorm:"foreignKey:Brand"`
}
