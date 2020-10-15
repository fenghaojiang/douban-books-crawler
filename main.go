package main

import (
	"books/model"
	"books/parse"
	"log"
)

var (
	baseUrl = "https://book.douban.com/top250"
)

func add(books []parse.DoubanBook) {
	for i, book := range books {
		if err := model.DB.Create(&book).Error; err != nil {
			log.Printf("db.Create index: %d, err : %v", i, err)
		}
	}
}

func startCrawler() {
	var books []parse.DoubanBook

	pages := parse.GetPages(baseUrl)
	for _, page := range pages {
		doc := parse.GetDoc(page.Url)
		books = append(books, parse.ParseBook(doc)...)
	}

	add(books)
}

func main() {
	startCrawler()

	defer model.DB.Close()
}
