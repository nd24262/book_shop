package main

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"
)

var database *sql.DB

func homePage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "Ошибка загрузки страницы", http.StatusInternalServerError)
		log.Println(err)
		return
	}

	err = tmpl.Execute(w, nil)
}

func addBookPage(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		title := r.FormValue("title")
		author := r.FormValue("author")
		price := r.FormValue("price")
		quantity := r.FormValue("quantity")

		_, err := database.Exec(
			"INSERT INTO books (title, author, price, quantity) VALUES (?, ?, ?, ?)",
			title, author, price, quantity,
		)

		if err != nil {
			http.Error(w, "Ошибка добавления книги", http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	tmpl, err := template.ParseFiles("templates/add_book.html")
	if err != nil {
		http.Error(w, "Ошибка загрузки страницы", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, nil)
}

func main() {
	database, err := connectDB()
	if err != nil {
		log.Fatal("Ошибка подключения к базе данных:", err)
	}
	defer database.Close()

	err = createTables(database)
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", homePage)
	http.HandleFunc("/add-book", addBookPage)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	log.Println("Сервер запущен: http://localhost:8080")
	http.ListenAndServe(":8080", nil)

}
