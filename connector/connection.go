package connector

import "fmt"

func (c *DbConfig) GetConnectionString() string {
	return fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable",
		c.Host, c.Port, c.DbName, c.User, c.Password)
}
