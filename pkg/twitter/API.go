package twitter

import (
	"context"

	twitterscraper "github.com/n0madic/twitter-scraper"
)

func GetTweets(username string) <-chan *twitterscraper.Result {
	scraper := twitterscraper.New()
	scraper.SetSearchMode(twitterscraper.SearchLatest)
	scraper.WithReplies(true)

	tweets := scraper.GetTweets(context.Background(), username, 5)

	return tweets
}
