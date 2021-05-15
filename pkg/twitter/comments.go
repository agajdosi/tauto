package twitter

import (
	"fmt"
	"strings"
	"time"

	"github.com/agajdosi/tauto/pkg/generate"
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/input"
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
		chromedp.Click(commentTextPath, chromedp.BySearch),
		input.InsertText(text),
		chromedp.Sleep(2*time.Second),
		chromedp.Click(commentSubmitPath, chromedp.BySearch),
	)
	if err != nil {
		fmt.Println(err)
	}

	time.Sleep(5 * time.Second)
	return
}

//Reply makes sure that the bot replied to the tweet with the text
func (b Bot) Reply(tweetURL, text string) {
	b.EnsureCommented(tweetURL, strings.TrimLeft(b.Username, "@"), text)
	time.Sleep(5 * time.Second)

	return
}

//ReplyFromTemplate makes sure that bot replied to the tweet with text generated from template
func (b Bot) ReplyFromTemplate(tweetURL, template string) {
	text := generate.FromTemplate(template)
	b.EnsureCommented(tweetURL, strings.TrimLeft(b.Username, "@"), text)
	time.Sleep(5 * time.Second)

	return
}

//TrollReply makes sure that the bot reacts to the tweet with some trolly comment
func (b Bot) TrollReply(tweetURL string) {
	text := generate.FromTemplateByName("stupidQuestions")
	b.EnsureCommented(tweetURL, strings.TrimLeft(b.Username, "@"), text)
	time.Sleep(5 * time.Second)

	return
}
