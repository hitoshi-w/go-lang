package infrastructure

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Conn *gorm.DB

type DBConfig struct {
	Charset   string
	Password  string
	Collation string
	Host      string
	Name      string
	Port      string
	Username  string
}

// gorm
func InitializeDB(cf *DBConfig) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&collation=%s&parseTime=True&loc=Local", cf.Username, cf.Password, cf.Host, cf.Port, cf.Name, cf.Charset, cf.Collation)

	if _, err := gorm.Open(mysql.Open(dsn), &gorm.Config{}); err != nil {
    panic("failed to connect database: " + err.Error())
  }
}
