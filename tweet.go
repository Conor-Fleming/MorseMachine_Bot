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

// syn api url: https://www.stands4.com/services/v2/syno.php
// sample: https://www.stands4.com/services/v2/syno.php?uid=1001&tokenid=tk324324&word=consistent&format=xml
type Result struct {
	Word Word
}

type Word struct {
	Term       string
	Definition string
	Speech     string
	Synonyms   string
	Antonyms   string
}

func tweet(client *gotwi.Client, s string, tweetID string) (string, error) {
	repsondingMessage := fmt.Sprintf(`%s - responding to your mention`, s)
	message := &types.CreateInput{
		Text: gotwi.String(repsondingMessage),
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
	return strings.ReplaceAll(tweet, "@AutomatedAndy", "")
}

func getWordDetails(w string) (Word, error) {
	token := os.Getenv("API_TOKEN")
	uid := os.Getenv("API_UID")

	url := fmt.Sprintf("https://www.stands4.com/services/v2/syno.php?uid=%s&tokenid=%s&word=%s&format=json", uid, token, w)
	res, err := http.Get(url)
	if err != nil {
		fmt.Println("error")
		return Word{}, err
	}

	responseData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var word Word
	json.Unmarshal(responseData, &word)

	return word, nil
}
