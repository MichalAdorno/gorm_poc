package domain

type Role struct {
	ID        uint       `gorm:"primary_key"`
	Name      string     `gorm:"unique;not null;default:'EMPLOYEE'"`
	Employees []Employee `gorm:"foreignkey:RoleId"`
}

func (Role) TableName() string {
	return "Roles"
}
