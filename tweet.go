package main

import (
	"context"

	"github.com/michimani/gotwi"
	"github.com/michimani/gotwi/tweet/managetweet"
	"github.com/michimani/gotwi/tweet/managetweet/types"
)

func SendTweet(client *gotwi.Client, s string, tweetID string) (string, error) {
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
