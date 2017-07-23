package main

import (
	"encoding/json"
	"fmt"
	"github.com/nanobox-io/golang-scribble"
	"html/template"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"
)

const (
	PORT_ADDR    = ":3000"
	TODO_DB_ADDR = "todos"
)

var db, err = scribble.New("scribble", nil)

func generateTodoFromName(name string) Todo {
	rand.Seed(time.Now().UnixNano())
	return Todo{
		name,
		false,
		rand.Int63(),
	}
}

func generateJsonStringArray(records []string) string {
	data := "["
	length := len(records) - 1
	for pos, record := range records {
		data += record
		if pos != length {
			data += ","
		}
	}

	data += "]"
	return data
}

// TODO: реализовать удаление по id елемента!!!
func addNewTodo(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	bodyString := string(body)
	fmt.Println(bodyString)
	if err != nil {
		err.Error()
	}
	var resp Resp
	json.Unmarshal(body, &resp)
	fmt.Println("Respname ", resp.Name)
	db.Write(TODO_DB_ADDR, resp.Name, generateTodoFromName(resp.Name))

	fmt.Fprintf(w, bodyString)
}

func getAllPosts(w http.ResponseWriter, r *http.Request) {
	records, err := db.ReadAll(TODO_DB_ADDR)

	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	fmt.Fprintf(w,
		generateJsonStringArray(records))

}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/index.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	err = t.ExecuteTemplate(w, "index", nil)
	if err != nil {
		log.Println(err.Error())
	}

}

func main() {

	if err != nil {
		fmt.Println("Error", err)
	}

	http.Handle("/dist/", http.StripPrefix(
		"/dist/",
		http.FileServer(http.Dir("./dist/"))))

	http.HandleFunc("/", mainHandler)
	http.HandleFunc("/get_all", getAllPosts)
	http.HandleFunc("/add_new", addNewTodo)

	fmt.Println("Listening on port :3000")

	http.ListenAndServe(PORT_ADDR, nil)
}
