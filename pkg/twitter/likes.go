package twitter

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/chromedp/chromedp"
)

const likePath = `//*[@id="react-root"]/div/div/div[2]/main/div/div/div/div[1]/div/div[2]/div/section/div/div/div[*]/div/div[1]/article/div/div/div/div[3]/div[*]/div[3]/div`

//IsLiked checks whether specified tweet is liked by the bot
func (b Bot) IsLiked(tweetURL string) (bool, error) {
	var value string
	var ok bool

	err := chromedp.Run(*b.ctx,
		chromedp.Navigate(tweetURL),
		chromedp.WaitVisible(likePath, chromedp.BySearch),
		chromedp.AttributeValue(likePath, "aria-label", &value, &ok, chromedp.BySearch),
	)

	if value == "Liked" {
		return true, err
	}

	if value == "Like" {
		return false, err
	}

	fmt.Println("like value not found")
	return false, err
}

//EnsureLiked unsures that specified tweet is liked by the bot
func (b Bot) EnsureLiked(tweetURL string) {
	liked, err := b.IsLiked(tweetURL)
	if err != nil {
		fmt.Println(err)
	}
	time.Sleep(10 * time.Second)
	if liked == true {
		return
	}

	time.Sleep(5 * time.Second)
	err = chromedp.Run(*b.ctx,
		chromedp.WaitVisible(likePath, chromedp.BySearch),
		chromedp.Click(likePath, chromedp.BySearch),
	)
	if err != nil {
		fmt.Println(err)
	}

	time.Sleep(5 * time.Second)
	return
}

//MaybeLike likes the post with chance of 0.0-1.0.
func (b Bot) MaybeLike(tweetURL string, chance float32) {
	if chance > rand.Float32() {
		b.EnsureLiked(tweetURL)
	}
	return
}
