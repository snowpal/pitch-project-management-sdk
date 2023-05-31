package response

import (
	common2 "github.com/snowpal/pitch-building-projects-sdk/lib/structs/common"
)

type AllEvents struct {
	DueDateEvent    *DueDateEvent    `json:"dueDate"`
	EndDateEvent    *EndDateEvent    `json:"endDate"`
	SchedulerEvents []SchedulerEvent `json:"schedulerEvents"`
}

type EndDateEvent struct {
	Projects []BlockEvent `json:"projects"`
}

type DueDateEvent struct {
	Tasks    TasksEvent   `json:"tasks"`
	Projects []BlockEvent `json:"projects"`
	Pods     []PodEvent   `json:"pods"`
}

type BlockEvent struct {
	ID          string `json:"id"`
	Name        string `json:"projectName"`
	Description string `json:"projectDescription"`

	DueDate   *string `json:"projectDueDate"`
	StartTime *string `json:"projectStartTime"`
	EndTime   *string `json:"projectEndTime"`

	Key common2.SlimKey `json:"key"`
}

type PodEvent struct {
	ID      string `json:"id"`
	Name    string `json:"podName"`
	DueDate string `json:"podDueDate"`

	Key   common2.SlimKey   `json:"key"`
	Block common2.SlimBlock `json:"project"`
}

type TasksEvent struct {
	KeyTasks   []TaskEvent `json:"keys"`
	BlockTasks []TaskEvent `json:"projects"`
	PodTasks   []TaskEvent `json:"pods"`
}

type TaskEvent struct {
	ID      string             `json:"id"`
	Name    string             `json:"taskName"`
	DueDate string             `json:"taskDueDate"`
	Key     common2.SlimKey    `json:"key"`
	Block   *common2.SlimBlock `json:"project"`
	Pod     *common2.SlimPod   `json:"pod"`
}

type SchedulerEvent struct {
	SchedulerId      string            `json:"schedulerId"`
	StandaloneEvents []StandaloneEvent `json:"standaloneEvents"`
}

type StandaloneEvent struct {
	ID          string `json:"id"`
	Description string `json:"description"`
	StartTime   string `json:"eventStartTime"`
	EndTime     string `json:"eventEndTime"`

	Creator  *common2.ResourceCreator  `json:"creator"`
	Modifier *common2.ResourceModifier `json:"modifier"`
}
