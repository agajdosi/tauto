package twitter

import (
	"context"
	"fmt"
	"time"

	"github.com/agajdosi/twitter-storm-toolkit/pkg/browser"
	"github.com/chromedp/chromedp"
)

type user struct {
	username string
	password string
	ctx      *context.Context
}

// NewUser creates a new instance of user struct
func NewUser(username, password string) user {
	ctx, _ := browser.CreateBrowser()

	return user{username, password, ctx}
}

//Post sends a new tweet
func (u user) Post(text string) error {
	u.Login()

	fmt.Println("login done")
	err := chromedp.Run(*u.ctx,
		chromedp.WaitVisible(`//*[@id="react-root"]/div/div/div[2]/header/div/div/div/div[1]/div[3]/a/div`, chromedp.BySearch),
		chromedp.Click(`//*[@id="react-root"]/div/div/div[2]/header/div/div/div/div[1]/div[3]/a/div`, chromedp.BySearch),
		chromedp.Sleep(3000*time.Second),
		chromedp.WaitVisible(`//*[@id="layers"]/div[2]/div/div/div/div/div/div[2]/div[2]/div/div[3]/div/div/div/div[1]/div/div/div/div/div[2]/div[1]/div/div/div/div/div/div/div/div/div/div[1]/div/div/div/div[2]/div/div/div/div`, chromedp.BySearch),
		chromedp.SendKeys(`//*[@id="layers"]/div[2]/div/div/div/div/div/div[2]/div[2]/div/div[3]/div/div/div/div[1]/div/div/div/div/div[2]/div[1]/div/div/div/div/div/div/div/div/div/div[1]/div/div/div/div[2]/div/div/div/div`, text, chromedp.BySearch),
		chromedp.Sleep(10*time.Second),
		chromedp.Click(`//*[@id="react-root"]/div/div/div[2]/main/div/div/div/div[1]/div/div[2]/div/div[2]/div[1]/div/div/div/div[2]/div[4]/div/div/div[2]/div[3]/div/span/span`, chromedp.BySearch),
		chromedp.Sleep(2000*time.Second),
	)

	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println("posting finished")

	return nil
}

//Login logs user into the platform
func (u user) Login() error {
	chromedp.Run(*u.ctx,
		chromedp.Navigate("https://twitter.com"),
		chromedp.WaitVisible(`//*[@id="react-root"]/div/div/div/main/div/div/div/div[1]/div[1]/div/form/div/div[1]/div/label/div/div[2]/div/input`, chromedp.BySearch),
		chromedp.SendKeys(`//*[@id="react-root"]/div/div/div/main/div/div/div/div[1]/div[1]/div/form/div/div[1]/div/label/div/div[2]/div/input`, u.username, chromedp.BySearch),
		chromedp.SendKeys(`//*[@id="react-root"]/div/div/div/main/div/div/div/div[1]/div[1]/div/form/div/div[2]/div/label/div/div[2]/div/input`, u.password, chromedp.BySearch),
		chromedp.Click(`//*[@id="react-root"]/div/div/div/main/div/div/div/div[1]/div[1]/div/form/div/div[3]/div/div/span/span`, chromedp.BySearch),
	)

	return nil
}
