package survey

import (
	"fmt"
	"os"

	"github.com/AlecAivazis/survey/v2"
	"github.com/AlecAivazis/survey/v2/terminal"
)

func SingleChoice(label string, opt []string) string {
	var res string
	prompt := &survey.Select{
		Message:  label,
		Options:  opt,
		PageSize: 10,
		Help:     "위아래 방향키로 커서를 옮기고, 엔터로 선택합니다.",
	}
	err := survey.AskOne(prompt, &res)
	if err != nil {
		if err == terminal.InterruptErr {
			os.Exit(128)
		}
		fmt.Println(err.Error())
		os.Exit(1)
	}

	return res
}

func MultipleChoice(label string, opt []string, maxCount int) []string {
	var res []string
	prompt := &survey.MultiSelect{
		Message:  label,
		Options:  opt,
		PageSize: 10,
		Help:     "위아래 방향키로 커서를 옮기고, 스페이스바로 선택합니다. 엔터로 선택을 완료합니다.",
	}
	err := survey.AskOne(prompt, &res)
	if err != nil {
		if err == terminal.InterruptErr {
			os.Exit(128)
		}
		fmt.Println(err.Error())
		os.Exit(1)
	}

	if len(res) > maxCount {
		fmt.Printf("최대 %d개까지 선택할 수 있습니다.\n", maxCount)
		return MultipleChoice(label, opt, maxCount)
	}

	return res
}
