package cmd

import (
	"mercury-x/internal"
	"mercury-x/internal/wanted"
	"mercury-x/pkg/survey"
	"mercury-x/pkg/tui"
)

func search() error {
	var payload internal.JobSearchPayload
	platform := survey.SingleChoice("검색할 플랫폼을 골라주세요", internal.Platforms)
	benefitOpts, benefitMaximum := setBenefitSearchOptions(platform)
	payload.Role = survey.SingleChoice("직군을 골라주세요", internal.Roles)
	payload.Experience = survey.SingleChoice("경력을 골라주세요", internal.Experiences)
	payload.Stacks = survey.MultipleChoice("기술 스택을 골라주세요", internal.Stacks, 5)
	payload.Benefits = survey.MultipleChoice("복지 여건을 선택해주세요", benefitOpts, benefitMaximum)

	var postings []internal.JobPosting

	wc := wanted.NewWantedCrawler()
	postings, err := wc.SearchJobPostings(payload)
	if err != nil {
		return err
	}

	tui.JobList(postings)

	return nil
}

func setBenefitSearchOptions(platform string) (opts []string, maximum int) {
	switch platform {
	case "wanted":
		opts = internal.Benefits
		maximum = 3
	}
	return
}
