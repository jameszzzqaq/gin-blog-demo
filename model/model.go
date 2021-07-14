package model

import (
	"fmt"
	"log"

	"github.com/yu1er/gin-blog/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type Model struct {
	ID         int `gorm:"primary_key" json:"id"`
	CreatedOn  int `json:"created_on"`
	ModifiedOn int `json:"modified_on"`
}

var db *gorm.DB // db

func InitDB() {

	sec, err := config.Cfg.GetSection("database")
	if err != nil {
		log.Fatal(2, "Fail to get section 'database': %v", err)
	}

	// dbType := sec.Key("type").String()
	dbName := sec.Key("db_name").String()
	username := sec.Key("username").String()
	password := sec.Key("password").String()
	host := sec.Key("host").String()
	tablePrefix := sec.Key("table_prefix").String()

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		username, password, host, dbName)

	db, err = gorm.Open(mysql.New(mysql.Config{
		DSN:               dsn,
		DefaultStringSize: 171,
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   tablePrefix,
			SingularTable: true,
		},
	})
	if err != nil {
		log.Fatalf("Fail to open data-source-name: %v", err)
	}

	// set connect pool
	sqlDb, _ := db.DB()
	sqlDb.SetMaxIdleConns(10)
	sqlDb.SetMaxOpenConns(100)
}
