package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

func handleFuncMain(w http.ResponseWriter, r *http.Request) {
	tmp, err := template.ParseFiles("templates/index.html")
	if err != nil {
		log.Println(err)
	}
	err = tmp.ExecuteTemplate(w, "index", nil)
	if err != nil {
		log.Println(err)
	}
}

func getAllTodos(w http.ResponseWriter, r *http.Request) {
	var todos []Todo
	db.Find(&todos)
	jsonify, err := json.Marshal(&todos)
	if err != nil {
		log.Println(err)
	}
	w.Write(jsonify)
}

func addTodo(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
	}
	var resp Resp
	err = json.Unmarshal(body, &resp)
	log.Println(resp)
	db.Create(&Todo{
		Name:    resp.Name,
		Checked: false,
	})
	if err != nil {
		fmt.Fprintf(w, "error")
	} else {
		fmt.Fprintf(w, "ok")
	}
}

func deleteTodo(w http.ResponseWriter, r *http.Request) {
	var err error
	if err != nil {
		fmt.Fprintln(w, "error")
	} else {
		fmt.Fprintln(w, "ok")
	}
}

func complitedTodo(w http.ResponseWriter, r *http.Request) {
	var err error
	if err != nil {
		fmt.Fprintln(w, "error")
	} else {
		fmt.Fprintln(w, "ok")
	}
}

func editTodo(w http.ResponseWriter, r *http.Request) {
	var err error
	if err != nil {
		fmt.Fprintln(w, "error")
	} else {
		fmt.Fprintln(w, "ok")
	}
}
