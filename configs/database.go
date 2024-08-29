package configs

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var (
	DBConn *gorm.DB
)

func InitDB(config Config) {
	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s database=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", config.Database.DBHost, config.Database.Username, config.Database.Password, config.Database.DBName, config.Database.DBPort)
	DBConn, err = gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   config.Database.SchemaName + ".",
			SingularTable: false,
		}})
	if err != nil {
		panic("Failed to connect database")
	}

	fmt.Println("Database Connected")
}
