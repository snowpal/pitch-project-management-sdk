package request

type ScaleReqBody struct {
	Name        string   `json:"scaleName"`
	Type        string   `json:"scaleType"`
	ScaleValues []string `json:"scaleValues"`
}

type UpdateScaleValueReqBody struct {
	ScaleValue string `json:"scaleValue"`
}

type ScaleIdParam struct {
	ScaleId   string
	KeyId     string
	ProjectId *string
	CardId    *string
}
