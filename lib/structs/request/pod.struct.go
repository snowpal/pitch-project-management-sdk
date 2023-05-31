package request

type AddCardReqBody struct {
	Name string `json:"podName"`
}

type UpdateCardDescReqBody struct {
	Description   string  `json:"podDescription"`
	TaggedUserIds *string `json:"taggedUserIds"`
}

type BulkArchiveCardsReqBody struct {
	CardIds string `json:"podIds"`
}

type UpdateCardStatusReqBody struct {
	Completed bool `json:"podCompleted"`
}

type UpdateCardReqBody struct {
	Name              *string `json:"podName"`
	SimpleDescription *string `json:"simpleDescription"`
	DueDate           *string `json:"podDueDate"`
	Color             *string `json:"podColor"`
	Tags              *string `json:"podTags"`
	KanbanMode        *bool   `json:"kanbanMode"`
}

type CardAclReqBody struct {
	Acl string `json:"podAcl"`
}

type CardBulkShareReqBody struct {
	Acl     string `json:"podAcl"`
	CardIds string `json:"podIds"`
}

type GetCardsParam struct {
	KeyId      string
	ProjectId  *string
	BatchIndex int
}

type AddCardTypeIdParam struct {
	CardId     string
	CardTypeId string
	KeyId      string
	ProjectId  *string
}

type CardByTemplateParam struct {
	KeyId        string
	ProjectId    *string
	TemplateId   string
	ExcludeTasks bool
}

type CopyMoveCardParam struct {
	CardId      string
	KeyId       string
	TargetKeyId string

	ProjectId       string
	TargetProjectId string

	AllTasks      bool
	AllChecklists bool
}
