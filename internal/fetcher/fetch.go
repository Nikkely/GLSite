package fetcher

import (
	"context"
	"fmt"
	"io"
	"log"
	"regexp"
	"strconv"
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

var idReg = regexp.MustCompile(`https://www\.dlsite\.com/home/work/=/product_id/(.+)\.html`)

func parse(input io.Reader) (*[]model.Work, error) {
	doc, err := goquery.NewDocumentFromReader(input)
	if err != nil {
		return nil, err
	}

	works := []model.Work{}
	list := doc.Find(`li.search_result_img_box_inner`)
	list.Each(func(index int, s *goquery.Selection) {
		var (
			work model.Work
			ok   bool
		)
		if work.URL, ok = s.Find(`dd.work_name > div.multiline_truncate > a`).Attr("href"); !ok {
			log.Println("url not found")
			return
		}
		var matches []string
		if matches = idReg.FindStringSubmatch(work.URL); len(matches) < 1 {
			log.Println("failed parsing id")
			return
		}
		work.ID = matches[1]
		if work.Name, ok = s.Find(`dd.work_name > div.multiline_truncate > a`).Attr("title"); !ok {
			log.Println("work_name not found")
			return
		}
		if work.MakerName = s.Find(`dd.maker_name > a`).Text(); work.MakerName == "" {
			log.Println("maker_name not found")
			return
		}
		if work.Author = s.Find(`span.author > a`).Text(); work.Author == "" {
			log.Println("author not found")
			return
		}
		if t := s.Find(`span.discount`).Text(); t == "" {
			log.Println("discount not found")
			return
		} else {
			if work.Discount, err = strconv.Atoi(strings.Replace(strings.Replace(t, ",", "", -1), "円", "", -1)); err != nil {
				log.Println("discount format unexpected")
				return
			}
		}
		if t := s.Find(`span.strike`).Text(); t != "" {
			if work.Price, err = strconv.Atoi(strings.Replace(strings.Replace(t, ",", "", -1), "円", "", -1)); err != nil {
				log.Println("price format unexpected")
				return
			}
		}
		if t := s.Find(`span._dl_count_` + work.ID).Text(); t == "" {
			log.Println("dl count not found")
			return
		} else {
			if work.DL, err = strconv.Atoi(strings.Replace(t, ",", "", -1)); err != nil {
				log.Println("dl count fomat unexpected")
				return
			}
		}
		if r := s.Find(`dd.work_rating`); len(r.Nodes) == 0 {
			log.Println("rating star not found")
			return
		} else {
			star := -1
			for s := 0; s <= 50; s += 5 {
				if len(r.Find(`div.star_`+strconv.Itoa(s)).Nodes) > 0 {
					star = s
					break
				}
			}
			if star == -1 {
				log.Println("rating star format unexpected")
				return
			}
			work.RatingStar = star
			if t := r.Find(`div.star_` + strconv.Itoa(star)).Text(); t == "" {
				log.Println("rating total not found")
				return
			} else {
				if work.RatingTotal, err = strconv.Atoi(
					strings.Replace(
						strings.Replace(
							strings.Replace(t, ",", "", -1), "(", "", -1), ")", "", -1)); err != nil {
					log.Println("rating total unexpected")
					return
				}
			}
		}

		works = append(works, work)
		fmt.Printf("progress: %d / %d", index, len(list.Nodes))
	})
	return &works, nil
}
