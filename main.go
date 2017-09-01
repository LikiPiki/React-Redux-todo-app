package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

const (
	PORT_ADDR    = ":3000"
	TODO_DB_ADDR = "todos"
)

var (
	db *gorm.DB
)

func init() {
	var err error
	db, err = gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database", err)
	}
	defer db.Close()

	db.AutoMigrate(
		&Todo{},
	)
}

func main() {
}
