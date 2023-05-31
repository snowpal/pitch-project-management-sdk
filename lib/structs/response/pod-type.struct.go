package response

import (
	common2 "github.com/snowpal/pitch-building-projects-sdk/lib/structs/common"
)

type CardTypes struct {
	CardTypes []CardType `json:"podTypes"`
}

type CardType struct {
	ID    string              `json:"id"`
	Name  string              `json:"podTypeName"`
	Cards *[]common2.SlimCard `json:"pods"`

	TeacherReadOnly *bool `json:"teacherReadOnly"`

	Modifier     common2.ResourceModifier `json:"modifier"`
	LastModified string                   `json:"lastModified"`
}
