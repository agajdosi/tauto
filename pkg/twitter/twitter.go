package twitter

import (
	"context"
	"fmt"
	"time"

	"github.com/agajdosi/twitter-storm-toolkit/pkg/browser"
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
)

type user struct {
	username string
	password string
	ctx      *context.Context
}

// NewUser creates a new instance of user struct
func NewUser(username, password string) user {
	ctx, _ := browser.CreateBrowser(username)

	return user{username, password, ctx}
}

//Post sends a new tweet
func (u user) Post(text string) error {
	err := u.Login()
	if err != nil {
		return err
	}

	err = chromedp.Run(*u.ctx,
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

//Login logs user into the platform
func (u user) Login() error {
	logged, err := u.isLoggedIn()
	if err != nil {
		return err
	}

	if logged == true {
		return nil
	}

	fmt.Println("logging in!")
	chromedp.Run(*u.ctx,
		chromedp.Navigate("https://twitter.com"),
		chromedp.WaitVisible(`//*[@id="react-root"]/div/div/div/main/div/div/div/div[1]/div[1]/div/form/div/div[1]/div/label/div/div[2]/div/input`, chromedp.BySearch),
		chromedp.SendKeys(`//*[@id="react-root"]/div/div/div/main/div/div/div/div[1]/div[1]/div/form/div/div[1]/div/label/div/div[2]/div/input`, u.username, chromedp.BySearch),
		chromedp.SendKeys(`//*[@id="react-root"]/div/div/div/main/div/div/div/div[1]/div[1]/div/form/div/div[2]/div/label/div/div[2]/div/input`, u.password, chromedp.BySearch),
		chromedp.Click(`//*[@id="react-root"]/div/div/div/main/div/div/div/div[1]/div[1]/div/form/div/div[3]/div/div/span/span`, chromedp.BySearch),
	)

	return nil
}

func (u user) isLoggedIn() (bool, error) {
	var nodes []*cdp.Node
	err := chromedp.Run(*u.ctx,
		chromedp.Navigate("https://twitter.com"),
		chromedp.Sleep(time.Second*2),
		chromedp.Nodes(`/html/body/div/div/div/div[2]/header/div/div/div/div[1]/div[3]/a/div`, &nodes, chromedp.AtLeast(0), chromedp.BySearch),
	)

	if len(nodes) == 0 {
		return false, err
	}

	return true, err
}
