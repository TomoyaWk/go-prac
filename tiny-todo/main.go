package main

import (
	"html/template"
	"log"
	"net/http"
)

var todoList []string

// template
func handleTodo(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/todo.html")
	t.Execute(w, todoList)
}

func main() {
	todoList = append(todoList, "顔洗う", "朝食食べる", "歯を磨く")
	//static
	http.Handle("/static/",
		http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	//todo
	http.HandleFunc("/todo/", handleTodo)

	//start & error handle
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("failed to Start:", err)
	}

}
