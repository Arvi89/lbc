package entity

type Ad struct {
	Id        int
	Title     string
	Content   string
	Car       *int
	CarStruct *Car `gorm:"foreignKey:Car"`
}
