package twitter

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/chromedp/chromedp"
)

//IsRetweeted checks whether specified tweet is already retweeted by the bot
func (b bot) IsRetweeted(tweetURL string) (bool, error) {
	var value string
	var ok bool

	err := chromedp.Run(*b.ctx,
		chromedp.Navigate(tweetURL),
		chromedp.WaitVisible(`//*[@id="react-root"]/div/div/div[2]/main/div/div/div/div/div/div[2]/div/section/div/div/div[1]/div/div/article/div/div/div/div[3]/div[5]/div[2]/div`, chromedp.BySearch),
		chromedp.AttributeValue(`//*[@id="react-root"]/div/div/div[2]/main/div/div/div/div/div/div[2]/div/section/div/div/div[1]/div/div/article/div/div/div/div[3]/div[5]/div[2]/div`, "aria-label", &value, &ok, chromedp.BySearch),
	)

	if value == "Retweeted" {
		return true, err
	}

	if value == "Retweet" {
		return false, err
	}

	return false, err
}

//EnsureRetweeted unsures that specified tweet is retweeted by the bot
func (b bot) EnsureRetweeted(tweetURL string) {
	liked, err := b.IsRetweeted(tweetURL)
	if err != nil {
		fmt.Println(err)
	}
	time.Sleep(10 * time.Second)
	if liked == true {
		return
	}

	time.Sleep(5 * time.Second)
	err = chromedp.Run(*b.ctx,
		chromedp.WaitVisible(`//*[@id="react-root"]/div/div/div[2]/main/div/div/div/div/div/div[2]/div/section/div/div/div[1]/div/div/article/div/div/div/div[3]/div[5]/div[2]/div`, chromedp.BySearch),
		chromedp.Click(`//*[@id="react-root"]/div/div/div[2]/main/div/div/div/div/div/div[2]/div/section/div/div/div[1]/div/div/article/div/div/div/div[3]/div[5]/div[2]/div`, chromedp.BySearch),
		chromedp.Sleep(5*time.Second),
		chromedp.Click(`//*[@id="layers"]/div[2]/div/div/div/div[2]/div[3]/div/div/div/div`, chromedp.BySearch),
	)
	if err != nil {
		fmt.Println(err)
	}

	time.Sleep(5 * time.Second)
	return
}

//MaybeRetweet retweets the post with chance of 0.0-1.0.
func (b bot) MaybeRetweet(tweetURL string, chance float32) {
	if chance > rand.Float32() {
		b.EnsureRetweeted(tweetURL)
	}
	return
}
