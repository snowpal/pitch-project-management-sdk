package request

type AddBlockReqBody struct {
	Name string `json:"projectName"`
}

type BlockAclReqBody struct {
	Acl string `json:"projectAcl"`
}

type GetProjectsParam struct {
	KeyId            string
	BatchIndex       int
	WriteOrHigherAcl bool
	Filter           string
}

type CopyMoveBlockParam struct {
	BlockId     string
	KeyId       string
	TargetKeyId string

	PodIds        []string
	AllPods       bool
	AllTasks      bool
	AllChecklists bool
}
