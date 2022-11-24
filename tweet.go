package main

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/michimani/gotwi"
	"github.com/michimani/gotwi/tweet/managetweet"
	"github.com/michimani/gotwi/tweet/managetweet/types"
)

// syn api url: https://www.stands4.com/services/v2/syno.php
//sample: https://www.stands4.com/services/v2/syno.php?uid=1001&tokenid=tk324324&word=consistent&format=xml

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

func getWordDetails(word string) {
	res, err := http.Get("https://www.stands4.com/services/v2/syno.php?uid=1001&tokenid={********}&word=consistent&format=xml")
}
