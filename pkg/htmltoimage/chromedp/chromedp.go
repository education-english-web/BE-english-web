package chromedp

import (
	"context"
	"fmt"
	"io"
	"sync/atomic"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/emulation"
	"github.com/chromedp/cdproto/page"
	chromedpLib "github.com/chromedp/chromedp"

	"github.com/education-english-web/BE-english-web/pkg/htmltoimage"
)

type chromedp struct{}

func New() htmltoimage.Generator {
	return &chromedp{}
}

func (c *chromedp) GenerateByID(html, elementID string, out io.Writer) error {
	opts := append(chromedpLib.DefaultExecAllocatorOptions[:],
		chromedpLib.DisableGPU,
		chromedpLib.NoSandbox,
	)

	ctx, cancel := chromedpLib.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	ctx, cancel = chromedpLib.NewContext(ctx)
	defer cancel()

	if err := chromedpLib.Run(
		ctx,
		chromedpLib.Navigate("about:blank"),
		actionLoadHTMLContent(html),
		chromedpLib.ActionFunc(func(ctx context.Context) error {
			var buf []byte

			// transparent image
			if err := emulation.SetDefaultBackgroundColorOverride().WithColor(&cdp.RGBA{}).Do(ctx); err != nil {
				return fmt.Errorf("chromedp set transparent background: %w", err)
			}

			// scale image to has better quality
			if err := emulation.SetDeviceMetricsOverride(800, 800, 2, false).Do(ctx); err != nil {
				return fmt.Errorf("chromedp set device metrics: %w", err)
			}

			if err := chromedpLib.Screenshot(elementID, &buf, chromedpLib.NodeVisible).Do(ctx); err != nil {
				return fmt.Errorf("chromedp run action: %w", err)
			}

			if _, err := out.Write(buf); err != nil {
				return fmt.Errorf("write out result: %w", err)
			}

			return nil
		}),
	); err != nil {
		return fmt.Errorf("generate image from html: %w", err)
	}

	return nil
}

// actionLoadHTMLContent load static html content into the page
func actionLoadHTMLContent(html string) chromedpLib.ActionFunc {
	return func(ctx context.Context) error {
		var isLoaded atomic.Bool

		// create a new cancel context that will be canceled
		// when `page.EventLoadEventFired` is received
		listenerCtx, cancel := context.WithCancel(ctx)

		defer cancel()

		chromedpLib.ListenTarget(listenerCtx, func(ev interface{}) {
			if _, ok := ev.(*page.EventLoadEventFired); ok {
				_ = isLoaded.CompareAndSwap(false, true)
				// stop listener
				cancel()
			}
		})

		frameTree, err := page.GetFrameTree().Do(ctx)
		if err != nil {
			return err
		}

		// page.SetDocumentContent will trigger the page.EventLoadEventFired event
		if err := page.SetDocumentContent(frameTree.Frame.ID, html).Do(ctx); err != nil {
			return err
		}

		// wait until the page loaded (page.EventLoadEventFired event be handled)
		// or be canceled by parent context
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-listenerCtx.Done():
			if isLoaded.Load() {
				return nil
			}

			return listenerCtx.Err()
		}
	}
}
