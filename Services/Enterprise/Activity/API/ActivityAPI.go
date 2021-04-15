package Activity

type activity struct {
	id               string `json:"id"`
	activityNumber   string `json:"activity-number"`
	plannedStartDate string `json:"planned_start_date"`
	plannedEndDate   string `json:"planned_end_date"`
	startDate        string `json:"start-date"`
	endDate          string `json:"end-date"`
	activityType     string `json:"activity-type"`
	description      string `json:"description"`
	priority         int    `json:"priority"`
	targetTime       string `json:"target_time"`
	actualTime       string `json:"actual_time"`
	complexity       int    `json:"complexity"`
	actions          []activityAction
	issues           []activityIssue
}

type activityAction struct {
	id               int
	activityID       string
	plannedStartDate string `json:"planned_start_date"`
	plannedEndDate   string `json:"planned_end_date"`
	startDate        string `json:"start-date"`
	endDate          string `json:"end-date"`
	targetTime       string `json:"target_time"`
	actualTime       string `json:"actual_time"`
}

type activityIssue struct {
}

func (a *activity ) newActivity( )

