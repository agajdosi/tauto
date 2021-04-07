package browser

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/chromedp/chromedp"
)

//CreateBrowser creates a new instance of browser - opens a new window.
func CreateBrowser(username string) (*context.Context, *context.CancelFunc) {
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.DisableGPU,
		chromedp.Flag("disable-extensions", false),
		chromedp.Flag("headless", false),
		chromedp.Flag("disable-session-crashed-bubble", true),
	)

	configDir, err := Location()
	if err != nil {
		fmt.Println("Error finding the profiles directory.")
	}

	if username != "" {
		userDataDir := filepath.Join(configDir, username)
		opts = append(opts, chromedp.UserDataDir(userDataDir))
	}

	ctx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	ctx, cancel = chromedp.NewContext(ctx)
	ctx, cancel = context.WithTimeout(ctx, 300*time.Second)

	return &ctx, &cancel
}

//Location returns a location where profiles folders are located.
func Location() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	configDir := filepath.Join(home, ".tst")
	err = os.MkdirAll(configDir, 0700)
	if err != nil {
		return "", err
	}

	location := filepath.Join(configDir, "profiles")
	return location, nil
}
