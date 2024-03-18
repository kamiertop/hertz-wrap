package storage

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// InitGorm init gorm with mysql driver.
func InitGorm() error {
	_, err := gorm.Open(mysql.Open(""), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{},
	})
	if err != nil {
		return fmt.Errorf("open gorm db error: %v", err)
	}

	return nil
}
