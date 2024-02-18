package handler

import (
	"fmt"
	"html/template"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		fmt.Printf("报错: %v\n", err)
		return
	}

	data := PageData{
		Title:   "My Vue Webpage",
		Message: "Hello from the server!",
	}

	tmpl.Execute(w, data)
	if err != nil {

		fmt.Printf("报错: %v\n", err)
		return
	}
}

type PageData struct {
	Title   string
	Message string
}
type ResponseData struct {
	Message string `json:"message"`
}
