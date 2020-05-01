package domain

type Team struct {
	ID           uint        `gorm:"primary_key"`
	Name         string      `gorm:"size:255"`
	Department   *Department `gorm:"foreignkey:DepartmentId"`
	DepartmentId uint        `gorm:"column:DepartmentId"`
	Regions      []Region    `gorm:"many2many:TeamRegion";association_jointable_foreignkey:TeamId;jointable_foreignkey:RegionId`
	Employees    []Employee  `gorm:"foreignkey:TeamId"`
}

func (Team) TableName() string {
	return "Teams"
}
