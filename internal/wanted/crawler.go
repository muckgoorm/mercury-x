package wanted

import (
	"fmt"
	"mercury-x/internal"
	"mercury-x/pkg/webdriver"
	"strings"
	"time"
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

	ub := strings.Builder{}
	ub.WriteString("https://www.wanted.co.kr/wdlist/")
	ub.WriteString(dev)
	ub.WriteString(role)
	ub.WriteString("?country=kr&job_sort=job.latest_order&years=")
	ub.WriteString(exp)
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

	if len(j.Stacks) > 0 {
		if err := wd.ClickButtonByDivClass(skillAddButton); err != nil {
			return nil, err
		}

		for _, stack := range j.Stacks {
			if err := wd.FulfillInput(skillInput, stack); err != nil {
				return nil, err
			}
		}

		if err := wd.ClickButtonByDataAttributeId(skillApplyButton); err != nil {
			return nil, err
		}
		wd.WaitForDivClass(jobList)
	}

	if len(j.Benefits) > 0 {
		benefitTags := mapBenefitsToTags(j.Benefits)
		for _, tag := range benefitTags {
			switch tag {
			case workFromHomeTag:
				if err := wd.ClickButtonByDataTag(workFromHomeTag); err != nil {
					return nil, err
				}
			case flexibleTag:
				if err := wd.ClickButton(nextButton); err != nil {
					return nil, err
				}
				time.Sleep(1 * time.Second)
				if err := wd.ClickButtonByDataTag(flexibleTag); err != nil {
					return nil, err
				}
			case flatTag:
				if err := wd.ClickButtonByDataTag(flatTag); err != nil {
					return nil, err
				}
			case snackTag:
				if err := wd.ClickButton(nextButton); err != nil {
					return nil, err
				}
				time.Sleep(1 * time.Second)
				if err := wd.ClickButtonByDataTag(snackTag); err != nil {
					return nil, err
				}
			}
			time.Sleep(1 * time.Second)
		}
	}

	scrollCount := (j.Count / 20) - 1
	if scrollCount > 0 {
		if err := wd.ScrollToBottom(scrollCount); err != nil {
			return nil, err
		}
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
