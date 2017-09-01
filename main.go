package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
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
		panic("failed to connect database")
	}
	defer db.Close()

	db.AutoMigrate(
		&Todo{},
	)
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", handleFuncMain).Path("GET")
	router.HandleFunc("/all_todos", getAllTodos).Path("GET")
	router.HandleFunc("/show_todo", showTodo).Path("GET")
	router.HandleFunc("/add_todo", addTodo).Path("POST")
	router.HandleFunc("/delete_todo", deleteTodo).Path("POST")
	router.HandleFunc("is_checked_todo", complitedTodo).Path("POST")
	router.HandleFunc("edit_todo", editTodo).Path("POST")

	http.Handle("/", router)
	http.ListenAndServe(PORT_ADDR, nil)
}

func handleFuncMain(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w /*Main*/)
}

func getAllTodos(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w /*JSON witch have all todos*/)
}

func showTodo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w /*JSON witch have only one todo*/)
}

func addTodo(w http.ResponseWriter, r *http.Request) {
	// Add todo to db
	var err error
	if err != nil { // Error which can appeared when todo is added
		fmt.Fprintln(w, "error")
	} else {
		fmt.Fprintln(w, "ok")
	}
}

func deleteTodo(w http.ResponseWriter, r *http.Request) {
	// Delete todo from db
	var err error
	if err != nil { // Error which can appeared when todo is deleted
		fmt.Fprintln(w, "error")
	} else {
		fmt.Fprintln(w, "ok")
	}
}

func complitedTodo(w http.ResponseWriter, r *http.Request) {
	// Checked todo
	var err error
	if err != nil { // Error which can appeared when todo is complited
		fmt.Fprintln(w, "error")
	} else {
		fmt.Fprintln(w, "ok")
	}
}

func editTodo(w http.ResponseWriter, r *http.Request) {
	// Edit todo to db
	var err error
	if err != nil { // Error which can appeared when todo is edited
		fmt.Fprintln(w, "error")
	} else {
		fmt.Fprintln(w, "ok")
	}
}
