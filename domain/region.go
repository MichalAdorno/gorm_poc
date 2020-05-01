package domain

import (
	"time"

	"github.com/jinzhu/gorm/dialects/postgres"
)

type Region struct {
	ID          uint   `gorm:"primary_key"`
	Name        string `gorm:"size:255"`
	Description postgres.Jsonb
	Teams       []Team    `gorm:"many2many:TeamRegion";association_jointable_foreignkey:RegionId;jointable_foreignkey:TeamId`
	CreatedAt   time.Time `gorm:"column:createdAt"`
}

func (Region) TableName() string {
	return "Regions"
}
