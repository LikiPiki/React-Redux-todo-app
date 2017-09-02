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
	PORT_ADDR    = ":3000"
	TODO_DB_ADDR = "todos"
)

var (
	db *gorm.DB
)

func init() {
	var err error
	db, err = gorm.Open("sqlite3", "test.db")
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
	router.HandleFunc("/", handleFuncMain)
	router.HandleFunc("/all_todos", getAllTodos)
	router.HandleFunc("/show_todo", showTodo)
	router.HandleFunc("/add_todo", addTodo)
	router.HandleFunc("/delete_todo", deleteTodo)
	router.HandleFunc("is_checked_todo", complitedTodo)
	router.HandleFunc("edit_todo", editTodo)
	fmt.Println("Starting server on port", PORT_ADDR)
	err := http.ListenAndServe(PORT_ADDR, router)
	if err != nil {
		log.Println("Cant start the server")
	}
}

func handleFuncMain(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `
        / - main
        /all_todos
        /show_todo
        /add_todo
        /delete_todo
        /is_checked_todo
        /edit_todo`)
}

func getAllTodos(w http.ResponseWriter, r *http.Request) {
	var todos []Todo
	db.Find(&todos)
	fmt.Fprintln(w)
}

func showTodo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Something")
}

func addTodo(w http.ResponseWriter, r *http.Request) {
	todo, err := getTodoFromRequest(r)
	db.Create(&Todo{Name: todo.Name, Checked: todo.Checked})
	if err != nil {
		fmt.Fprintln(w, "error")
	} else {
		fmt.Fprintln(w, "ok")
	}
}

func deleteTodo(w http.ResponseWriter, r *http.Request) {
	todo, err := getTodoFromRequest(r)
	db.Delete(&todo)
	if err != nil {
		fmt.Fprintln(w, "error")
	} else {
		fmt.Fprintln(w, "ok")
	}
}

func complitedTodo(w http.ResponseWriter, r *http.Request) {
	todo, err := getTodoFromRequest(r)
	db.Model(&todo).Update("Checked", todo.Checked)
	if err != nil {
		fmt.Fprintln(w, "error")
	} else {
		fmt.Fprintln(w, "ok")
	}
}

func editTodo(w http.ResponseWriter, r *http.Request) {
	todo, err := getTodoFromRequest(r)
	db.Save(&todo)
	if err != nil {
		fmt.Fprintln(w, "error")
	} else {
		fmt.Fprintln(w, "ok")
	}
}

func getTodoFromRequest(r *http.Request) (Todo, error) {
	todo := new(Todo)
	return *todo, nil
}
