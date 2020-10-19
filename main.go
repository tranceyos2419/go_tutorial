package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Ttodo struct {
	UserId    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

type DocumentTone struct {
	Tones []Tone `json:"tones"`
}

type Tone struct {
	Score    float64 `json:"score"`
	ToneId   string  `json:"tone_id"`
	ToneName string  `json:"tone_name"`
}

func main() {
	testPost()
	// testGet()

}

func testPost() {
	url := "http://localhost:8080/analyze"

	text := "Fuck off. I don't know who you are.sales have been disappointing for the past three quarters. We have a competitive product, but we need to do a better job of selling it! "
	values := map[string]string{"text": text}

	jsonValue, _ := json.Marshal(values)

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonValue))

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))

	// var documentTone DocumentTone
	var tones []Tone
	json.Unmarshal(body, &tones)
	for i, tone := range tones {
		fmt.Println(i, tone)
	}
}

func testGet() {
	url := "https://jsonplaceholder.typicode.com/todos"

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

	//* want to use a json response
	var todos []Ttodo
	json.Unmarshal(body, &todos)

	if err != nil {
		log.Fatal(err)
	}

	for i, todo := range todos {
		fmt.Println(i, todo.Title)
	}
}
