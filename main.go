package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	url := "https://jsonplaceholder.typicode.com/todos/1"

	resp, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {

		panic(err)
	}

	fmt.Println(string(body))
}
