package response

import (
	common2 "github.com/snowpal/pitch-building-projects-sdk/lib/structs/common"
)

type StudentGradeForProjectAndCard struct {
	ID           string                 `json:"id"`
	Name         string                 `json:"projectName"`
	Key          common2.SlimKey        `json:"key"`
	Card         *common2.SlimCard      `json:"pod"`
	StudentGrade *StudentGrade          `json:"scaleValue"`
	Cards        *[]StudentGradeForCard `json:"pods"`
	Students     *[]Student             `json:"students"`
}

type StudentGradeForCard struct {
	ID           string        `json:"id"`
	Name         string        `json:"podName"`
	StudentGrade *StudentGrade `json:"scaleValue"`
}

type StudentGrade struct {
	ScaleValue   string `json:"scaleValue"`
	NumericScale int    `json:"numericScale"`
	Published    bool   `json:"published"`
	PublishedOn  string `json:"publishedOn"`
}

type Students struct {
	Students []Student `json:"students"`
}

type Student struct {
	ID            string               `json:"id"`
	ProfileID     string               `json:"profileId"`
	Email         string               `json:"email"`
	Username      string               `json:"username"`
	FirstName     string               `json:"firstName"`
	MiddleName    string               `json:"middleName"`
	LastName      string               `json:"lastName"`
	PhoneNumber   string               `json:"phoneNumber"`
	AddressUserBy string               `json:"addressUserBy"`
	UserInitial   string               `json:"userInitial"`
	AvatarName    string               `json:"avatarName"`
	AvatarUrl     string               `json:"avatarUrl"`
	ProjectName   string               `json:"projectName"`
	StudentGrade  *StudentGrade        `json:"scaleValue"`
	Key           *common2.SlimKey     `json:"key"`
	Project       *common2.SlimProject `json:"project"`
}
