package fetcher

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/Nikkely/GLSite/model"
	"github.com/PuerkitoBio/goquery"
)

const (
	urlPrefix  = `https://www.dlsite.com/home/works/type/=/language/jp/age_category%5B0%5D/general/work_category%5B0%5D/doujin/order%5B0%5D/trend/work_type_category%5B0%5D/audio/work_type_category_name%5B0%5D/%E3%83%9C%E3%82%A4%E3%82%B9%E3%83%BBASMR/options_and_or/and/options%5B0%5D/JPN/options%5B1%5D/NM/per_page/100/show_type/3/lang_options%5B0%5D/%E6%97%A5%E6%9C%AC%E8%AA%9E/lang_options%5B1%5D/%E8%A8%80%E8%AA%9E%E4%B8%8D%E8%A6%81/page/`
	waitSecond = 1
)

// joinFilePath makes path
func joinFilePath(filename string, paths ...string) (path string) {
	for _, p := range paths {
		if !strings.HasSuffix(p, "/") {
			p += "/"
		}
		path += p
	}
	path += filename
	return
}

// Fetch scrape target
func Fetch(dir string, options ...FetchOption) error {
	option := defaultOption(dir)
	for _, f := range options {
		f(&option)
	}

	var works []model.Work
	for p := 1; option.unlimitPage || p <= option.limitPage; p++ {
		log.Printf("fetch page %d", p)
		res, err := fetchHTML(urlPrefix + strconv.Itoa(p))
		end := false
		if err != nil {
			switch e := err.(type) {
			case notFoundErr:
				log.Printf("page %d is not found\n", p)
				end = true
			case otherErr:
				return fmt.Errorf("page %d is failed; %s", p, e.Error())
			default:
				return e
			}
		}
		if end {
			break
		}
		log.Println("start to parse ")

		result, err := parse(strings.NewReader(res))
		if err != nil {
			return err
		}
		works = append(works, result...)

		time.Sleep(waitSecond * time.Second)
	}
	return option.writer.Write(works)
}

// fetchHTML get HTML with chronium
func fetchHTML(url string) (string, error) {
	res, err := http.Get(url)
	if err != nil {
		return "", err
	}
	if strings.Contains(res.Status, "404") {
		return "", newNotFoundErr()
	}
	if !strings.Contains(res.Status, "200") {
		return "", newOtherErr(res.Status)
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	buf := bytes.NewBuffer(body)
	return buf.String(), nil
}

var workIDReg = regexp.MustCompile(`https://www\.dlsite\.com/home/work/=/product_id/(.+)\.html`)
var makerIDReg = regexp.MustCompile(`https://www\.dlsite\.com/home/circle/profile/=/maker_id/(.+)\.html`)
var numReg = regexp.MustCompile(`\d+`)

// parse parses model from html
func parse(input io.Reader) ([]model.Work, error) {
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
		work.FetchedAt = time.Now()
		if work.URL, ok = s.Find(`dd.work_name > div.multiline_truncate > a`).Attr("href"); !ok {
			log.Println("url not found")
			return
		}
		if matches := workIDReg.FindStringSubmatch(work.URL); len(matches) < 1 {
			log.Println("failed parsing work id")
			return
		} else {
			work.ID = matches[1]
		}
		if work.Name, ok = s.Find(`dd.work_name > div.multiline_truncate > a`).Attr("title"); !ok {
			log.Println("work_name not found")
			return
		}
		work.Maker = model.Maker{}
		if work.Maker.URL, ok = s.Find(`dd.maker_name > a`).Attr("href"); !ok {
			log.Println("maker url not found")
			return
		}
		if matches := makerIDReg.FindStringSubmatch(work.Maker.URL); len(matches) < 1 {
			log.Println("failed parsing maker id")
			return
		} else {
			work.Maker.ID = matches[1]
		}
		if work.Maker.Name = s.Find(`dd.maker_name > a`).Text(); work.Maker.Name == "" {
			log.Println("maker_name not found")
			return
		}
		if work.Author = s.Find(`span.author > a`).Text(); work.Author == "" {
			log.Printf("author not found: %s", work.URL) // continue if discount was not found
		}
		if t := s.Find(`span.discount`).Text(); t == "" {
			if t := s.Find(`span.work_price`).Text(); t == "" {
				log.Printf("discount and price not found")
				return
			} else {
				if work.Price, err = strconv.Atoi(strings.Join(numReg.FindAllString(t, -1), "")); err != nil {
					log.Printf("price format unexpected: %s\n", t)
					return
				}
			}
		} else {
			if work.Discount, err = strconv.Atoi(strings.Join(numReg.FindAllString(t, -1), "")); err != nil {
				log.Printf("discount format unexpected: %s\n", strings.Join(numReg.FindAllString(t, -1), ""))
				return
			}
			if t := s.Find(`span.strike`).Text(); t == "" {
				log.Printf("discount found but price not found: %s", work.URL) // continue if discount was not found
				return
			} else {
				if work.Price, err = strconv.Atoi(strings.Join(numReg.FindAllString(t, -1), "")); err != nil {
					log.Printf("price format unexpected: %s\n", t)
					return
				}
			}
		}
		if t := s.Find(`span._dl_count_` + work.ID).Text(); t == "" {
			log.Printf("dl count not found: %s", work.URL) // continue if discount was not found
		} else {
			if work.DL, err = strconv.Atoi(strings.Join(numReg.FindAllString(t, -1), "")); err != nil {
				log.Printf("dl count fomat unexpected: %s\n", t)
				return
			}
		}
		if r := s.Find(`dd.work_rating`); len(r.Nodes) == 0 {
			log.Printf("rating not found: %s", work.URL) // continue if discount was not found
		} else {
			star := -1
			for s := 0; s <= 50; s += 5 {
				if len(r.Find(`div.star_`+strconv.Itoa(s)).Nodes) > 0 {
					star = s
					break
				}
			}
			if star == -1 {
				log.Println("rating star not found")
				return
			}
			work.RatingStar = star
			if t := r.Find(`div.star_` + strconv.Itoa(star)).Text(); t == "" {
				log.Println("rating total not found")
				return
			} else {
				if work.RatingTotal, err = strconv.Atoi(strings.Join(numReg.FindAllString(t, -1), "")); err != nil {
					log.Printf("rating total format unexpected: %s\n", t)
					return
				}
			}
		}

		works = append(works, work)
		log.Printf("progress: %d / %d\n", index+1, len(list.Nodes))
	})
	return works, nil
}
