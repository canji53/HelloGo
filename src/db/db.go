package db

import (
	"fmt"
    "os"
	"github.com/joho/godotenv"
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
	"entity"
)

// global
var (
	db *gorm.DB
	err error
)

func Init() {

	// Read Dotenv
	ENV_PATH := ".env"
	err := godotenv.Load(fmt.Sprintf(ENV_PATH))
    if err != nil {
        panic(err)
    }

	// Connect MySQL
	DBMS     := "mysql"
	DB_USER  := os.Getenv("DB_USER")
	DB_PASS  := os.Getenv("DB_PASS")
	PROTOCOL := "tcp(" + os.Getenv("DB_HOST") + ":3306)"
	DB_NAME  := os.Getenv("DB_NAME")
	OPTION   := "?parseTime=true"
	CONNECT  := DB_USER + ":" + DB_PASS + "@" + PROTOCOL + "/" + DB_NAME + OPTION

    db, err = gorm.Open(DBMS, CONNECT)
    if err != nil {
        panic(err)
	}

	autoMigration()
}

func GetDB() *gorm.DB {
	return db
}

func Close() {
	if err := db.Close(); err != nil {
		panic(err)
	}
}

func autoMigration() {
	db.AutoMigrate(&entity.Memo{})
}