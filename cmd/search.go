package cmd

import (
	"mercury-x/internal"
	"mercury-x/internal/wanted"
	"mercury-x/pkg/survey"
	"mercury-x/pkg/tui"
	"strconv"
)

func search() {
	platform := survey.SingleChoice("검색할 플랫폼을 골라주세요", internal.Platforms)
	benefitOpts, benefitMaximum := setBenefitSearchOptions(platform)
	role := survey.SingleChoice("직군을 골라주세요", internal.Roles)
	experience := survey.SingleChoice("경력을 골라주세요", internal.Experiences)
	stacks := survey.MultipleChoice("기술 스택을 골라주세요", internal.Stacks, 5)
	benefits := survey.MultipleChoice("복지 여건을 선택해주세요", benefitOpts, benefitMaximum)
	count, _ := strconv.Atoi(survey.SingleChoice("보고 싶은 채용 공고의 개수를 입력해주세요", internal.Counts))

	payload := internal.JobSearchPayload{
		Role:       role,
		Experience: experience,
		Stacks:     stacks,
		Benefits:   benefits,
		Count:      count,
	}

	var postings []internal.JobPosting

	switch platform {
	case "wanted":
		wc := wanted.NewWantedCrawler()
		postings, _ = wc.SearchJobPostings(payload)
	}

	tui.SearchTable(postings)
}

func setBenefitSearchOptions(platform string) (opts []string, maximum int) {
	switch platform {
	case "wanted":
		opts = internal.Benefits
		maximum = 3
	}
	return
}
