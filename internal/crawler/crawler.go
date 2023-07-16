package crawler

type crawler interface {
	checkValidation() bool
	searchJobPostings(JobSearchPayload) []JobPosting
	parseJobDescription(url string) JobDescription
}

type JobSearchPayload struct {
	Platform   string
	Role       string
	Experience string
	Stacks     []string
	Benefits   []string
}

type JobPosting struct {
	Company string
	Role    string
	URL     string
}

type JobDescription struct {
	MainTasks []string
	Required  []string
	Preferred []string
	Benefits  []string
	Process   []string
	Location  string
}
