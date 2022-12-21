package main

import (
	"fmt"
	"os"

	twitterstream "github.com/fallenstedt/twitter-stream"
	"github.com/fallenstedt/twitter-stream/rules"
)

func AddRules() {
	tok, err := twitterstream.NewTokenGenerator().SetApiKeyAndSecret(os.Getenv("GOTWI_API_KEY"), os.Getenv("GOTWI_API_KEY_SECRET")).RequestBearerToken()
	if err != nil {
		fmt.Println("bad token")
		panic(err)
	}
	api := twitterstream.NewTwitterStream(tok.AccessToken)
	rules := twitterstream.NewRuleBuilder().AddRule("@TheMorseMachine", "Tweets that mention me").Build()

	res, err := api.Rules.Create(rules, false)
	if err != nil {
		fmt.Println("create rule")
		panic(err)
	}
	if res.Errors != nil && len(res.Errors) > 0 {
		panic(fmt.Sprintf("Recieved Error from Twitter: %v", res.Errors))
	}
	fmt.Println(res.Data)
}

func listRules() {
	tok, err := twitterstream.NewTokenGenerator().SetApiKeyAndSecret(os.Getenv("GOTWI_API_KEY"), os.Getenv("GOTWI_API_KEY_SECRET")).RequestBearerToken()
	if err != nil {
		panic(err)
	}

	api := twitterstream.NewTwitterStream(tok.AccessToken)
	res, err := api.Rules.Get()
	if err != nil {
		panic(err)
	}

	if res.Errors != nil && len(res.Errors) > 0 {
		panic(fmt.Sprintf("Received an error from twitter: %v", res.Errors))
	}

	if len(res.Data) > 0 {
		fmt.Println("Rules: ")
		for _, data := range res.Data {
			fmt.Printf("ID: %v\n", data.Id)
			fmt.Printf("Tag: %v\n", data.Tag)
			fmt.Printf("Data: %v\n", data.Value)
		}
	} else {
		fmt.Println("No rules found.")
	}
}

func deleteRules() {
	tok, err := twitterstream.NewTokenGenerator().SetApiKeyAndSecret(os.Getenv("GOTWI_API_KEY"), os.Getenv("GOTWI_API_KEY_SECRET")).RequestBearerToken()
	if err != nil {
		panic(err)
	}

	api := twitterstream.NewTwitterStream(tok.AccessToken)

	// use api.Rules.Get to find the ID number for an existing rule
	res, err := api.Rules.Delete(rules.NewDeleteRulesRequest( /* enter rule id to delete */ ), false)

	if err != nil {
		panic(err)
	}

	if res.Errors != nil && len(res.Errors) > 0 {
		//https://developer.twitter.com/en/support/twitter-api/error-troubleshooting
		panic(fmt.Sprintf("Received an error from twitter: %v", res.Errors))
	}

	fmt.Println("I have deleted rules")
}
