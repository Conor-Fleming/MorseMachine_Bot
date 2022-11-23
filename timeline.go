package main

import (
	"context"

	"github.com/michimani/gotwi"
	"github.com/michimani/gotwi/fields"
	"github.com/michimani/gotwi/tweet/timeline"
	"github.com/michimani/gotwi/tweet/timeline/types"
)

func getMentions(client *gotwi.Client, id string) (string, error) {
	input := &types.ListMentionsInput{
		ID:          userId,
		MaxResults:  5,
		TweetFields: fields.TweetFieldList{fields.TweetFieldText},
	}
	res, err := timeline.ListMentions(context.Background(), client, input)
	if err != nil {
		return "", err
	}
	count := 0
	result := ""
	for _, v := range res.Data {
		if count > 1 {
			break
		}
		result = gotwi.StringValue(v.Text)
		count++
	}

	return result, nil
}
