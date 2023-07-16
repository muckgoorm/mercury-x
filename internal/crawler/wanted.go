package crawler

import "github.com/gocolly/colly/v2"

type WantedCrawler struct {
	Crawler
	Collector *colly.Collector
}

func NewWantedCrawler() Crawler {
	collector := colly.NewCollector()
	return &WantedCrawler{
		Collector: collector,
	}
}

func (c *WantedCrawler) SearchJobPostings(JobSearchPayload) []JobPosting {
	// url := "https://www.wanted.co.kr/wdlist/518?country=kr&job_sort=job.latest_order&years=-1&locations=all"

	return []JobPosting{}
}
