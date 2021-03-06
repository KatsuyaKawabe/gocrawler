package crawler

import (
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"os"
)

func Crawl() error {
	doc, err := goquery.NewDocument("http://www.ise.shibaura-it.ac.jp")
	if err != nil {
		fmt.Println("url scraping failed")
	}

	res, err := doc.Html()
	if err != nil {
		fmt.Println("dom get failed")
	}

	getLink(doc)
	//ioutil.WriteFile("./sample.html", []byte(res), os.ModePerm)
	return nil
}

func getLink(doc *goquery.Document) ([]string, error) {
	var url_list []string
	doc.Find("a").Each(func(_ int, s *goquery.Selection) {
		url, _ := s.Ttr("href")
		url_list = append(url_list, url)
	})
	if len(url_list) > 256 {
		return url_list[0:255], errors.New("url_list is too long")
	}
	return url_list, nil
}
