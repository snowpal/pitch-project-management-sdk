package request

type AddProjectReqBody struct {
	Name string `json:"projectName"`
}

type ProjectAclReqBody struct {
	Acl string `json:"projectAcl"`
}

type GetProjectsParam struct {
	KeyId            string
	BatchIndex       int
	WriteOrHigherAcl bool
	Filter           string
}

type CopyMoveProjectParam struct {
	ProjectId   string
	KeyId       string
	TargetKeyId string

	CardIds       []string
	AllCards      bool
	AllTasks      bool
	AllChecklists bool
}
