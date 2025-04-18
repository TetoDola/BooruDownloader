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

func main() {
	url := "https://safebooru.org/index.php?page=post&s=view&id=5708103"
	response, err := http.Get(url)
	if err != nil {
		fmt.Println(err)

		return
	} else {
		_, _ = fmt.Println(response) // the underscore is a way to say ik it returns something and something but i will ignore it
		defer response.Body.Close()
		html, err := io.ReadAll(response.Body)
		if err != nil {
			fmt.Println(err)
			return
		}
		html_str := string(html)
		doc, err := goquery.NewDocumentFromReader(strings.NewReader(html_str))
		if err != nil {
			fmt.Println(err)
			return
		}
		image_url := doc.Find("meta[property='og:image']").AttrOr("content", "")
		fmt.Println(image_url)
		download, err := http.Get(image_url)
		if err != nil {
			log.Fatal(err)
		}
		defer download.Body.Close()
		file, _ := os.Create("image.jpg")
		defer file.Close()
		_, _ = io.Copy(file, download.Body)
		return
	}
	// TODO: Properly handle error
	// TODO: Do it in a loop between image ids
	// ToDO: Create Folder before Saving
}
