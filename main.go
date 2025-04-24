package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

// PS Copy pasted from a random website, no clue how this function actually works
func make_dir() error {
	home, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("Get Home Directory: %w", err)
	}
	path := home + "/Documents/Safebooru/images"
	check_existence, err := os.Stat(path)
	switch {
	case os.IsNotExist(err):
		err = os.MkdirAll(path, 0755)
		if err != nil {
			return fmt.Errorf("Create Directory: %w", err)
		}
		return nil
	case check_existence.IsDir():
		return nil

	case !check_existence.IsDir():
		return fmt.Errorf("Path is not a directory %w", path)

	default:
		return fmt.Errorf("Unknown error %w", err)
	}

}

func download_image(url string, id int) error {
	download, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return fmt.Errorf("Image %w Downloading error", id)
	}

	file, err := os.Create(fmt.Sprintf("C:\\Users\\Teto\\Documents\\Safebooru\\images\\%v.jpg", id))
	if err != nil {
		log.Println(err)
		return fmt.Errorf("Create File: %w", err)
	}
	_, err = io.Copy(file, download.Body)
	if err != nil {
		log.Println(err)
		return fmt.Errorf("Copy File: %w", err)
	}
	err = file.Close()
	if err != nil {
		log.Println(err)
		return fmt.Errorf("Close File: %w", err)
	}
	err = download.Body.Close()
	if err != nil {
		log.Println(err)
		return fmt.Errorf("Close Body: %w", err)
	}
	fmt.Printf("Downloaded Image: %v \n", id)
	return nil
}

func parse_html(html string, id int) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		log.Println(err)
	}
	image_url := doc.Find("meta[property='og:image']").AttrOr("content", "")
	err = download_image(image_url, id)
	if err != nil {
		fmt.Printf("Error Downloading Image %v \n", id)
	}
}
func main() {
	// Fetch HTML
	err := make_dir()
	if err != nil {
		log.Fatal(err)
	}
	for i := 1; i < 1001; i++ {
		url := fmt.Sprintf("https://safebooru.org/index.php?page=post&s=view&id=%v", i)
		response, err := http.Get(url)
		if err != nil {
			log.Println(err)
		}

		html, err := io.ReadAll(response.Body)
		response.Body.Close()
		if err != nil {
			log.Println(err)
			continue
		}
		parse_html(string(html), i)
		time.Sleep(1000 * time.Millisecond)
	}
}
