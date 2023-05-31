package response

import (
	common2 "github.com/snowpal/pitch-building-projects-sdk/lib/structs/common"
)

type Dashboard struct {
	Dashboard DashboardResources `json:"dashboard"`
}

type DashboardResources struct {
	RecentlyModifiedResources *RecentlyModifiedResources `json:"recentlyModified"`
	ShortlyDueResources       *DueShortlyResources       `json:"dueShortly"`
}

type RecentlyModifiedKeys struct {
	Keys []common2.SlimKey `json:"keys"`
}

type RecentlyModifiedResources struct {
	Projects []DashboardProject `json:"projects"`
	Cards    []DashboardCard    `json:"pods"`
}

type DueShortlyResources struct {
	Projects *[]DashboardProject `json:"projects"`
	Cards    []DashboardCard     `json:"pods"`
	Tasks    []DashboardTask     `json:"tasks"`
}

type DashboardProject struct {
	ID      string `json:"id"`
	Name    string `json:"projectName"`
	DueDate string `json:"projectDueDate"`

	Key         *common2.SlimKey `json:"key"`
	ProjectType *ProjectType     `json:"projectType"`

	Creator  common2.ResourceCreator  `json:"creator"`
	Modifier common2.ResourceModifier `json:"modifier"`
}

type DashboardCard struct {
	ID      string `json:"id"`
	Name    string `json:"podName"`
	DueDate string `json:"podDueDate"`

	Key      *common2.SlimKey       `json:"key"`
	Projects *[]ProjectWithCardType `json:"projects"`

	Creator  common2.ResourceCreator  `json:"creator"`
	Modifier common2.ResourceModifier `json:"modifier"`
}

type ProjectWithCardType struct {
	ID       string           `json:"id"`
	Name     string           `json:"projectName"`
	Key      *common2.SlimKey `json:"key"`
	CardType *CardType        `json:"podType"`
}

type DashboardTask struct {
	ID      string `json:"id"`
	Name    string `json:"taskName"`
	DueDate string `json:"taskDueDate"`

	Key      *common2.SlimKey       `json:"key"`
	Project  *common2.SlimProject   `json:"project"`
	Card     *common2.SlimCard      `json:"pod"`
	Projects *[]common2.SlimProject `json:"projects"`

	Creator  common2.ResourceCreator  `json:"creator"`
	Modifier common2.ResourceModifier `json:"modifier"`
}

type DashboardUnreadCount struct {
	DueTasks    int `json:"dueTasks"`
	DueProjects int `json:"dueProjects"`
	DueCards    int `json:"dueCards"`

	Notifications int `json:"notifications"`
	Conversations int `json:"conversations"`
}
