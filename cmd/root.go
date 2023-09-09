package cmd

import (
	"mercury-x/pkg/survey"
	"os"
)

const s = "search"

const (
	FAILED            = 1
	INVALID_INPUT     = 2
	COMMAND_NOT_FOUND = 127
)

func Root(args []string) {
	if len(args) == 0 || len(args) > 2 {
		os.Exit(INVALID_INPUT)
	} else if len(args) == 1 {
		args = append(args, selectCommand())
	}

	switch args[1] {
	case s:
		err := search()
		if err != nil {
			os.Exit(FAILED)
		}
	default:
		os.Exit(COMMAND_NOT_FOUND)
	}
}

func selectCommand() string {
	sDesc := "원하는 조건에 맞는 채용 공고 검색하기"
	res := survey.SingleChoice("무엇을 하고 싶으신가요?", []string{sDesc})

	var cmd string
	switch res {
	case sDesc:
		cmd = s
	}

	return cmd
}
