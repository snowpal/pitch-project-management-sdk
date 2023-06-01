package request

type NoteReqBody struct {
	NoteText string `json:"noteText"`
}

type NoteIdParam struct {
	KeyId     string
	ProjectId *string
	CardId    *string
	NoteId    *string
}
