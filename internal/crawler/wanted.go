package crawler

import (
	"fmt"

	"github.com/gocolly/colly/v2"
)

type WantedCrawler struct {
	crawler
	Collector *colly.Collector
}

func NewWantedCrawler() crawler {
	collector := colly.NewCollector()
	return &WantedCrawler{
		Collector: collector,
	}
}

func (c *WantedCrawler) SearchJobPostings(j JobSearchPayload) []JobPosting {
	// url := "https://www.wanted.co.kr/wdlist/518?country=kr&job_sort=job.latest_order&years=-1&locations=all"
	role := MapRole(j.Role)
	exp := MapExperience(j.Experience)
	stacks := MapStacks(j.Stacks)

	url := fmt.Sprintf(
		"https://www.wanted.co.kr/wdlist/%s%s?country=kr&job_sort=job.latest_order&years=%s%s&locations=all",
		dev, role, exp, stacks,
	)

	fmt.Println(url)

	return []JobPosting{}
}
