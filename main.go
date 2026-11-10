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
	books, err := getBooks(database)
	if err != nil {
		http.Error(w, "Ошибка получения книг", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, books)
}

func addBookPage(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		title := r.FormValue("title")
		author := r.FormValue("author")
		price := r.FormValue("price")
		quantity := r.FormValue("quantity")
		categoryID := r.FormValue("category_id")

		_, err := database.Exec(
			"INSERT INTO books (title, author, price, quantity, category_id) VALUES (?, ?, ?, ?, ?)",
			title, author, price, quantity, categoryID,
		)

		if err != nil {
			log.Println(err)
			http.Error(w, "Ошибка добавления книги", http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	categories, err := getCategories(database)
	if err != nil {
		log.Println(err)
		http.Error(w, "Ошибка получения категорий", http.StatusInternalServerError)
		return
	}

	tmpl, err := template.ParseFiles("templates/add_book.html")
	if err != nil {
		log.Println(err)
		http.Error(w, "Ошибка загрузки страницы", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, categories)
	if err != nil {
		log.Println(err)
	}
}

func addCategoryPage(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("name")

		_, err := database.Exec(
			"INSERT INTO categories (name) VALUES (?)",
			name,
		)

		if err != nil {
			log.Println(err)
			http.Error(w, "Ошибка добавления категории", http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/add-book", http.StatusSeeOther)
		return
	}

	tmpl, err := template.ParseFiles("templates/add_category.html")
	if err != nil {
		http.Error(w, "Ошибка загрузки страницы", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

func main() {
	var err error
	database, err = connectDB()
	if err != nil {
		log.Fatal("Ошибка подключения к базе данных:", err)
	}
	defer database.Close()

	err = createTables(database)
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", homePage)
	http.HandleFunc("/add-category", addCategoryPage)
	http.HandleFunc("/add-book", addBookPage)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	log.Println("Сервер запущен: http://localhost:8080")
	http.ListenAndServe(":8080", nil)

}
