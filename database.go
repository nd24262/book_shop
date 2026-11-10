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
