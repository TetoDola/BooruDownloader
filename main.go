package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func check_error(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func download_image(url string) {
	download, err := http.Get(url)
	check_error(err)
	file, err := os.Create("image.jpg")
	check_error(err)
	DownloadedImage, err := io.Copy(file, download.Body)
	check_error(err)
	err = file.Close()
	check_error(err)
	fmt.Println(download, DownloadedImage)
	err = download.Body.Close()
	check_error(err)
}

func parse_html(html string) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	check_error(err)
	image_url := doc.Find("meta[property='og:image']").AttrOr("content", "")
	fmt.Println(image_url)
	download_image(image_url)
}
func main() {
	// Fetch HTML
	url := "https://safebooru.org/index.php?page=post&s=view&id=5708103"
	response, err := http.Get(url)
	check_error(err)
	_, _ = fmt.Println(response) // the underscore is a way to say ik it returns something and something but i will ignore it
	html, err := io.ReadAll(response.Body)
	check_error(err)
	parse_html(string(html))
	defer response.Body.Close()

}

// TODO: Properly handle error
// TODO: Do it in a loop between image ids
// ToDO: Create Folder before Saving
