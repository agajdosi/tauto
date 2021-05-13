package twitter

import (
	"fmt"
	"strings"
	"time"

	"github.com/agajdosi/tauto/pkg/generate"
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
)

const commentAuthorPath = `//*[@href="/%v" and not(@aria-label="Profile")]`
const replyPath = `//*[@aria-label="Reply"]`
const commentTextPath = `//*[@aria-label="Tweet text"]`
const commentSubmitPath = `//*[@data-testid="tweetButton"]`

//IsCommented checks whether specified tweet was already commented by the bot
func (b Bot) IsCommented(tweetURL, nick string) (bool, error) {
	var replies []*cdp.Node

	err := chromedp.Run(*b.ctx,
		chromedp.Navigate(tweetURL),
		chromedp.WaitVisible(replyPath, chromedp.BySearch),
		chromedp.Sleep(2*time.Second),
		chromedp.Nodes(fmt.Sprintf(commentAuthorPath, nick), &replies, chromedp.AtLeast(0)),
	)

	fmt.Println(len(replies))

	if len(replies) == 0 {
		return false, err
	}

	return true, err
}

//EnsureCommented ensures that specified tweet is commented by the bot
func (b Bot) EnsureCommented(tweetURL, nick, text string) {
	commented, err := b.IsCommented(tweetURL, nick)
	if err != nil {
		fmt.Println(err)
	}
	time.Sleep(5 * time.Second)
	if commented == true {
		return
	}

	time.Sleep(5 * time.Second)
	err = chromedp.Run(*b.ctx,
		chromedp.WaitVisible(replyPath, chromedp.BySearch),
		chromedp.Click(replyPath, chromedp.BySearch),
		chromedp.Sleep(5*time.Second),
		chromedp.SendKeys(commentTextPath, text, chromedp.BySearch),
		chromedp.Sleep(2*time.Second),
		chromedp.Click(commentSubmitPath, chromedp.BySearch),
	)
	if err != nil {
		fmt.Println(err)
	}

	time.Sleep(15 * time.Second)
	return
}

//TrollComment makes sure that the bot reacts to the tweet with some trolly comment
func (b Bot) TrollComment(tweetURL, nick string) {
	text := generate.StupidQuestion()
	b.EnsureCommented(tweetURL, strings.TrimLeft(nick, "@"), text)
	return
}
