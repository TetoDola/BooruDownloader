package main

import (
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func check_error(err error) bool {
	if err != nil {
		log.Println(err)
		return true
	}
	return false
}

func make_dir() {
	path := "C:\\Users\\Teto\\Documents\\Safebooru\\images\\"
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		fmt.Println("Creating directory:", path)
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			log.Println("Failed to create directory:", err)
		}
	}
}

func download_image(url string, id int) {
	download, err := http.Get(url)
	if check_error(err) != true {
		file, err := os.Create(fmt.Sprintf("C:\\Users\\Teto\\Documents\\Safebooru\\images\\%v.jpg", id))
		check_error(err)
		_, err = io.Copy(file, download.Body)
		check_error(err)
		err = file.Close()
		check_error(err)
		err = download.Body.Close()
		check_error(err)
		fmt.Printf("Downloaded Image: %v \n", id)
	}

}

func parse_html(html string, id int) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	check_error(err)
	image_url := doc.Find("meta[property='og:image']").AttrOr("content", "")
	download_image(image_url, id)
}
func main() {
	// Fetch HTML
	make_dir()
	for i := 0; i < 1000; i++ {
		url := fmt.Sprintf("https://safebooru.org/index.php?page=post&s=view&id=%v", i)
		response, err := http.Get(url)
		if check_error(err) {
			continue
		}
		html, err := io.ReadAll(response.Body)
		response.Body.Close()
		if check_error(err) {
			continue
		}
		parse_html(string(html), i)
	}
}

// TODO: Properly handle error
// TODO: Do it in a loop between image ids
// ToDO: Create Folder before Saving
