package main

import (
	"context"

	"github.com/chromedp/chromedp"
)

func main() {
	ctx, cancel := chromedp.NewContext()
	defer cancel()
	chromedp.Run(ctx, chromedp.ActionFunc(func(ctx context.Context) error {
		_, err := domain.SomeAction().Do(ctx)
		return err
	}))
}
