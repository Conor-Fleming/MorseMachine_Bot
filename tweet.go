package main

import (
	"context"
	"fmt"
	"strings"

	"github.com/michimani/gotwi"
	"github.com/michimani/gotwi/tweet/managetweet"
	"github.com/michimani/gotwi/tweet/managetweet/types"
)

func tweet(client *gotwi.Client, s string) (string, error) {
	repsondingMessage := fmt.Sprintf(`"%s" responding to your mention`, s)
	message := &types.CreateInput{
		Text: gotwi.String(repsondingMessage),
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
