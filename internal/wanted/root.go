package wanted

import (
	"context"
	"fmt"
	"log"
	"mercury-x/internal"
	"strings"
	"time"

	"github.com/chromedp/chromedp"
)

type WantedCrawler struct {
	internal.Crawler
}

func NewWantedCrawler() internal.Crawler {
	return &WantedCrawler{}
}

func (c *WantedCrawler) SearchJobPostings(j internal.JobSearchPayload) ([]internal.JobPosting, error) {
	role := mapRole(j.Role)
	exp := mapExperience(j.Experience)
	stacks := mapStacks(j.Stacks)

	ub := strings.Builder{}
	ub.WriteString("https://www.wanted.co.kr/wdlist/")
	ub.WriteString(dev)
	ub.WriteString(role)
	ub.WriteString("?country=kr&job_sort=job.latest_order&years=")
	ub.WriteString(exp)
	ub.WriteString(stacks)
	ub.WriteString("&locations=all")
	url := ub.String()
	fmt.Println(url)

	// TODO: implement the crawling logic
	ctx, cancel := chromedp.NewContext(
		context.Background(),
		chromedp.WithLogf(log.Printf),
	)
	defer cancel()

	ctx, cancel = context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	postings, err := getPostings(ctx, url)
	if err != nil {
		return nil, err
	}

	return postings, nil
}

func getPostings(ctx context.Context, url string) ([]internal.JobPosting, error) {
	var postings []internal.JobPosting
	var companies string

	err := chromedp.Run(ctx,
		chromedp.Navigate(url),
		chromedp.Text(`h3`, &companies, chromedp.NodeVisible, chromedp.ByQueryAll),
	)
	if err != nil {
		return nil, err
	}

	fmt.Println(companies)

	return postings, nil
}
