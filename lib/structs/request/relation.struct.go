package request

type KeyToKeyRelationParam struct {
	KeyId       string
	TargetKeyId string
}

type KeyToProjectRelationParam struct {
	KeyId           string
	TargetProjectId string
}

type KeyToCardRelationParam struct {
	KeyId string

	TargetCardId    string
	TargetKeyId     string
	TargetProjectId string
}

type ProjectToProjectRelationParam struct {
	ProjectId       string
	TargetProjectId string
}

type ProjectToCardRelationParam struct {
	ProjectId string

	TargetCardId    string
	TargetKeyId     string
	TargetProjectId string
}

type CardToCardRelationParam struct {
	CardId          string
	SourceKeyId     string
	SourceProjectId string

	TargetCardId    string
	TargetKeyId     string
	TargetProjectId string
}
