package domain

type Car struct {
	ID         uint   `gorm:"primary_key"`
	Name       string `gorm:"size:255"`
	RegistryNr `gorm:"column:registryNr;size:255;not null"`
}

func (Car) TableName() string {
	return "Cars"
}
