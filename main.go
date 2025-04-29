package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func make_dir() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("Get Home Directory: %w", err)
	}
	path := home + "/Documents/Safebooru/images"
	check_existence, err := os.Stat(path)
	switch {
	case os.IsNotExist(err):
		err = os.MkdirAll(path, 0755)
		if err != nil {
			return "", fmt.Errorf("Create Directory: %w", err)
		}
		return path, nil
	case check_existence.IsDir():
		return path, nil

	case !check_existence.IsDir():
		return "", fmt.Errorf("Path is not a directory %w", path)

	default:
		return "", fmt.Errorf("Unknown error %w", err)
	}

}

func download_image(url string, id int, path string) error {
	download, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return fmt.Errorf("Image %w Downloading error", id)
	}

	file, err := os.Create(filepath.Join(path, fmt.Sprintf("%v.jpg", id)))
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

func parse_html(html string, id int, path string) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		log.Println(err)
	}
	image_url := doc.Find("meta[property='og:image']").AttrOr("content", "")
	err = download_image(image_url, id, path)
	if err != nil {
		fmt.Printf("Error Downloading Image %v \n", id)
	}
}
func main() {
	var id1 int
	var id2 int
	// Fetch HTML
	path, err := make_dir()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Enter the ID you want tor download from; ")
	_, err = fmt.Scanln(&id1)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Enter the ID you want tor download to; ")
	_, err = fmt.Scanln(&id2)
	if err != nil {
		log.Fatal(err)
	}

	for i := id1; i < id2+1; i++ {
		url := fmt.Sprintf("https://safebooru.org/index.php?page=post&s=view&id=%v", i)
		response, err := http.Get(url)
		if err != nil {
			log.Println(err)
		}

		html, err := io.ReadAll(response.Body)
		if err != nil {

			log.Println(err)
			continue
		}
		err = response.Body.Close()
		if err != nil {
			log.Println(err)
		}

		parse_html(string(html), i, path)
		time.Sleep(1000 * time.Millisecond)
	}
}
