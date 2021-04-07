package twitter

import (
	"context"

	twitterscraper "github.com/n0madic/twitter-scraper"
)

func GetTweets(username string) []string {
	scraper := twitterscraper.New()
	scraper.SetSearchMode(twitterscraper.SearchLatest)
	scraper.WithReplies(true)

	var tweets []string
	for tweet := range scraper.GetTweets(context.Background(), username, 5) {
		if tweet.Error != nil {
			panic(tweet.Error)
		}
		tweets = append(tweets, tweet.PermanentURL)
	}

	return tweets
}
