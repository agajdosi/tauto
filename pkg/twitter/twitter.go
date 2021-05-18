package twitter

import (
	"context"
	"fmt"
	"time"

	"github.com/agajdosi/tauto/pkg/browser"
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/input"
	"github.com/chromedp/chromedp"
)

type Bot struct {
	Username string
	password string
	ctx      *context.Context
}

// NewUser creates a new instance of user struct
func NewUser(id int, username, password string, timeout int) (Bot, context.CancelFunc, error) {
	ctx, cancel := browser.CreateBrowser(username, timeout)
	b := Bot{username, password, ctx}
	fmt.Printf("Loging in bot: %v", b.Username)
	err := b.Login()
	if err != nil {
		fmt.Printf(" - login error: %v\n", err)
	} else {
		fmt.Println(" - OK")
	}

	return b, cancel, err
}

//Login logs user into the Twitter
func (b Bot) Login() error {
	logged, err := b.IsLoggedIn()
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
		chromedp.Sleep(5*time.Second),
		//password
		chromedp.WaitVisible(`//*[@id="react-root"]/div/div/div[2]/main/div/div/div[2]/form/div/div[2]/label/div/div[2]/div/input`, chromedp.BySearch),
		chromedp.SendKeys(`//*[@id="react-root"]/div/div/div[2]/main/div/div/div[2]/form/div/div[2]/label/div/div[2]/div/input`, b.password, chromedp.BySearch),
		chromedp.Sleep(5*time.Second),
		//login button
		chromedp.Click(`//*[@id="react-root"]/div/div/div[2]/main/div/div/div[2]/form/div/div[3]/div/div`, chromedp.BySearch),
		chromedp.Sleep(5*time.Second),
	)

	return nil
}

//IsProfileAccessible checks whether profile is not blocked
func (b Bot) IsProfileAccessible() (bool, error) {
	fmt.Printf("Checking accessibility of profile: %v", b.Username)
	var nodes []*cdp.Node
	err := chromedp.Run(*b.ctx,
		chromedp.Navigate("https://twitter.com"),
		chromedp.Sleep(10*time.Second),
		chromedp.Nodes(`//*[@id="phone_number"]`, &nodes, chromedp.AtLeast(0)),
	)
	if err != nil {
		fmt.Println(" - an error occured checking if phone is required.")
	}

	if len(nodes) > 0 {
		fmt.Println(" - profile not accessible.")
		return false, err
	}

	fmt.Println(" - OK")
	return true, nil
}

//IsLoggedIn checks whether the bot is logged in in the browser
func (b Bot) IsLoggedIn() (bool, error) {
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

//IsAvailable navigates to tweets URL and checks whether the tweet is available
func (b Bot) IsTweetAvailable(tweetURL string) (bool, error) {
	var nodes []*cdp.Node
	err := chromedp.Run(*b.ctx,
		chromedp.Navigate(tweetURL),
		chromedp.Sleep(4*time.Second),
		chromedp.Nodes(`//*[@href="https://help.twitter.com/rules-and-policies/notices-on-twitter"]`, &nodes, chromedp.AtLeast(0)),
	)
	if err != nil {
		fmt.Println("en error occured checking if tweet is available")
	}

	if len(nodes) > 0 {
		fmt.Println("Content not available.")
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
		input.InsertText(text),
		chromedp.Click(`//*[@id="layers"]/div[2]/div/div/div/div/div/div[2]/div[2]/div/div[3]/div/div/div/div[1]/div/div/div/div/div[2]/div[4]/div/div/div[2]/div[4]/div/span/span`, chromedp.BySearch),
		chromedp.Sleep(2*time.Second),
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
