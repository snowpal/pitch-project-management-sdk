package response

import (
	common2 "github.com/snowpal/pitch-building-projects-sdk/lib/structs/common"
)

type Projects struct {
	Projects []Block `json:"projects"`
}

type Block struct {
	ID                string  `json:"id"`
	Name              string  `json:"projectName"`
	BlockId           string  `json:"projectId"`
	Description       string  `json:"projectDescription"`
	SimpleDescription string  `json:"simpleDescription"`
	Color             string  `json:"color"`
	Tags              string  `json:"tags"`
	ScaleValue        *string `json:"scaleValue"`

	Attributes  []common2.DisplayAttribute `json:"attributes"`
	BlockType   *BlockType                 `json:"projectType"`
	Scale       *Scale                     `json:"scale"`
	TaggedUsers []TaggedUser               `json:"taggedUsers"`
	Key         *common2.SlimKey           `json:"key"`

	// Boolean Attributes
	Completed  *bool `json:"completed"`
	Archived   *bool `json:"archived"`
	KanbanMode *bool `json:"kanbanMode"`
	CanUnlink  *bool `json:"canUnlink"`
	PublicKey  *bool `json:"publicKey"`

	// Project Key Attribute
	ProjectKanbanMode *bool `json:"projectKanbanMode"`

	// Time Attributes
	DueDate   string `json:"projectDueDate"`
	StartTime string `json:"projectStartTime"`
	EndTime   string `json:"projectEndTime"`

	// Acl Attributes
	Acl            *string       `json:"acl"`
	IsShared       *bool         `json:"isShared"`
	CanLeave       *bool         `json:"canLeave"`
	AllowDelete    *bool         `json:"allowDelete"`
	CanDelete      *bool         `json:"canDelete"`
	SharedUsers    *[]SharedUser `json:"sharedUsers"`
	CurrentUserAcl SharedUser    `json:"currentUserAcl"`

	// Count Attributes
	KeysCount        *int `json:"keysCount"`
	PodsCount        *int `json:"podsCount"`
	TasksCount       *int `json:"tasksCount"`
	ChecklistsCount  *int `json:"checklistsCount"`
	AttachmentsCount *int `json:"attachmentsCount"`
	CommentsCount    *int `json:"commentsCount"`
	NotesCount       *int `json:"notesCount"`

	Creator      common2.ResourceCreator  `json:"creator"`
	Modifier     common2.ResourceModifier `json:"modifier"`
	LastModified string                   `json:"lastModified"`
}

type UpdateBlockScaleValue struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	ScaleValue   string `json:"scaleValue"`
	NumericScale int    `json:"numericScale"`

	Key common2.SlimKey `json:"key"`
}
