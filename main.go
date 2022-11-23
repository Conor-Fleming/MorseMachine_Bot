package main

import (
	"fmt"
	"os"

	"github.com/michimani/gotwi"
	//"github.com/michimani/gotwi/tweet/managetweet"
	//"github.com/michimani/gotwi/tweet/managetweet/types"
)

const userId = "1595276630196248576"

func main() {
	fmt.Println("Twitt Talk Bot v0.01")

	//bearerToken := os.Getenv("BEARER_TOKEN")
	accessToken := os.Getenv("ACCESS_TOKEN")
	accessSecret := os.Getenv("ACCESS_TOKEN_SECRET")

	if accessToken == "" || accessSecret == "" {
		fmt.Fprintln(os.Stderr, "Please set the ACCESS_TOKEN and ACCESS_SECRET environment variables.")
		os.Exit(1)
	}

	client, err := newOAuth1Client(accessToken, accessSecret)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		fmt.Println("test")
		os.Exit(1)
	}

	//tweetId, err := tweet(client, os.Args[1])

	mentionedTweet, err := getMentions(client, userId)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}

	//call tweet func with message as response from synonyms api call
	tweetId, err := tweet(client, sanitize(mentionedTweet))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}

	fmt.Println("Tweet ID:", tweetId)
}

func newOAuth1Client(accessToken, accessSecret string) (*gotwi.Client, error) {
	client := &gotwi.NewClientInput{
		AuthenticationMethod: gotwi.AuthenMethodOAuth1UserContext,
		OAuthToken:           accessToken,
		OAuthTokenSecret:     accessSecret,
	}

	return gotwi.NewClient(client)
}
