package request

type AddCardReqBody struct {
	Name string `json:"cardName"`
}

type UpdateCardDescReqBody struct {
	Description   string  `json:"cardDescription"`
	TaggedUserIds *string `json:"taggedUserIds"`
}

type BulkArchiveCardsReqBody struct {
	CardIds string `json:"cardIds"`
}

type UpdateCardStatusReqBody struct {
	Completed bool `json:"cardCompleted"`
}

type UpdateCardReqBody struct {
	Name              *string `json:"cardName"`
	SimpleDescription *string `json:"simpleDescription"`
	DueDate           *string `json:"cardDueDate"`
	Color             *string `json:"color"`
	Tags              *string `json:"cardTags"`
	KanbanMode        *bool   `json:"kanbanMode"`
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
