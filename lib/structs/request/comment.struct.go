package request

type CommentReqBody struct {
	CommentText   string  `json:"commentText"`
	TaggedUserIds *string `json:"taggedUserIds"`
}

type CommentIdParam struct {
	KeyId     string
	ProjectId *string
	CardId    *string
	CommentId *string
}
