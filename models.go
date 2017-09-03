package main

type Todo struct {
	Id      int    `gorm:"primary_key" json:"id"`
	Name    string `json:"name"`
	Checked bool   `json:"checked"`
}

type Resp struct {
	Name string
}
