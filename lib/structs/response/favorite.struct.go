package response

import (
	"github.com/snowpal/pitch-building-projects-sdk/lib/structs/common"
)

type Favorites struct {
	Favorites []Favorite `json:"favorites"`
}

type AddFavorite struct {
	ID string `json:"id"`
}

type Favorite struct {
	ID       string                  `json:"id"`
	Modifier common.ResourceModifier `json:"modifier"`
	Resource FavoriteResource        `json:"resource"`
}

type FavoriteKey struct {
	ID       string                  `json:"id"`
	Name     string                  `json:"keyName"`
	Modifier common.ResourceModifier `json:"modifier"`
}

type FavoriteProject struct {
	ID       string                  `json:"id"`
	Name     string                  `json:"projectName"`
	Modifier common.ResourceModifier `json:"modifier"`
}

type FavoriteResource struct {
	ID           string                  `json:"id"`
	ResourceType string                  `json:"resourceType"`
	KeyName      *string                 `json:"keyName"`
	ProjectName  *string                 `json:"projectName"`
	CardName     *string                 `json:"cardName"`
	Key          *FavoriteKey            `json:"key"`
	Project      *FavoriteProject        `json:"project"`
	Modifier     common.ResourceModifier `json:"modifier"`
}
