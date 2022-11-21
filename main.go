package main

import (
	"fmt"
	"log"
	"os"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

type Credentials struct {
	ConsumerKey       string
	ConsumerSecret    string
	AccessToken       string
	AccessTokenSecret string
}

func main() {
	fmt.Println("Twitt Talk Bot v0.01")
	creds := Credentials{
		AccessToken:       os.Getenv("ACCESS_TOKEN"),
		AccessTokenSecret: os.Getenv("ACCESS_TOKEN_SECRET"),
		ConsumerKey:       os.Getenv("CONSUMER_KEY"),
		ConsumerSecret:    os.Getenv("CONSUMER_SECRET"),
	}

	client, err := getClient(&creds)
	if err != nil {
		log.Println("Error getting client")
		log.Println(err)
	}

	//tweet, resp, err := client.Statuses.Update("TEST TEST TEST - FROM A BOT.", nil)
	//if err != nil {
	//	log.Println(err)
	//}
	//log.Printf("%+v\n", tweet)
	log.Printf("%+v\n", client)
}

func getClient(creds *Credentials) (*twitter.Client, error) {
	/*config := &clientcredentials.Config{
		ClientID:     creds.ConsumerKey,
		ClientSecret: creds.ConsumerSecret,
		TokenURL:     "https://api.twitter.com/oauth2/token",
	}*/
	config := oauth1.NewConfig(creds.ConsumerKey, creds.ConsumerSecret)

	token := oauth1.NewToken(creds.AccessToken, creds.AccessTokenSecret)

	httpClient := config.Client(oauth1.NoContext, token)
	client := twitter.NewClient(httpClient)

	verifyParams := &twitter.AccountVerifyParams{
		SkipStatus:   twitter.Bool(true),
		IncludeEmail: twitter.Bool(true),
	}

	user, _, err := client.Accounts.VerifyCredentials(verifyParams)
	if err != nil {
		return nil, err
	}
	log.Printf("Accout: \n%+v\n", user)

	return client, nil
}
