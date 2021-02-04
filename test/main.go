package main

import (
	"context"
	"fmt"
	"time"

	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
)

func main() {

	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.DisableGPU,
		chromedp.Flag("disable-extensions", false),
		chromedp.Flag("headless", false),
		chromedp.Flag("disable-session-crashed-bubble", true),
	)

	ctx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	ctx, cancel = chromedp.NewContext(ctx)
	ctx, cancel = context.WithTimeout(ctx, 300*time.Second)

	defer cancel()

	chromedp.ListenTarget(ctx, func(ev interface{}) {

		switch ev := ev.(type) {
		case *page.EventNavigatedWithinDocument:
			fmt.Printf("* console.%s call:\n", ev)
		}
	})

	if err := chromedp.Run(ctx,
		chromedp.Navigate("https://gajdosik.org"),
		chromedp.Click("#alert", chromedp.ByID),
	); err != nil {
		panic(err)
	}

}
