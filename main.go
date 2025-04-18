package main

import (
	"fmt"
	"net/http"
)

func main() {
	url := "https://safebooru.org/index.php?page=post&s=view&id=5708103"
	req, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println(req)
	}
}
