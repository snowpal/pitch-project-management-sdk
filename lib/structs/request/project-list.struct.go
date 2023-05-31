package request

type AddProjectListReqBody struct {
	Name string `json:"projectListName"`
}

type AssignCardReqBody struct {
	UserIds string `json:"userIds"`
}

type ProjectListIdParam struct {
	KeyId         string
	ProjectId     string
	ProjectListId string
	CardId        *string
}

type CopyMoveProjectListParam struct {
	KeyId         string
	ProjectId     string
	ProjectListId string

	TargetKeyId     string
	TargetProjectId string
	TargetPosition  int
}

type CopyMoveProjectCardParam struct {
	CardId    string
	ProjectId string
	KeyId     string

	TargetKeyId         string
	TargetProjectId     string
	TargetProjectListId string
}

type CopyMoveProjectListCardsParam struct {
	KeyId         string
	ProjectId     string
	ProjectListId string

	TargetKeyId         string
	TargetProjectId     string
	TargetProjectListId string

	AllTasks bool
	AllCards bool

	CardIds []string
}
