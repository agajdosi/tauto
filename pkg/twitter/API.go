package twitter

import (
	"context"

	twitterscraper "github.com/n0madic/twitter-scraper"
)

func GetTweets(username string) <-chan *twitterscraper.Result {
	scraper := twitterscraper.New()
	scraper.SetSearchMode(twitterscraper.SearchLatest)
	scraper.WithReplies(true)

	return scraper.GetTweets(context.Background(), username, 5)
}

func GetUserID(username string) (string, error) {
	scraper := twitterscraper.New()
	profile, err := scraper.GetProfile(username)

	return profile.UserID, err
}
