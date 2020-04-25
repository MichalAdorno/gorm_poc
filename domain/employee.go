package domain

import (
	"time"

	"github.com/jinzhu/gorm/dialects/postgres"
)

type Employee struct {
	ID     uint   `gorm:"primary_key"`
	Name   string `gorm:"size:255"`
	Role   Role   `gorm:"foreignkey:RoleId"`
	RoleId uint   `gorm:"column:RoleId"`
	// ManagerId       uint `gorm:"column:ManagerId"`
	Car             Car  `gorm:"foreignkey:CarId"`
	CarId           uint `gorm:"column:CarId"`
	Team            Team `gorm:"foreignkey:TeamId"`
	TeamId          uint `gorm:"column:TeamId"`
	Active          bool
	CreatedAt       time.Time      `gorm:"column:createdAt"`
	UpdatedAt       time.Time      `gorm:"column:updatedAt"`
	HrDocumentation postgres.Jsonb `gorm:"column:hrDocumentation"`
}

func (Employee) TableName() string {
	return "Employees"
}
