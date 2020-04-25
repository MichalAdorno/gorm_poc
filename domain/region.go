package domain

type Region struct {
	ID    uint   `gorm:"primary_key"`
	Name  string `gorm:"size:255"`
	Teams []Team
}

func (Region) TableName() string {
	return "Regions"
}
