import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

//$ Post
type ResultFromToneAnalyzerAPI struct {
	DocumentTone DocumentTone   `json:"document_tone"`
	SentenceTone []SentenceTone `json:"sentences_tone"`
}
type DocumentTone struct {
	Tones []Tone `json:"tones"`
}

type Tone struct {
	Score    float64 `json:"score"`
	ToneId   string  `json:"tone_id"`
	ToneName string  `json:"tone_name"`
}

type SentenceTone struct {
	SentenceId float64 `json:"sentence_id"`
	Text       string  `json:"text"`
	Tones      []Tone  `json:"tones"`
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

	// fmt.Println(string(body))

	var result ResultFromToneAnalyzerAPI
	json.Unmarshal(body, &result)

	fmt.Println(result)
	// for i, tone := range result.SentenceTone.Tones {
	// 	fmt.Println(i, tone)
	// }
}
