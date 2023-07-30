package wanted

import (
	"fmt"
	"mercury-x/internal"
	"mercury-x/pkg/webdriver"
	"strings"
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

	wd, err := webdriver.NewWebDriver("http://localhost:9515")
	if err != nil {
		return nil, err
	}
	defer wd.Driver.Quit()

	if err := wd.Get(url); err != nil {
		return nil, err
	}

	if err := wd.ScrollToBottom(2); err != nil {
		return nil, err
	}

	companies, err := wd.FindByClassName(company)
	if err != nil {
		return nil, err
	}

	roles, err := wd.FindByClassName(position)
	if err != nil {
		return nil, err
	}

	urls, err := wd.FindLinks(card)
	if err != nil {
		return nil, err
	}

	var postings []internal.JobPosting
	for i := 0; i < len(companies); i++ {
		postings = append(postings, internal.JobPosting{
			Company: companies[i],
			Role:    roles[i],
			URL:     urls[i],
		})
	}

	return postings, nil
}
