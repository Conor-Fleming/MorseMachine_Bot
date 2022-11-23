package main

import (
	"fmt"
	"os"

	"github.com/michimani/gotwi"
	//"github.com/michimani/gotwi/tweet/managetweet"
	//"github.com/michimani/gotwi/tweet/managetweet/types"
)

func main() {
	fmt.Println("Automated Andy Bot v0.01")

	userID := os.Getenv("USER_ID")
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

	//call to getMentions() retrieves the most recent tweet the bot has been mentioned in
	// return values are the text of that tweet as well as the tweet id to enable replying
	tweetBody, tweetID, err := getMentions(client, userID)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}

	//call tweet func with message as response from synonyms api call
	tweetId, err := tweet(client, sanitize(tweetBody), tweetID)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}

	//indicated success
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
