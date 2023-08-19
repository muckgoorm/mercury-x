package tui

import (
	"fmt"
	"mercury-x/internal"
	"mercury-x/internal/wanted"
	"strconv"
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type search struct {
	jd           []internal.JobDescription
	app          *tview.Application
	contents     *tview.Flex
	jobs         *tview.List
	selector     *tview.TextView
	tasks        *tview.TextView
	requirements *tview.TextView
	notes        *tview.TextView
	commands     *tview.TextView
}

func SearchTable(postings []internal.JobPosting) {
	var s search
	s.jd = make([]internal.JobDescription, len(postings))
	s.app = tview.NewApplication()
	s.initJobs(postings)
	s.initSelector()
	s.initContents()
	s.initCommands()
	flex := s.generateFlex()
	s.app.SetRoot(flex, true).EnableMouse(true)
	s.app.SetInputCapture(func(e *tcell.EventKey) *tcell.EventKey {
		switch key := e.Rune(); key {
		case 'q':
			s.app.Stop()
		}
		return e
	})
	s.app.Run()
}

func (s *search) initJobs(postings []internal.JobPosting) {
	list := tview.NewList().ShowSecondaryText(false)
	list.SetBorder(true).SetTitle("채용 회사 / 직무")
	for _, p := range postings {
		list.AddItem(fmt.Sprintf("%s / %s", p.Company, p.Role), fmt.Sprintf("%s,%s", p.Platform, p.URL), 0, nil)
	}
	list.SetChangedFunc(s.jobChange)
	list.SetSelectedFunc(s.jobSelect)

	s.jobs = list
}

func (s *search) initSelector() {
	items := []string{"주요 업무", "자격 요건", "참고 사항"}
	tv := tview.NewTextView().
		SetTextAlign(tview.AlignCenter).
		SetDynamicColors(true).
		SetRegions(true).
		SetWrap(false).
		SetHighlightedFunc(s.selectorHighlight)
	tv.SetBorder(true)

	previusSlide := func() {
		slide, _ := strconv.Atoi(tv.GetHighlights()[0])
		slide = (slide - 1 + len(items)) % len(items)
		tv.Highlight(strconv.Itoa(slide)).ScrollToHighlight()
	}
	nextSlide := func() {
		slide, _ := strconv.Atoi(tv.GetHighlights()[0])
		slide = (slide + 1) % len(items)
		tv.Highlight(strconv.Itoa(slide)).ScrollToHighlight()
	}

	for i, item := range items {
		fmt.Fprintf(tv, `["%d"][green]%s[white][""]  `, i, item)
	}

	tv.Highlight("0")
	tv.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyEsc:
			s.app.SetFocus(s.jobs)
		case tcell.KeyLeft:
			previusSlide()
		case tcell.KeyRight:
			nextSlide()
		}
		return event
	})

	s.selector = tv
}

func (s *search) initContents() {
	s.contents = tview.NewFlex()
	s.contents.SetBorder(true)
	s.tasks = tview.NewTextView().SetDynamicColors(true).SetWrap(false)
	s.requirements = tview.NewTextView().SetDynamicColors(true).SetWrap(false)
	s.notes = tview.NewTextView().SetDynamicColors(true).SetWrap(false)
}

func (s *search) initCommands() {
	items := []string{"enter/click: 선택", "q: 종료"}
	tv := tview.NewTextView().
		SetTextAlign(tview.AlignCenter).
		SetDynamicColors(true).
		SetRegions(true).
		SetWrap(false)
	tv.SetBorder(true).SetTitle("단축키")
	for i, item := range items {
		fmt.Fprintf(tv, `["%d"][yellow]%s[white][""]  `, i, item)
	}

	s.commands = tv
}

// build from inside out
func (s *search) generateFlex() *tview.Flex {
	infoTab := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(s.selector, 3, 1, false).
		AddItem(s.contents, 0, 5, false)
	mainTab := tview.NewFlex().
		AddItem(s.jobs, 0, 4, true).
		AddItem(infoTab, 0, 5, false)
	entireTab := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(mainTab, 0, 5, true).
		AddItem(s.commands, 3, 1, false)

	return entireTab
}

func (s *search) jobChange(index int, mainText string, secondaryText string, shortcut rune) {
	s.loadJD(index, secondaryText)
}

func (s *search) selectorHighlight(added, removed, remaining []string) {
	if s.contents == nil {
		return
	}
	s.contents.Clear()
	idx, _ := strconv.Atoi(added[0])
	switch idx {
	case 0:
		s.contents.AddItem(s.tasks, 0, 1, true)
	case 1:
		s.contents.AddItem(s.requirements, 0, 1, true)
	case 2:
		s.contents.AddItem(s.notes, 0, 1, true)
	}
}

func (s *search) jobSelect(index int, mainText string, secondaryText string, shortcut rune) {
	s.tasks.Clear()
	s.requirements.Clear()
	s.notes.Clear()
	s.loadJD(index, secondaryText)
	for _, task := range s.jd[index].MainTasks {
		fmt.Fprintf(s.tasks, "[white]%s\n", task)
	}
	for _, required := range s.jd[index].Required {
		fmt.Fprintf(s.requirements, "[white]%s\n", required)
	}
	fmt.Fprintf(s.requirements, "[white]---\n")
	for _, preferred := range s.jd[index].Preferred {
		fmt.Fprintf(s.requirements, "[white]%s\n", preferred)
	}
	for _, note := range s.jd[index].Benefits {
		fmt.Fprintf(s.notes, "[white]%s\n", note)
	}
	s.app.SetFocus(s.selector)
}

func (s *search) loadJD(index int, secondaryText string) {
	// TODO: implement loading indicator
	go func() {
		link := strings.Split(secondaryText, ",")
		platform, url := link[0], link[1]
		var job internal.JobDescription
		switch platform {
		case "wanted":
			wc := wanted.NewWantedCrawler()
			job, _ = wc.ParseJobDescription(url)
		}
		s.jd[index] = job
	}()
}
