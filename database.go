package main

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

func connectDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite", "./book_shop.db")
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func createTables(db *sql.DB) error {
	//Таблицу пользователей
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT,
		email TEXT,
		password TEXT,
		role TEXT
	)`)
	if err != nil {
		return err
	}

	//Таблица категорий
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS categories (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT
	)`)
	if err != nil {
		return err
	}

	//Таблица книг
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS books (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT,
		author TEXT,
		price REAL,
		quantity INTEGER,
		category_id INTEGER
	)`)
	if err != nil {
		return err
	}

	//Таблица заказов
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS orders (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id INTEGER,
		order_date TEXT,
		total REAL
	)`)
	if err != nil {
		return err
	}

	return nil
}

/*Представляем в каталоге*/
type Book struct {
	ID           int
	Title        string
	Author       string
	Price        float64
	Quantity     int
	CategoryName string
}

func getBooks(db *sql.DB) ([]Book, error) {
	rows, err := db.Query(`
    SELECT books.id, books.title, books.author, books.price,
           books.quantity, categories.name
    FROM books
    LEFT JOIN categories ON books.category_id = categories.id
`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []Book

	for rows.Next() {
		var book Book
		rows.Scan(&book.ID, &book.Title, &book.Author, &book.Price, &book.Quantity, &book.CategoryName)
		books = append(books, book)
	}

	return books, nil
}

/*Для страницы, чтобы добавлять категории*/
type Category struct {
	ID   int
	Name string
}

func getCategories(db *sql.DB) ([]Category, error) {
	rows, err := db.Query("SELECT id, name FROM categories")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []Category

	for rows.Next() {
		var category Category
		rows.Scan(&category.ID, &category.Name)
		categories = append(categories, category)
	}

	return categories, nil
}
