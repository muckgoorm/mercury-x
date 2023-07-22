package internal

type Crawler interface {
	CheckValidation() bool
	SearchJobPostings(JobSearchPayload) ([]JobPosting, error)
	ParseJobDescription(url string) JobDescription
}

type JobSearchPayload struct {
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
