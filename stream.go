package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	twitterstream "github.com/fallenstedt/twitter-stream"
	"github.com/fallenstedt/twitter-stream/stream"
	"github.com/michimani/gotwi"
)

type StreamDataExample struct {
	Data struct {
		Text      string    `json:"text"`
		ID        string    `json:"id"`
		CreatedAt time.Time `json:"created_at"`
		AuthorID  string    `json:"author_id"`
	} `json:"data"`
	Includes struct {
		Users []struct {
			ID       string `json:"id"`
			Name     string `json:"name"`
			Username string `json:"username"`
		} `json:"users"`
	} `json:"includes"`
	MatchingRules []struct {
		ID  string `json:"id"`
		Tag string `json:"tag"`
	} `json:"matching_rules"`
}

func newOAuth1Client(accessToken, accessSecret string) (*gotwi.Client, error) {
	client := &gotwi.NewClientInput{
		AuthenticationMethod: gotwi.AuthenMethodOAuth1UserContext,
		HTTPClient: &http.Client{
			Timeout: 0,
		},
		OAuthToken:       accessToken,
		OAuthTokenSecret: accessSecret,
	}

	return gotwi.NewClient(client)
}

func initiateStream() {
	fmt.Println("Authenticating...")
	accessToken := os.Getenv("ACCESS_TOKEN")
	accessSecret := os.Getenv("ACCESS_TOKEN_SECRET")

	if accessToken == "" || accessSecret == "" {
		fmt.Fprintln(os.Stderr, "Please set the ACCESS_TOKEN and ACCESS_SECRET environment variables.")
		os.Exit(1)
	}

	client, err := newOAuth1Client(accessToken, accessSecret)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println("Stream starting...", client)

	api := getTweets()
	defer getTweets()

	for tweet := range api.GetMessages() {
		if tweet.Err != nil {
			fmt.Printf("got error from twitter: %v", tweet.Err)

			api.StopStream()
			continue
		}

		tweetBody := tweet.Data.(StreamDataExample).Data.Text
		tweetID := tweet.Data.(StreamDataExample).Data.ID
		translation, err := getTranslation(tweetBody)
		if err != nil {
			fmt.Println("err")
		}
		fmt.Println(translation)

		replyId, err := SendTweet(client, translation, tweetID)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			api.StopStream()
			continue
		}
		//indicates success
		fmt.Println("Tweet ID:", replyId)
	}
	fmt.Println("Stopped Stream")
}

func getTweets() stream.IStream {
	//Authentication
	tok, err := twitterstream.NewTokenGenerator().SetApiKeyAndSecret(os.Getenv("GOTWI_API_KEY"), os.Getenv("GOTWI_API_KEY_SECRET")).RequestBearerToken()
	if err != nil {
		panic(err)
	}
	//create twitterstream api instance
	api := twitterstream.NewTwitterStream(tok.AccessToken).Stream

	api.SetUnmarshalHook(func(bytes []byte) (interface{}, error) {
		data := StreamDataExample{}

		if err := json.Unmarshal(bytes, &data); err != nil {
			fmt.Printf("failed to unmarshal bytes: %v", err)
		}

		return data, err
	})

	streamExpansions := twitterstream.NewStreamQueryParamsBuilder().
		AddExpansion("author_id").
		AddTweetField("created_at").
		Build()

	err = api.StartStream(streamExpansions)
	if err != nil {
		panic(err)
	}

	return api
}
