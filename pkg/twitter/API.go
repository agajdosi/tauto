package twitter

import (
	"context"
	"fmt"

	twitterscraper "github.com/n0madic/twitter-scraper"
)

func GetTweets(username string) []string {
	scraper := twitterscraper.New()
	scraper.SetSearchMode(twitterscraper.SearchLatest)
	//scraper.WithReplies(true)

	for tweet := range scraper.GetTweets(context.Background(), username, 5) {
		if tweet.Error != nil {
			panic(tweet.Error)
		}

		fmt.Println(tweet.Text, tweet.PermanentURL)
	}

	return nil
}
