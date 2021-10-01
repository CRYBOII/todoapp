package main

import (
	"bufio"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

func main() {
	log.SetFlags(log.Lmicroseconds)
	http.HandleFunc("/todos", todosHandler)
	http.HandleFunc("/todos/new", newHandler)
	http.HandleFunc("/todos/create", createHandler)
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))
	http.ListenAndServe(":8080", nil)
}

func todosHandler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("index.html", "todos.html"))

	t.Execute(w, readLines("todos.txt"))
}
func newHandler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("index.html", "new.html"))

	t.Execute(w, nil)
}
func createHandler(w http.ResponseWriter, r *http.Request) {

	f, err := os.OpenFile("todos.txt", os.O_APPEND, os.ModeAppend)
	check(err)
	_, err = fmt.Fprintln(f, r.FormValue("item"))

	check(err)
	http.Redirect(w, r, "/todos", http.StatusFound)
}

func check(err error) {
	if err != nil {
		panic(err)
	}

}

func readLines(fn string) []string {

	f, err := os.Open(fn)
	if os.IsNotExist(err) {
		return nil
	}
	check(err)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	check(scanner.Err())

	return lines

}
