package response

import (
	common2 "github.com/snowpal/pitch-project-management-sdk/lib/structs/common"
)

type Comments struct {
	Comments []Comment `json:"comments"`
}

type Comment struct {
	ID          string       `json:"id"`
	CommentText string       `json:"commentText"`
	TaggedUsers []TaggedUser `json:"taggedUsers"`

	CanEdit   *bool `json:"canEdit"`
	CanDelete *bool `json:"canDelete"`

	Key     *common2.SlimKey     `json:"key"`
	Project *common2.SlimProject `json:"project"`
	Card    *common2.SlimCard    `json:"card"`

	Creator      common2.ResourceCreator  `json:"creator"`
	Modifier     common2.ResourceModifier `json:"modifier"`
	LastModified string                   `json:"lastModified"`
}

type RecentComments struct {
	Comments []RecentComment `json:"comments"`
}

type RecentComment struct {
	ID          string `json:"id"`
	CommentText string `json:"commentText"`

	Key      *common2.SlimKey       `json:"key"`
	Project  *common2.SlimProject   `json:"project"`
	Projects *[]common2.SlimProject `json:"projects"`
	Card     *common2.SlimCard      `json:"card"`

	Creator      common2.ResourceCreator  `json:"creator"`
	Modifier     common2.ResourceModifier `json:"modifier"`
	LastModified string                   `json:"lastModified"`
}
