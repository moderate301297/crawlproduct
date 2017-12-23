package savedata

import (
	"fmt"

	"../dbconnection"
)

type DataWeb struct {
	Body string
}

// save data to mysql
func SaveData(title string, linkPath string, linkImage string, link string) {
	_, err := dbconnection.ConnectNew.Exec("insert products set title= ?, link_path = ?, link_image = ?, link = ?", title, linkPath, linkImage, link)
	if err != nil {
		fmt.Println("Error: ", err)
	}
}

// save url error
func SaveUrlErrorProduct(url string) {
	_, err := dbconnection.ConnectNew.Exec("insert url_error set url = ?", url)
	if err != nil {
		fmt.Println("Error: ", err)
	}
}
