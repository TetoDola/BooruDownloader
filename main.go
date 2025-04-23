package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func check_error(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	// Fetch HTML
	url := "https://safebooru.org/index.php?page=post&s=view&id=5708103"
	response, err := http.Get(url)
	check_error(err)
	_, _ = fmt.Println(response) // the underscore is a way to say ik it returns something and something but i will ignore it
	defer response.Body.Close()

	// Parse HTML
	html, err := io.ReadAll(response.Body)
	check_error(err)
	html_str := string(html)
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html_str))
	check_error(err)
	image_url := doc.Find("meta[property='og:image']").AttrOr("content", "")
	fmt.Println(image_url)

	// Fetch Img + Download
	download, err := http.Get(image_url)
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
	return
}

// TODO: Properly handle error
// TODO: Do it in a loop between image ids
// ToDO: Create Folder before Saving
