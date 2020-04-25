package domain

type Region struct {
	ID    uint   `gorm:"primary_key"`
	Name  string `gorm:"size:255"`
	Teams []Team `gorm:"many2many:TeamRegion";association_jointable_foreignkey:RegionId;jointable_foreignkey:TeamId`
}

func (Region) TableName() string {
	return "Regions"
}
