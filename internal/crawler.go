package internal

type Crawler interface {
	CheckValidation() bool
	SearchJobPostings(JobSearchPayload) ([]JobPosting, error)
	ParseJobDescription(url string) (JobDescription, error)
}

type JobSearchPayload struct {
	Role       string
	Experience string
	Stacks     []string
	Benefits   []string
	Count      int
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
	Process   []string
	Location  string
	Benefits  []string
}
