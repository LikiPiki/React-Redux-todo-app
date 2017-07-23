package main

type Todo struct {
	Name    string `json:"name"`
	Checked bool   `json:"checked"`
	Todo_id int64  `json:"todo_id"`
}

type Resp struct {
	Name string
}
