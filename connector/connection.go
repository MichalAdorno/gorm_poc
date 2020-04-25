package connector

import "fmt"

func (c *DbConfig) GetConnectionString() string {
	return fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s", c.Host, c.Port, c.User, c.DbName, c.Password)
}
