package domain

import (
	"time"

	"github.com/jinzhu/gorm/dialects/postgres"
)

type Department struct {
	ID              uint           `gorm:"primary_key"`
	Name            string         `gorm:"size:255"`
	ManagerId       uint           `gorm:"column:ManagerId"`
	CreatedAt       time.Time      `gorm:"column:createdAt"`
	UpdatedAt       time.Time      `gorm:"column:updatedAt"`
	HrDocumentation postgres.Jsonb `gorm:"column:hrDocumentation"`
	Teams           []Team         `gorm:"foreignkey:DepartmentId"`
}

func (Department) TableName() string {
	return "Departments"
}
