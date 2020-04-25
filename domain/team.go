package domain

type Team struct {
	ID           uint   `gorm:"primary_key"`
	Name         string `gorm:"size:255"`
	DepartmentId uint   `gorm:"column:DepartmentId"`
	Regions      []Region
}

func (Team) TableName() string {
	return "Teams"
}
