package main

import (
	"context"

	"github.com/michimani/gotwi"
	"github.com/michimani/gotwi/fields"
	"github.com/michimani/gotwi/tweet/timeline"
	"github.com/michimani/gotwi/tweet/timeline/types"
)

func getMentions(client *gotwi.Client, userId string) (string, string, error) {
	input := &types.ListMentionsInput{
		ID:          userId,
		MaxResults:  5,
		TweetFields: fields.TweetFieldList{fields.TweetFieldText},
	}
	res, err := timeline.ListMentions(context.Background(), client, input)
	if err != nil {
		return "", "", err
	}

	body := gotwi.StringValue(res.Data[0].Text)
	tweetID := gotwi.StringValue(res.Data[0].ID)

	return body, tweetID, nil
}
