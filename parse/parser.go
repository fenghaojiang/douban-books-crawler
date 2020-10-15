package parse

import (
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

type DoubanBook struct {
	Title      string
	Author     string
	Translator string
	Press      string
	Date       string
	Price      string
	Star       string
	Comment    string
	Quote      string
}

type Page struct {
	Page int
	Url  string
}

//伪造请求头
func GetDoc(url string) *goquery.Document {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.75 Safari/537.36")

	rsp, err := client.Do(req)
	defer rsp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	if rsp.StatusCode != 200 {
		log.Fatalf("%s", rsp.Status)
	}

	doc, err := goquery.NewDocumentFromReader(rsp.Body)
	return doc
}

// 分析分页
func ParsePage(doc *goquery.Document) (pages []Page) {
	pages = append(pages, Page{Page: 1, Url: ""})
	doc.Find(".paginator > a").Each(func(i int, s *goquery.Selection) {
		page, _ := strconv.Atoi(s.Text())
		url, _ := s.Attr("href")

		pages = append(pages, Page{
			Page: page,
			Url:  url,
		})
	})

	return pages
}

func ParseBook(doc *goquery.Document) (books []DoubanBook) {
	doc.Find("#content > div > div.indent > table > tbody > tr > td").Each(func(i int, selection *goquery.Selection) {
		title := strings.TrimSpace(selection.Find("td a").Text())

		info := selection.Find("td .pl").Text()
		bookInfo := strings.Split(info, "/")
		author := strings.TrimRight(bookInfo[0], "著")
		author = strings.TrimSpace(author)
		var translator, press, date, price string
		var press string
		if len(bookInfo) == 4 {
			press = strings.TrimSpace(bookInfo[1])
			date = strings.TrimSpace(bookInfo[2])
			price = strings.TrimSpace(bookInfo[3])
		} else {
			translator = strings.TrimSpace(bookInfo[1])
			press = strings.TrimSpace(bookInfo[2])
			date = strings.TrimSpace(bookInfo[3])
			price = strings.TrimSpace(bookInfo[4])
		}

		star := selection.Find("td div span").Eq(1).Text()
		comment := selection.Find("td div span").Eq(2).Text()
		compile := regexp.MustCompile("[0-9]")
		comment = strings.Join(compile.FindAllString(comment, -1), "")

		quote := selection.Find("td .quote .inq").Text()

		book := DoubanBook{
			Title:      title,
			Author:     author,
			Translator: translator,
			Press:      press,
			Date:       date,
			Price:      price,
			Star:       star,
			Comment:    comment,
			Quote:      quote,
		}

		log.Printf("i: %d, book: %v\n", i, book)

		books = append(books, book)
	})

	return books
}
