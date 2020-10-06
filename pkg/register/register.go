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

		chromedp.WaitVisible(`//*[@id="react-root"]/div/div/div/main/div/div/div/div[1]/div[2]/a[1]/div`, chromedp.BySearch),
		chromedp.Click(`//*[@id="react-root"]/div/div/div/main/div/div/div/div[1]/div[2]/a[1]/div`, chromedp.BySearch),

		chromedp.WaitVisible(`//*[@id="layers"]/div/div/div/div/div/div/div[2]/div[2]/div/div/div[2]/div[2]/div/div/div[2]/label/div/div[2]/div/input`, chromedp.BySearch),
		chromedp.SendKeys(`//*[@id="layers"]/div/div/div/div/div/div/div[2]/div[2]/div/div/div[2]/div[2]/div/div/div[2]/label/div/div[2]/div/input`, name+" "+surname, chromedp.BySearch),
		chromedp.Click(`//*[@id="layers"]/div/div/div/div/div/div/div[2]/div[2]/div/div/div[2]/div[2]/div/div/div[4]/span`, chromedp.BySearch),
		chromedp.SendKeys(`//*[@id="layers"]/div/div/div/div/div/div/div[2]/div[2]/div/div/div[2]/div[2]/div/div/div[3]/label/div/div[2]/div/input`, email, chromedp.BySearch),

		chromedp.SendKeys(`//*[@id="layers"]/div/div/div/div/div/div/div[2]/div[2]/div/div/div[2]/div[2]/div/div/div[5]/div[3]/div/div[1]/div[2]/select`, month, chromedp.BySearch),
		chromedp.SendKeys(`//*[@id="layers"]/div/div/div/div/div/div/div[2]/div[2]/div/div/div[2]/div[2]/div/div/div[5]/div[3]/div/div[2]/div[2]/select`, day, chromedp.BySearch),
		chromedp.SendKeys(`//*[@id="layers"]/div/div/div/div/div/div/div[2]/div[2]/div/div/div[2]/div[2]/div/div/div[5]/div[3]/div/div[3]/div[2]/select`, year, chromedp.BySearch),

		chromedp.Sleep(time.Second*1),
		chromedp.Click(`//*[@id="layers"]/div/div/div/div/div/div/div[2]/div[2]/div/div/div[2]/div[1]/div/div/div/div[3]/div/div/span/span`, chromedp.BySearch),

		// Screen 2
		chromedp.WaitVisible(`//*[@id="layers"]/div/div/div/div/div/div/div[2]/div[2]/div/div/div[2]/div[2]/div/div/label[1]/div[2]/input`, chromedp.BySearch),
		chromedp.Click(`//*[@id="layers"]/div/div/div/div/div/div/div[2]/div[2]/div/div/div[2]/div[2]/div/div/label[1]/div[2]/input`, chromedp.BySearch),
		chromedp.Click(`//*[@id="layers"]/div/div/div/div/div/div/div[2]/div[2]/div/div/div[2]/div[2]/div/div/label[2]/div[2]/input`, chromedp.BySearch),
		chromedp.Click(`//*[@id="layers"]/div/div/div/div/div/div/div[2]/div[2]/div/div/div[2]/div[2]/div/div/label[3]/div[2]/input`, chromedp.BySearch),
		chromedp.Click(`//*[@id="layers"]/div/div/div/div/div/div/div[2]/div[2]/div/div/div[2]/div[1]/div/div/div/div[3]/div/div/span/span`, chromedp.BySearch),

		// Screen 3
		chromedp.WaitVisible(`//*[@id="layers"]/div/div/div/div/div/div/div[2]/div[2]/div/div/div[2]/div[2]/div/div/div/div[5]/div`, chromedp.BySearch),
		chromedp.Click(`//*[@id="layers"]/div/div/div/div/div/div/div[2]/div[2]/div/div/div[2]/div[2]/div/div/div/div[5]/div`, chromedp.BySearch),

		// Screen 4 - paste vericode manually

		// Screen 5 - inserts password
		chromedp.WaitVisible(`//*[@id="layers"]/div/div/div/div/div/div/div[2]/div[2]/div/div/div[2]/div[2]/div/div/div[2]/div/label/div/div[2]/div/input`, chromedp.BySearch),
		chromedp.SendKeys(`//*[@id="layers"]/div/div/div/div/div/div/div[2]/div[2]/div/div/div[2]/div[2]/div/div/div[2]/div/label/div/div[2]/div/input`, password, chromedp.BySearch),
		chromedp.Click(`//*[@id="layers"]/div/div/div/div/div/div/div[2]/div[2]/div/div/div[2]/div[1]/div/div/div/div[3]/div/div/span/span`, chromedp.BySearch),

		//skip profile picture
		//*[@id="layers"]/div[2]/div/div/div/div/div/div[2]/div[2]/div/div/div[2]/div[1]/div/div/div/div[3]/div/div/span/span

		//skip bio
		//*[@id="layers"]/div[2]/div/div/div/div/div/div[2]/div[2]/div/div/div[2]/div[1]/div/div/div/div[3]/div/div/span/span

		//skip interests
		//*[@id="layers"]/div[2]/div/div/div/div/div/div[2]/div[2]/div/div/div[2]/div[1]/div/div/div/div[3]/div/div/span/span

		//skip follow suggestions
		//*[@id="layers"]/div[2]/div/div/div/div/div/div[2]/div[2]/div/div/div[2]/div[1]/div/div/div/div[3]/div/div/span/span

		//skip notifications
		//*[@id="layers"]/div[2]/div/div/div/div/div/div[2]/div[2]/div/div/div[2]/div[2]/div/div/div[2]/div/div[2]/div[2]/div

		

		chromedp.Sleep(time.Second*2000),
	)

	//#react-root > div > div > div > main > div > div > div > div:nth-child(1) > div.css-1dbjc4n.r-1kihuf0.r-17w48nw > a.css-4rbku5.css-18t94o4.css-1dbjc4n.r-urgr8i.r-42olwf.r-sdzlij.r-1phboty.r-rs99b7.r-1loqt21.r-1w2pmg.r-1ifxtd0.r-5prp13.r-1hw0jha.r-1ny4l3l.r-1fneopy.r-o7ynqc.r-6416eg.r-lrvibr > div
	//*[@id="react-root"]/div/div/div/main/div/div/div/div[1]/div[2]/a[1]/div

	chromedp.Cancel(*ctx)
	chromedp.Cancel(*ctx2)

	return nil
}
