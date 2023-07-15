package cmd

import "github.com/urfave/cli/v2"

func checkFunction(c *cli.Context) error {
	return nil
}

func checkCommand() *cli.Command {
	return &cli.Command{
		Name:        "check",
		Usage:       "채용 플랫폼들에 대한 크롤링 작업 유효성 확인",
		Description: "채용 플랫폼의 웹 페이지 양식 변경 등의 이유로 크롤러가 오작동하는지를 미리 확인합니다.",
		Aliases:     []string{"c"},
		Action:      checkFunction,
	}
}
