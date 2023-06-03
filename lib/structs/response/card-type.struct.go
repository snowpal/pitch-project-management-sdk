package response

import (
	common2 "github.com/snowpal/pitch-project-management-sdk/lib/structs/common"
)

type CardTypes struct {
	CardTypes []CardType `json:"cardTypes"`
}

type CardType struct {
	ID    string              `json:"id"`
	Name  string              `json:"cardTypeName"`
	Cards *[]common2.SlimCard `json:"cards"`

	Modifier     common2.ResourceModifier `json:"modifier"`
	LastModified string                   `json:"lastModified"`
}
