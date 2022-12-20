package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

type Result struct {
	Success struct {
		Total int `json:"total"`
	} `json:"success"`
	Contents struct {
		Translated  string `json:"translated"`
		Text        string `json:"text"`
		Translation string `json:"translation"`
	} `json:"contents"`
}

func sanitize(tweet string) string {
	//cleaning and normalizing word to use
	tweet = strings.ReplaceAll(tweet, "@AutomatedAndy", "")
	tweet = strings.ToLower(tweet)
	return tweet
}

func getTranslation(payload string) (string, error) {
	client := &http.Client{}

	payload = sanitize(payload)
	morseAPI := os.Getenv("MORSE_KEY")

	var data = strings.NewReader(`text=` + payload)
	req, err := http.NewRequest("POST", "https://api.funtranslations.com/translate/morse.json", data)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("X-Funtranslations-API-Secret", morseAPI)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(responseData))

	var result Result
	err = json.Unmarshal(responseData, &result)
	if err != nil {
		fmt.Println("There was an error unMarshalling the data.")
		return "", err
	}

	//format data for tweeting
	output := formatTweet(result)

	return output, nil
}

func formatTweet(data Result) string {
	fmt.Println(string(data.Contents.Text))

	return ""
}
