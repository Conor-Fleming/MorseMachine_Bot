package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/michimani/gotwi"
)

func main() {
	fmt.Println("Automated Andy Bot v0.01")

	accessToken := os.Getenv("ACCESS_TOKEN")
	accessSecret := os.Getenv("ACCESS_TOKEN_SECRET")
	userID := os.Getenv("USER_ID")

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

	CreateSearchRules(client)

	//call to getMentions() retrieves the most recent tweet the bot has been mentioned in
	// return values are the text of that tweet as well as the tweet id to enable replying

	//tweetBody, tweetID, err := getMentions(client, userID)
	//use filtered stream to get this info and then send tweets

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}

	//call tweet func with message as response from synonyms api call
	payload, err := getWordDetails(sanitize(tweetBody))
	if err != nil {
		log.Println(err)
	}

	for _, v := range payload.Words {
		formattedTweet := fmt.Sprintf("Word: %s\n\tDefinition: %s\n\tPart of Speech: %s\n\tSynonyms: %s\n\n", v.Term, v.Definition, v.Partofspeech, v.Synonyms)
		fmt.Println(formattedTweet)
		replyId, err := tweet(client, formattedTweet, tweetID)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(2)
		}
		//indicates success
		fmt.Println("Tweet ID:", replyId)
	}

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
