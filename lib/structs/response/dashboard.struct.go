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
	Projects []DashboardBlock `json:"projects"`
	Pods     []DashboardPod   `json:"pods"`
}

type DueShortlyResources struct {
	Projects *[]DashboardBlock `json:"projects"`
	Pods     []DashboardPod    `json:"pods"`
	Tasks    []DashboardTask   `json:"tasks"`
}

type DashboardBlock struct {
	ID      string `json:"id"`
	Name    string `json:"projectName"`
	DueDate string `json:"projectDueDate"`

	Key       *common2.SlimKey `json:"key"`
	BlockType *BlockType       `json:"projectType"`

	Creator  common2.ResourceCreator  `json:"creator"`
	Modifier common2.ResourceModifier `json:"modifier"`
}

type DashboardPod struct {
	ID      string `json:"id"`
	Name    string `json:"podName"`
	DueDate string `json:"podDueDate"`

	Key      *common2.SlimKey    `json:"key"`
	Projects *[]BlockWithPodType `json:"projects"`

	Creator  common2.ResourceCreator  `json:"creator"`
	Modifier common2.ResourceModifier `json:"modifier"`
}

type BlockWithPodType struct {
	ID      string           `json:"id"`
	Name    string           `json:"projectName"`
	Key     *common2.SlimKey `json:"key"`
	PodType *PodType         `json:"podType"`
}

type DashboardTask struct {
	ID      string `json:"id"`
	Name    string `json:"taskName"`
	DueDate string `json:"taskDueDate"`

	Key      *common2.SlimKey     `json:"key"`
	Block    *common2.SlimBlock   `json:"project"`
	Pod      *common2.SlimPod     `json:"pod"`
	Projects *[]common2.SlimBlock `json:"projects"`

	Creator  common2.ResourceCreator  `json:"creator"`
	Modifier common2.ResourceModifier `json:"modifier"`
}

type DashboardUnreadCount struct {
	DueTasks    int `json:"dueTasks"`
	DueProjects int `json:"dueProjects"`
	DuePods     int `json:"duePods"`

	Notifications int `json:"notifications"`
	Conversations int `json:"conversations"`
}
