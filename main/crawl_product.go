package crawlproduct

import (
	"fmt"
	"strings"

	"./dbconnection"
	"./savedata"
	"github.com/PuerkitoBio/goquery"
	"github.com/buger/jsonparser"
)

func OptimizeUrl(value string) (url string) {
	if strings.Index(value, "http") == 0 {
		url = value
		return url
	} else {
		url = "https://www.walmart.com" + value
		return url
	}
}
func main() {
	q, err := dbconnection.Connect.Query("select link from all_links")
	if err != nil {
		fmt.Println("ERROR: ", err)
	}
	var value string
	for q.Next() {
		// get link type string in mysql
		err := q.Scan(&value)
		if err != nil {
			fmt.Println("ERROR: ", err)
		}
		// parse link
		doc, err := goquery.NewDocument(value)
		if err != nil {
			fmt.Println("Error: ", err)
			savedata.SaveUrlErrorProduct(value)
		}
		var body string
		doc.Find("head script").Each(func(i int, s *goquery.Selection) {
			var band string
			band = s.Text()
			check := strings.Index(band, "__WML_REDUX_INITIAL_STATE__") != -1
			if check {
				index := strings.Index(band, "{")
				for i := 0; i < index; i++ {
					band = strings.Replace(band, string(band[i]), " ", 1)
				}
				body = band
			}
		})
		data := []byte(body)
		dataProduct, _, _, _ := jsonparser.Get(data, "preso")
		dataPath, _, _, _ := jsonparser.Get(dataProduct, "adContext", "categoryPathName")
		jsonparser.ArrayEach(dataProduct, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
			title, _, _, _ := jsonparser.Get(value, "title")
			imageUrl, _, _, _ := jsonparser.Get(value, "imageUrl")
			url, _, _, _ := jsonparser.Get(value, "productPageUrl")
			link := OptimizeUrl(string(url))
			savedata.SaveData(string(title), string(dataPath), string(imageUrl), link)
		}, "items")
	}
}
