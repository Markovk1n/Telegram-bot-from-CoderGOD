package main

//func TakeScreen(url string) {
//	ctx, cancel := chromedp.NewContext(context.Background())
//	defer cancel()
//
//	var imageBuf []byte
//	fileName := slug.Make(url) + ".png"
//
//	if err := chromedp.Run(ctx, ScreenshotTasks(url, &imageBuf)); err != nil {
//		log.Fatalln(err)
//	}
//
//	if err := ioutil.WriteFile(fileName, imageBuf, 0644); err != nil {
//		log.Fatalln(err)
//	}
//
//}
//
//func ScreenshotTasks(url string, imageBuf *[]byte) chromedp.Tasks {
//	return chromedp.Tasks{
//		chromedp.Navigate(url),
//		chromedp.FullScreenshot(imageBuf, 100),
//	}
//}
