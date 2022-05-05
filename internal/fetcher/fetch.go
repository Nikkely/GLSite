package fetcher

import (
	"context"
	"io"
	"log"
	"strings"
	"time"

	"github.com/Nikkely/GLSite/internal/model"
	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/cdproto/dom"
	"github.com/chromedp/chromedp"
)

const (
	endpoint   = `https://www.dlsite.com/home/works/type/=/work_type_category/audio`
	waitSecond = 2
)

// Fetch scrape target
func Fetch() error {
	res, err := fetchHTML(endpoint)
	if err != nil {
		return err
	}

	// TODO: output 
	// var works *[]model.Work
	_, err = parse(strings.NewReader(res))
	return nil
}

// fetchHTML get HTML with chronium
func fetchHTML(url string) (res string, err error) {
	// create chrome instance
	ctx, cancel := chromedp.NewContext(
		context.Background(),
		// chromedp.WithDebugf(log.Printf),
	)
	defer cancel()

	// create a timeout
	ctx, cancel = context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	// 	chromedp.Value(`li#search_result_img_box_inner`, &example),
	if err = chromedp.Run(ctx,
		chromedp.Navigate(endpoint),
		chromedp.Sleep(time.Second*waitSecond),
		chromedp.ActionFunc(func(ctx context.Context) error {
			node, e := dom.GetDocument().Do(ctx)
			if e != nil {
				return e
			}
			res, e = dom.GetOuterHTML().WithNodeID(node.NodeID).Do(ctx)
			return e
		}),
	); err != nil {
		return
	}
	return
}

func parse(input io.Reader) (*[]model.Work, error) {
	doc, err := goquery.NewDocumentFromReader(input)
	if err != nil {
		return nil, err
	}

	works := []model.Work{}
	list := doc.Find(`li.search_result_img_box_inner`)
	list.Each(func(i int, s *goquery.Selection) {
		var (
			work model.Work
			ok   bool
		)
		if work.Name, ok = s.Find(`dd.work_name > div.multiline_truncate > a`).Attr("title"); !ok {
			log.Printf("work_name not found")
			return
		}
		works = append(works, work)
	})
	return &works, nil
}
