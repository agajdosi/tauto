package twitter

import (
	"context"
	"fmt"
	"time"

	"github.com/agajdosi/twitter-storm-toolkit/pkg/browser"
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
)

type Bot struct {
	Username string
	password string
	ctx      *context.Context
}

// NewUser creates a new instance of user struct
func NewUser(id int, username, password string, timeout int) (Bot, context.CancelFunc) {
	ctx, cancel := browser.CreateBrowser(username, timeout)

	return Bot{username, password, ctx}, cancel
}

//Login logs user into the Twitter
func (b Bot) Login() error {
	logged, err := b.isLoggedIn()
	if err != nil {
		return err
	}

	if logged == true {
		return nil
	}

	chromedp.Run(*b.ctx,
		chromedp.Navigate("https://twitter.com/login"),
		//username
		chromedp.WaitVisible(`//*[@id="react-root"]/div/div/div[2]/main/div/div/div[2]/form/div/div[1]/label/div/div[2]/div/input`, chromedp.BySearch),
		chromedp.SendKeys(`//*[@id="react-root"]/div/div/div[2]/main/div/div/div[2]/form/div/div[1]/label/div/div[2]/div/input`, b.Username, chromedp.BySearch),
		//password
		chromedp.WaitVisible(`//*[@id="react-root"]/div/div/div[2]/main/div/div/div[2]/form/div/div[2]/label/div/div[2]/div/input`, chromedp.BySearch),
		chromedp.SendKeys(`//*[@id="react-root"]/div/div/div[2]/main/div/div/div[2]/form/div/div[2]/label/div/div[2]/div/input`, b.password, chromedp.BySearch),
		//login button
		chromedp.Click(`//*[@id="react-root"]/div/div/div[2]/main/div/div/div[2]/form/div/div[3]/div/div`, chromedp.BySearch),
	)

	return nil
}

func (b Bot) isLoggedIn() (bool, error) {
	var nodes []*cdp.Node
	err := chromedp.Run(*b.ctx,
		chromedp.Navigate("https://twitter.com"),
		chromedp.Sleep(time.Second*2),
		chromedp.Nodes(`//*[@id="react-root"]/div/div/div[2]/header/div/div/div/div[1]/div[2]/nav/a[7]`, &nodes, chromedp.AtLeast(0), chromedp.BySearch),
	)

	if len(nodes) == 0 {
		return false, err
	}

	return true, err
}

//Post sends a new tweet
func (b Bot) Post(text string) error {
	err := b.Login()
	if err != nil {
		return err
	}

	err = chromedp.Run(*b.ctx,
		chromedp.Navigate("https://twitter.com"),
		chromedp.Sleep(time.Second*5),
		chromedp.WaitVisible(`//*[@id="react-root"]/div/div/div[2]/header/div/div/div/div[1]/div[3]/a/div`, chromedp.BySearch),
		chromedp.Click(`//*[@id="react-root"]/div/div/div[2]/header/div/div/div/div[1]/div[3]/a/div`, chromedp.BySearch),
		chromedp.Click(`//*[@id="layers"]/div[2]/div/div/div/div/div/div[2]/div[2]/div/div[3]/div/div/div/div[1]/div/div/div/div/div[2]/div[1]/div/div/div/div/div/div/div/div/div/div[1]/div/div/div/div[2]/div`, chromedp.BySearch),
		chromedp.KeyEvent(text),
		chromedp.Click(`//*[@id="layers"]/div[2]/div/div/div/div/div/div[2]/div[2]/div/div[3]/div/div/div/div[1]/div/div/div/div/div[2]/div[4]/div/div/div[2]/div[4]/div/span/span`, chromedp.BySearch),
		chromedp.Sleep(2*time.Second),
	)

	return err
}

//Follow will open and follow selected Twitter account.
func (b Bot) Follow(who string) error {
	fmt.Println("going to log in!")
	err := b.Login()
	if err != nil {
		return err
	}

	//twitter handles both "username" and "@username" formats in the URL so we do not care about it
	address := "https://twitter.com/" + who

	err = chromedp.Run(*b.ctx,
		chromedp.Navigate(address),
		chromedp.Sleep(time.Second*4),
		chromedp.WaitVisible(`//*[@id="react-root"]/div/div/div[2]/main/div/div/div/div[1]/div/div[2]/div/div/div[1]/div/div[1]/div/div[last()]/div/div`, chromedp.BySearch),
		chromedp.Click(`//*[@id="react-root"]/div/div/div[2]/main/div/div/div/div[1]/div/div[2]/div/div/div[1]/div/div[1]/div/div[last()]/div/div`, chromedp.BySearch),
		chromedp.Sleep(2*time.Second),
	)

	return err
}

func (b Bot) Reply(tweet, where string) error {
	err := b.Login()
	if err != nil {
		return nil
	}

	err = chromedp.Run(*b.ctx,
		chromedp.Navigate(where),
		chromedp.WaitVisible(`//*[@id="react-root"]/div/div/div[2]/main/div/div/div/div[1]/div/div[2]/div/section/div/div/div[1]/div/div/article/div/div/div/div[3]/div[5]/div[1]/div/div/div/div`, chromedp.BySearch),
		chromedp.Click(`//*[@id="react-root"]/div/div/div[2]/main/div/div/div/div[1]/div/div[2]/div/section/div/div/div[1]/div/div/article/div/div/div/div[3]/div[5]/div[1]/div/div/div/div`, chromedp.BySearch),
		chromedp.WaitVisible(`//*[@id="layers"]/div[2]/div/div/div/div/div/div[2]/div[2]/div/div[3]/div/div/div/div[2]/div/div/div/div/div[2]/div[1]/div/div/div/div/div/div/div/div/div/div[1]/div/div/div/div[2]/div/div/div/div`, chromedp.BySearch),
		chromedp.KeyEvent(tweet),
		chromedp.Click(`//*[@id="layers"]/div[2]/div/div/div/div/div/div[2]/div[2]/div/div[3]/div/div/div/div[2]/div/div/div/div/div[2]/div[2]/div/div/div[2]/div[2]/div/span/span`, chromedp.BySearch),
		chromedp.Sleep(time.Second*10000000),
	)

	return err
}

//Open opens the profile and leaves it open for manual tweaks
func (b Bot) Open() error {
	err := b.Login()
	if err != nil {
		return err
	}

	err = chromedp.Run(*b.ctx,
		chromedp.Navigate("https://twitter.com"),
		chromedp.Sleep(time.Second*1000000000),
	)

	return err
}
