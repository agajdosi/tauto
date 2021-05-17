package browser

import (
	"context"
	"fmt"
	"path/filepath"
	"time"

	"github.com/agajdosi/tauto/pkg/database"
	"github.com/chromedp/chromedp"
)

//CreateBrowser creates a new instance of browser - opens a new window.
func CreateBrowser(username string, timeout int) (*context.Context, context.CancelFunc) {
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.DisableGPU,
		chromedp.Flag("disable-extensions", false),
		chromedp.Flag("headless", false),
		chromedp.Flag("disable-session-crashed-bubble", true),
		chromedp.Flag("disable-infobars", true),
		//chromedp.Flag("kiosk", true),
		chromedp.UserAgent("Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.212 Safari/537.36"),
	)

	configDir, err := BrowserProfilesPath()
	if err != nil {
		fmt.Println("Error finding the profiles directory.")
	}

	if username != "" {
		userDataDir := filepath.Join(configDir, username)
		opts = append(opts, chromedp.UserDataDir(userDataDir))
	}

	ctx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	ctx, cancel = chromedp.NewContext(ctx)
	ctx, cancel = context.WithTimeout(ctx, time.Duration(timeout)*time.Second)

	return &ctx, cancel
}

//Location returns a location where profiles folders are located.
func BrowserProfilesPath() (string, error) {
	configDir := database.ConfigDirectory()
	location := filepath.Join(configDir, "profiles")

	return location, nil
}
