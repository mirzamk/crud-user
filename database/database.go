package database

import (
    "log"
    "os"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

var DB *gorm.DB
var err error

func ConnectDB() {
    dsn := os.Getenv("MYSQL_DSN")
    if dsn == "" {
        log.Fatal("MYSQL_DSN environment variable not set")
    }

    DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to database: ", err)
    }
}
