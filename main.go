package main

import (
	"context"
	"github.com/chromedp/chromedp"
	"github.com/gosimple/slug"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	TakeScreen(os.Args[1])

}

//func ScreenshotTasks1(url string, imageBuf *[]byte) chromedp.Tasks {
//	return chromedp.Tasks{
//
//		chromedp.Navigate(url),
//		chromedp.Sleep(2 * time.Second),
//		chromedp.FullScreenshot(imageBuf, 100),
//
//		//chromedp.Navigate(url),
//		//chromedp.ActionFunc(func(ctx context.Context) (err error) {
//		//	*imageBuf, err = page.CaptureScreenshot().Do(ctx)
//		//	return err
//
//	}
//}

func TakeScreen(url string) {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	var imageBuf []byte
	fileName := slug.Make(url) + ".png"

	if err := chromedp.Run(ctx, ScreenshotTasks(url, &imageBuf)); err != nil {
		log.Fatalln(err)
	}

	if err := ioutil.WriteFile(fileName, imageBuf, 0644); err != nil {
		log.Fatalln(err)
	}

}

func ScreenshotTasks(url string, imageBuf *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(url),
		chromedp.FullScreenshot(imageBuf, 100),
	}
}
