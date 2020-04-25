package domain

import (
	"time"

	"github.com/jinzhu/gorm/dialects/postgres"
)

type Employee struct {
	ID              uint   `gorm:"primary_key"`
	Name            string `gorm:"size:255"`
	RoleId          uint   `gorm:"column:RoleId"`
	ManagerId       uint   `gorm:"column:ManagerId"`
	CardId          uint   `gorm:"column:CardId"`
	TeamId          uint   `gorm:"column:TeamId"`
	Active          bool
	CreatedAt       time.Time      `gorm:"column:createdAt"`
	UpdatedAt       time.Time      `gorm:"column:updatedAt"`
	HrDocumentation postgres.Jsonb `gorm:"column:hrDocumentation"`
}

func (Employee) TableName() string {
	return "Employees"
}
