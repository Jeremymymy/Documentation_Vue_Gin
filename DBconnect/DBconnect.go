package dbconnect

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	//merge
)

var MySQLcon *gorm.DB
var err error

func init() {
	// refer htps://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	// dsn := "root:root@tcp(127.0.0.1:3306)/tsmcdocs?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := "root:root@tcp(db)/tsmcdocs?charset=utf8mb4&parseTime=True&loc=Local"
	MySQLcon, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
}
