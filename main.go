package main

import (
	"context"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
	"github.com/gosimple/slug"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	url := os.Args[1]
	var filename string

	if len(os.Args) == 3 {
		filename = os.Args[2]
	} else {
		filename = slug.Make(url) + ".png"
	}

	var imageBuf []byte

	if err := chromedp.Run(
		ctx,
		ScreenshotTasks(url, &imageBuf),
	); err != nil {
		log.Fatal(err)

	}

	if err := ioutil.WriteFile(filename, imageBuf, 0644); err != nil {
		log.Fatal(err)
	}
}

func ScreenshotTasks(url string, imageBuf *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(url),
		chromedp.ActionFunc(func(ctx context.Context) (err error) {
			*imageBuf, err = page.CaptureScreenshot().Do(ctx)
			return err
		}),
	}
}
