package main

import "fmt"

type Book struct {
	title  string
	author string
	price  int
}

type BookStore struct {
	BookList []*Book
}

func main() {
	fmt.Println("主線程開始...")

	book1 := Book{title: "Vue", author: "尤大神", price: 20}
	book2 := Book{title: "React", author: "rinrin", price: 20}
	book3 := Book{title: "Golang", author: "rin大神", price: 20}

	bookStore := BookStore{}
	bookStore.BookList = append(bookStore.BookList, &book1, &book2, &book3)

	for i, book := range bookStore.BookList {
		fmt.Println(i, book)
	}
}
