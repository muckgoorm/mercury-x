package pkg

import (
	"github.com/AlecAivazis/survey/v2"
)

func SingleChoice(label string, opt []string) string {
	res := ""
	prompt := &survey.Select{
		Message:  label,
		Options:  opt,
		PageSize: 6,
		Help:     "위아래 방향키로 커서를 옮기고, 엔터로 선택합니다.",
	}
	survey.AskOne(prompt, &res)

	return res
}

func MultipleChoice(label string, opt []string) []string {
	res := []string{}
	prompt := &survey.MultiSelect{
		Message:  label,
		Options:  opt,
		PageSize: 6,
		Help:     "위아래 방향키로 커서를 옮기고, 스페이스바로 선택합니다. 엔터로 선택을 완료합니다.",
	}
	survey.AskOne(prompt, &res)

	return res
}
