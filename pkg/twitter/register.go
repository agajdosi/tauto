package twitter

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/agajdosi/twitter-storm-toolkit/pkg/database"
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

	var username string
	var getUsernameOK bool

	ctx, _ := browser.CreateBrowser("new-user")
	ctx2, _ := browser.CreateBrowser("")

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
		chromedp.WaitVisible(`//*[@id="react-root"]/div/div/div/main/div/div/div/div[1]/div/div[3]/a[1]`, chromedp.BySearch),
		chromedp.Click(`//*[@id="react-root"]/div/div/div/main/div/div/div/div[1]/div/div[3]/a[1]`, chromedp.BySearch),

		// Screen 1 - input name, email and birth date
		chromedp.WaitVisible(`//*[@id="layers"]/div/div/div/div/div/div/div[2]/div[2]/div/div/div[2]/div[2]/div/div/div[2]/label/div/div[2]/div/input`, chromedp.BySearch),
		chromedp.SendKeys(`//*[@id="layers"]/div/div/div/div/div/div/div[2]/div[2]/div/div/div[2]/div[2]/div/div/div[2]/label/div/div[2]/div/input`, name+" "+surname, chromedp.BySearch),
		chromedp.Click(`//*[@id="layers"]/div/div/div/div/div/div/div[2]/div[2]/div/div/div[2]/div[2]/div/div/div[4]/span`, chromedp.BySearch),
		chromedp.SendKeys(`//*[@id="layers"]/div/div/div/div/div/div/div[2]/div[2]/div/div/div[2]/div[2]/div/div/div[3]/label/div/div[2]/div/input`, email, chromedp.BySearch),

		chromedp.SendKeys(`//*[@id="Month"]`, month, chromedp.BySearch),
		chromedp.SendKeys(`//*[@id="Day"]`, day, chromedp.BySearch),
		chromedp.SendKeys(`//*[@id="Year"]`, year, chromedp.BySearch),

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
		chromedp.Sleep(time.Second*2),
		chromedp.Click(`//*[@id="layers"]/div/div/div/div/div/div/div[2]/div[2]/div/div/div[2]/div[1]/div/div/div/div[3]/div`, chromedp.BySearch),

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
		chromedp.WaitVisible(`//*[@id="layers"]/div[2]/div/div/div/div/div/div[2]/div[2]/div/div/div[2]/div[2]/div/div/div/div/div[2]/div[2]/div[2]`, chromedp.BySearch),
		chromedp.Click(`//*[@id="layers"]/div[2]/div/div/div/div/div/div[2]/div[2]/div/div/div[2]/div[2]/div/div/div/div/div[2]/div[2]/div[2]`, chromedp.BySearch),

		//get username
		chromedp.WaitVisible(`//*[@id="react-root"]/div/div/div[2]/header/div/div/div/div[1]/div[2]/nav/a[7]`, chromedp.BySearch),
		chromedp.AttributeValue(`//*[@id="react-root"]/div/div/div[2]/header/div/div/div/div[1]/div[2]/nav/a[7]`, "href", &username, &getUsernameOK, chromedp.BySearch),
	)

	username = strings.TrimLeft(username, "/")
	fmt.Printf("- registration OK, adding user %v into database... ", username)

	id, err := database.AddBot(username, password)
	if err != nil {
		return err
	}

	fmt.Println("user successfuly added: ")
	fmt.Printf("- %v %v %v %v %v %v\n\n", id, username, password, name, surname, email)

	time.Sleep(2000 * time.Second)

	chromedp.Cancel(*ctx)
	chromedp.Cancel(*ctx2)

	return nil
}
