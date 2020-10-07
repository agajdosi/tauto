package register

import (
	"fmt"
	"strconv"
	"time"

	"github.com/agajdosi/twitter-storm-toolkit/pkg/identity"

	"github.com/agajdosi/twitter-storm-toolkit/pkg/browser"
	"github.com/chromedp/chromedp"
)

//Register registers a new user to twitter.
func Register() error {
	var email string
	name, surname := identity.GenerateName("M")
	birthtime := identity.GenerateBirthday()
	month := birthtime.Month().String()
	day := strconv.Itoa(birthtime.Day())
	year := strconv.Itoa(birthtime.Year())
	password := identity.GeneratePassword(12)

	ctx, _ := browser.CreateBrowser()
	ctx2, _ := browser.CreateBrowser()

	chromedp.Run(*ctx2,
		chromedp.Navigate("https://www.fakemail.net/"),
		chromedp.Sleep(time.Second*1),
		chromedp.WaitVisible(`#email`, chromedp.ByQuery),
		chromedp.Text(`#email`, &email, chromedp.ByQuery),
	)

	fmt.Println(name, surname, email, password)

	chromedp.Run(*ctx,
		chromedp.Navigate("https://twitter.com"),

		// Click to sign up
		chromedp.WaitVisible(`//*[@id="react-root"]/div/div/div/main/div/div/div/div[1]/div[2]/a[1]/div`, chromedp.BySearch),
		chromedp.Click(`//*[@id="react-root"]/div/div/div/main/div/div/div/div[1]/div[2]/a[1]/div`, chromedp.BySearch),

		// Screen 1 - input name, email and birth date
		chromedp.WaitVisible(`//*[@id="layers"]/div/div/div/div/div/div/div[2]/div[2]/div/div/div[2]/div[2]/div/div/div[2]/label/div/div[2]/div/input`, chromedp.BySearch),
		chromedp.SendKeys(`//*[@id="layers"]/div/div/div/div/div/div/div[2]/div[2]/div/div/div[2]/div[2]/div/div/div[2]/label/div/div[2]/div/input`, name+" "+surname, chromedp.BySearch),
		chromedp.Click(`//*[@id="layers"]/div/div/div/div/div/div/div[2]/div[2]/div/div/div[2]/div[2]/div/div/div[4]/span`, chromedp.BySearch),
		chromedp.SendKeys(`//*[@id="layers"]/div/div/div/div/div/div/div[2]/div[2]/div/div/div[2]/div[2]/div/div/div[3]/label/div/div[2]/div/input`, email, chromedp.BySearch),

		chromedp.SendKeys(`//*[@id="layers"]/div/div/div/div/div/div/div[2]/div[2]/div/div/div[2]/div[2]/div/div/div[5]/div[3]/div/div[1]/div[2]/select`, month, chromedp.BySearch),
		chromedp.SendKeys(`//*[@id="layers"]/div/div/div/div/div/div/div[2]/div[2]/div/div/div[2]/div[2]/div/div/div[5]/div[3]/div/div[2]/div[2]/select`, day, chromedp.BySearch),
		chromedp.SendKeys(`//*[@id="layers"]/div/div/div/div/div/div/div[2]/div[2]/div/div/div[2]/div[2]/div/div/div[5]/div[3]/div/div[3]/div[2]/select`, year, chromedp.BySearch),

		chromedp.Sleep(time.Second*1),
		chromedp.Click(`//*[@id="layers"]/div/div/div/div/div/div/div[2]/div[2]/div/div/div[2]/div[1]/div/div/div/div[3]/div/div/span/span`, chromedp.BySearch),

		// Screen 2 - allow marketing emails
		chromedp.WaitVisible(`//*[@id="layers"]/div/div/div/div/div/div/div[2]/div[2]/div/div/div[2]/div[2]/div/div/label[1]/div[2]/input`, chromedp.BySearch),
		chromedp.Click(`//*[@id="layers"]/div/div/div/div/div/div/div[2]/div[2]/div/div/div[2]/div[2]/div/div/label[1]/div[2]/input`, chromedp.BySearch),
		chromedp.Click(`//*[@id="layers"]/div/div/div/div/div/div/div[2]/div[2]/div/div/div[2]/div[2]/div/div/label[2]/div[2]/input`, chromedp.BySearch),
		chromedp.Click(`//*[@id="layers"]/div/div/div/div/div/div/div[2]/div[2]/div/div/div[2]/div[2]/div/div/label[3]/div[2]/input`, chromedp.BySearch),
		chromedp.Click(`//*[@id="layers"]/div/div/div/div/div/div/div[2]/div[2]/div/div/div[2]/div[1]/div/div/div/div[3]/div/div/span/span`, chromedp.BySearch),

		// Screen 3 -
		chromedp.WaitVisible(`//*[@id="layers"]/div/div/div/div/div/div/div[2]/div[2]/div/div/div[2]/div[2]/div/div/div/div[5]/div`, chromedp.BySearch),
		chromedp.Click(`//*[@id="layers"]/div/div/div/div/div/div/div[2]/div[2]/div/div/div[2]/div[2]/div/div/div/div[5]/div`, chromedp.BySearch),

		// Screen 4 - paste vericode manually

		// Screen 5 - inserts password
		chromedp.WaitVisible(`//*[@id="layers"]/div/div/div/div/div/div/div[2]/div[2]/div/div/div[2]/div[2]/div/div/div[2]/div/label/div/div[2]/div/input`, chromedp.BySearch),
		chromedp.SendKeys(`//*[@id="layers"]/div/div/div/div/div/div/div[2]/div[2]/div/div/div[2]/div[2]/div/div/div[2]/div/label/div/div[2]/div/input`, password, chromedp.BySearch),
		chromedp.Click(`//*[@id="layers"]/div/div/div/div/div/div/div[2]/div[2]/div/div/div[2]/div[1]/div/div/div/div[3]/div/div/span/span`, chromedp.BySearch),

		//skip profile picture
		chromedp.WaitVisible(`//*[@id="layers"]/div[2]/div/div/div/div/div/div[2]/div[2]/div/div/div[2]/div[1]/div/div/div/div[3]/div/div/span/span`, chromedp.BySearch),
		chromedp.Click(`//*[@id="layers"]/div[2]/div/div/div/div/div/div[2]/div[2]/div/div/div[2]/div[1]/div/div/div/div[3]/div/div/span/span`, chromedp.BySearch),

		//skip bio
		chromedp.WaitVisible(`//*[@id="layers"]/div[2]/div/div/div/div/div/div[2]/div[2]/div/div/div[2]/div[1]/div/div/div/div[3]/div/div/span/span`, chromedp.BySearch),
		chromedp.Click(`//*[@id="layers"]/div[2]/div/div/div/div/div/div[2]/div[2]/div/div/div[2]/div[1]/div/div/div/div[3]/div/div/span/span`, chromedp.BySearch),

		//skip interests
		chromedp.WaitVisible(`//*[@id="layers"]/div[2]/div/div/div/div/div/div[2]/div[2]/div/div/div[2]/div[1]/div/div/div/div[3]/div/div/span/span`, chromedp.BySearch),
		chromedp.Click(`//*[@id="layers"]/div[2]/div/div/div/div/div/div[2]/div[2]/div/div/div[2]/div[1]/div/div/div/div[3]/div/div/span/span`, chromedp.BySearch),

		//skip follow suggestions
		chromedp.WaitVisible(`//*[@id="layers"]/div[2]/div/div/div/div/div/div[2]/div[2]/div/div/div[2]/div[1]/div/div/div/div[3]/div/div/span/span`, chromedp.BySearch),
		chromedp.Click(`//*[@id="layers"]/div[2]/div/div/div/div/div/div[2]/div[2]/div/div/div[2]/div[1]/div/div/div/div[3]/div/div/span/span`, chromedp.BySearch),

		//skip notifications
		chromedp.WaitVisible(`//*[@id="layers"]/div[2]/div/div/div/div/div/div[2]/div[2]/div/div/div[2]/div[2]/div/div/div[2]/div/div[2]/div[2]/div`, chromedp.BySearch),
		chromedp.Click(`//*[@id="layers"]/div[2]/div/div/div/div/div/div[2]/div[2]/div/div/div[2]/div[2]/div/div/div[2]/div/div[2]/div[2]/div`, chromedp.BySearch),

		chromedp.Sleep(time.Second*2000),
	)

	chromedp.Cancel(*ctx)
	chromedp.Cancel(*ctx2)

	return nil
}
