package response

import (
	common2 "github.com/snowpal/pitch-project-management-sdk/lib/structs/common"
)

type AllEvents struct {
	DueDateEvent    *DueDateEvent    `json:"dueDate"`
	EndDateEvent    *EndDateEvent    `json:"endDate"`
	SchedulerEvents []SchedulerEvent `json:"schedulerEvents"`
}

type EndDateEvent struct {
	Projects []ProjectEvent `json:"projects"`
}

type DueDateEvent struct {
	Tasks    TasksEvent     `json:"tasks"`
	Projects []ProjectEvent `json:"projects"`
	Cards    []CardEvent    `json:"cards"`
}

type ProjectEvent struct {
	ID          string `json:"id"`
	Name        string `json:"projectName"`
	Description string `json:"projectDescription"`

	DueDate   *string `json:"projectDueDate"`
	StartTime *string `json:"projectStartTime"`
	EndTime   *string `json:"projectEndTime"`

	Key common2.SlimKey `json:"key"`
}

type CardEvent struct {
	ID      string `json:"id"`
	Name    string `json:"cardName"`
	DueDate string `json:"cardDueDate"`

	Key     common2.SlimKey     `json:"key"`
	Project common2.SlimProject `json:"project"`
}

type TasksEvent struct {
	KeyTasks     []TaskEvent `json:"keys"`
	ProjectTasks []TaskEvent `json:"projects"`
	CardTasks    []TaskEvent `json:"cards"`
}

type TaskEvent struct {
	ID      string               `json:"id"`
	Name    string               `json:"taskName"`
	DueDate string               `json:"taskDueDate"`
	Key     common2.SlimKey      `json:"key"`
	Project *common2.SlimProject `json:"project"`
	Card    *common2.SlimCard    `json:"card"`
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
