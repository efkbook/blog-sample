package scenario

import (
	"log"
	"math/rand"
	"net/http"
	"path"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func Crawler() error {
	for {
		doc, err := goquery.NewDocument(indexURL)
		if err != nil {
			return err
		}
		doc.Find("a").Each(func(_ int, s *goquery.Selection) {
			url, _ := s.Attr("href")
			if strings.HasPrefix(url, "/") {
				if rand.Intn(rand.Intn(10)+1) != 0 {
					return
				}
				log.Printf("crawler: get %v", url)
				http.Get(path.Join(indexURL, url))
			}
			time.Sleep(100 * time.Millisecond)
		})
		time.Sleep(10 * time.Second)
	}
}
