package twitter

import (
	"fmt"
	"time"

	"github.com/chromedp/chromedp"
)

//IsLiked checks whether specified tweet is liked by the bot
func (b bot) IsLiked(tweetURL string) (bool, error) {
	var value string
	var ok bool

	err := chromedp.Run(*b.ctx,
		chromedp.Navigate(tweetURL),
		chromedp.WaitVisible(`//*[@id="react-root"]/div/div/div[2]/main/div/div/div/div[1]/div/div[2]/div/section/div/div/div[*]/div/div[1]/article/div/div/div/div[3]/div[*]/div[3]/div`, chromedp.BySearch),
		chromedp.AttributeValue(`//*[@id="react-root"]/div/div/div[2]/main/div/div/div/div[1]/div/div[2]/div/section/div/div/div[*]/div/div[1]/article/div/div/div/div[3]/div[*]/div[3]/div`, "aria-label", &value, &ok, chromedp.BySearch),
	)

	if value == "Liked" {
		return true, err
	}

	if value == "Like" {
		return false, err
	}

	return false, err
}

//EnsureLiked unsures that specified tweet is liked by the bot
func (b bot) EnsureLiked(tweetURL string) {
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
		chromedp.WaitVisible(`//*[@id="react-root"]/div/div/div[2]/main/div/div/div/div[1]/div/div[2]/div/section/div/div/div[*]/div/div[1]/article/div/div/div/div[3]/div[*]/div[3]/div`, chromedp.BySearch),
		chromedp.Click(`//*[@id="react-root"]/div/div/div[2]/main/div/div/div/div[1]/div/div[2]/div/section/div/div/div[*]/div/div[1]/article/div/div/div/div[3]/div[*]/div[3]/div`, chromedp.BySearch),
	)
	if err != nil {
		fmt.Println(err)
	}

	time.Sleep(5 * time.Second)
	return
}
