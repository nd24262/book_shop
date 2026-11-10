package main

import (
	"html/template"
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "Ошибка загрузки страницы", http.StatusInternalServerError)
		log.Println(err)
		return
	}

	err = tmpl.Execute(w, nil)
}

func main() {
	http.HandleFunc("/", homePage)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	log.Println("Сервер запущен: http://localhost:8080")
	http.ListenAndServe(":8080", nil)

}
