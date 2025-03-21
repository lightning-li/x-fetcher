package main

import (
	// "encoding/json"
	"fmt"
	"context"
	"log"
	"github.com/chromedp/chromedp"
	"time"
	// "strconv"
	"strings"
	"math/rand"
)

func FetchXPostByUri(uri string) (string, error) {
	opts := []chromedp.ExecAllocatorOption{
        chromedp.Flag("headless", true), // Set to true for headless mode in production
        chromedp.Flag("disable-gpu", true),
        chromedp.Flag("no-sandbox", true),
		chromedp.Flag("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/123.0.0.0 Safari/537.36"),
		chromedp.Flag("window-size", "1920,1080"),
		chromedp.Flag("ignore-certificate-errors", true),
		chromedp.Flag("disable-blink-features", "AutomationControlled"),
		chromedp.Flag("disable-web-security", true),
    }
    allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()
	ctx, cancel := chromedp.NewContext(allocCtx, chromedp.WithDebugf(log.Printf))
	// ctx, cancel := chromedp.NewContext(allocCtx)
	defer cancel()

	// Set a timeout
	ctx, cancel = context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	// Variable to store the tweet text
	var tweetText string

	// Run Chrome tasks
	err := chromedp.Run(ctx,
		chromedp.Navigate(uri),
		chromedp.WaitVisible(`div[data-testid='tweetText']`, chromedp.ByQuery),
		chromedp.WaitReady(`div[data-testid='tweetText']`, chromedp.ByQuery),
		chromedp.Sleep(time.Duration(rand.Intn(1000)+1000)*time.Millisecond),
		chromedp.Text(`div[data-testid='tweetText']`, &tweetText, chromedp.ByQuery),
	)
	if err != nil {
		return "", fmt.Errorf("could not fetch tweet content: %v", err)
	}

	return strings.TrimSpace(tweetText), nil	
}

func main() {
	// Fetch the tweet
	tweet, err := FetchXPostByUri("https://x.com/elonmusk/status/1902929242620416387")
	if err != nil {
		log.Fatalf("could not fetch tweet: %v", err)
	}

	fmt.Println(tweet)
}