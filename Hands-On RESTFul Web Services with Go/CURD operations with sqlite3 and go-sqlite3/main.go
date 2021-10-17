package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

type Book struct {
	id     int
	name   string
	author string
	isbn   int64
}

func main() {
	db, err := sql.Open("sqlite3", "./db/books.db")
	if err != nil {
		log.Fatal(err)
	}

	statement, err := db.Prepare("create table if not exists books (id primary key, name varchar(64) null, author varchar(64), isbn integer);")
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("sqlite3: successfully prepare the create table statement")
	}

	_, err = statement.Exec()
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("sqlite3: successfully create books table")
	}

	dbOperations(db)
}

func dbOperations(db *sql.DB) {
	// 创建
	statement, err := db.Prepare("insert into books (id, name, author, isbn) values (?, ?, ?, ?);")
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("sqlite3: successfully prepare the insert statement")
	}

	_, err = statement.Exec(1, "燃烧的大洋 1941-1942", "Ian W.Toll", 1234)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("successfully insert a book record")
	}

	// 读取
	var tmpBook Book
	rows, err := db.Query("select id, name, author, isbn from books;")
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("sqlite3: successfully read book records")
	}
	for rows.Next() {
		if err = rows.Scan(&tmpBook.id, &tmpBook.name, &tmpBook.author, &tmpBook.isbn); err != nil {
			log.Fatal(err)
		}
		log.Printf("book - id: %d, name: %s, author: %s, isbn: %d\n", tmpBook.id, tmpBook.name, tmpBook.author, tmpBook.isbn)
	}

	// 更新
	statement, err = db.Prepare("update books set isbn=? where id=?")
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("sqlite3: successfully prepare the update statement")
	}
	_, err = statement.Exec(2345, 1)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("sqlite3: successfully update a book record")
	}

	// 删除
	statement, err = db.Prepare("delete from books where id=?")
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("sqlite3: successfully prepare the delete statement")
	}
	_, err = statement.Exec(1)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("sqlite3: successfully delete a book record")
	}
}
