package main

import (
	"fmt"
	"io"
	"net/http"
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
		_, _ = fmt.Println(string(html))
	}
}

// TODO: Now that we fetched the html, we need to identify main booru image and download it somehow
