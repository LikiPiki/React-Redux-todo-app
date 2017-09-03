package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

const (
	PORT_ADDR     = ":3000"
	TODO_DB_ADDR  = "todos"
	STATIC_FOLDER = "/static/"
)

var (
	db *gorm.DB
)

func init() {
	var err error
	db, err = gorm.Open("sqlite3", "data.sqlite3")
	db.LogMode(true)
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(
		&Todo{},
	)
}

func main() {
	defer db.Close()

	router := mux.NewRouter().StrictSlash(true)
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./dist/"))))
	router.HandleFunc("/", handleFuncMain)
	router.HandleFunc("/get_all", getAllTodos)
	router.HandleFunc("/add_new", addTodo)
	router.HandleFunc("/delete_todo", deleteTodo)
	router.HandleFunc("is_checked_todo", complitedTodo)
	router.HandleFunc("edit_todo", editTodo)
	// static handle

	fmt.Println("Starting server on port", PORT_ADDR)
	err := http.ListenAndServe(PORT_ADDR, router)
	if err != nil {
		log.Println("Cant start the server")
	}
}
