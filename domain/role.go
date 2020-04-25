package domain

type Role struct {
	ID   uint   `gorm:"primary_key"`
	Name string `gorm:"unique;not null;default:'EMPLOYEE'"`
}

func (Role) TableName() string {
	return "Roles"
}
