package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/michimani/gotwi"
	"github.com/michimani/gotwi/tweet/managetweet"
	"github.com/michimani/gotwi/tweet/managetweet/types"
)

type Result struct {
	Words []struct {
		Term         string `json:"term"`
		Definition   string `json:"definition"`
		Example      string `json:"example"`
		Partofspeech string `json:"partofspeech"`
		Synonyms     string `json:"synonyms"`
		Antonyms     string `json:"antonyms"`
	} `json:"result"`
}

func tweet(client *gotwi.Client, s string, tweetID string) (string, error) {
	message := &types.CreateInput{
		Text: gotwi.String(s),
		Reply: &types.CreateInputReply{
			InReplyToTweetID: tweetID,
		},
	}
	res, err := managetweet.Create(context.Background(), client, message)
	if err != nil {
		return "", err
	}

	return gotwi.StringValue(res.Data.ID), nil
}

func sanitize(tweet string) string {
	//cleaning and normalizing word to use
	tweet = strings.ReplaceAll(tweet, "@AutomatedAndy", "")
	tweet = strings.ReplaceAll(tweet, " ", "")
	tweet = strings.ToLower(tweet)
	return tweet
}

func getWordDetails(w string) (Result, error) {
	token := os.Getenv("API_TOKEN")
	uid := os.Getenv("API_UID")

	url := fmt.Sprintf("https://www.stands4.com/services/v2/syno.php?uid=%s&tokenid=%s&word=%s&format=json", uid, token, w)
	res, err := http.Get(url)
	if err != nil {
		fmt.Println("error")
		return Result{}, err
	}

	responseData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var result Result
	err = json.Unmarshal(responseData, &result)
	if err != nil {
		fmt.Println("There was an error unMarshalling the data.")
		return Result{}, err
	}

	return result, nil
}
