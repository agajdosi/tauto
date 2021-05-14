package twitter

import (
	"fmt"
	"time"

	"github.com/chromedp/chromedp"
)

const followButton = `//*[@data-testid="%v-unfollow" or @data-testid="%v-follow"]`

//Follow will open and follow selected Twitter account.
func (b Bot) Follow(who string) error {
	//twitter handles both "username" and "@username" formats in the URL so we do not care about it
	address := "https://twitter.com/" + who

	err := chromedp.Run(*b.ctx,
		chromedp.Navigate(address),
		chromedp.Sleep(time.Second*4),
		chromedp.WaitVisible(`//*[@id="react-root"]/div/div/div[2]/main/div/div/div/div[1]/div/div[2]/div/div/div[1]/div/div[1]/div/div[last()]/div/div`, chromedp.BySearch),
		chromedp.Click(`//*[@id="react-root"]/div/div/div[2]/main/div/div/div/div[1]/div/div[2]/div/div/div[1]/div/div[1]/div/div[last()]/div/div`, chromedp.BySearch),
		chromedp.Sleep(5*time.Second),
	)

	return err
}

func (b Bot) IsFollowed(who string) (bool, error) {
	var isFollowing string
	var ok bool
	ID, _ := GetUserID(who)
	specificFollowButton := fmt.Sprintf(followButton, ID, ID)

	url := "https://twitter.com/" + who
	err := chromedp.Run(*b.ctx,
		chromedp.Navigate(url),
		chromedp.Sleep(5*time.Second),
		chromedp.WaitVisible(specificFollowButton, chromedp.BySearch),
		chromedp.AttributeValue(specificFollowButton, "data-testid", &isFollowing, &ok, chromedp.BySearch),
	)

	if isFollowing == ID+"-unfollow" {
		return true, err
	} else if isFollowing == ID+"-follow" {
		return false, err
	}

	fmt.Println("Could not detect follow data.")
	return false, err
}

func (b Bot) EnsureFollowed(who string) error {
	isFollowing, err := b.IsFollowed(who)
	if err != nil {
		return err
	}

	if isFollowing == false {
		err = b.Follow(who)
	}

	return err
}
