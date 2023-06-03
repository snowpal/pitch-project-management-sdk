package response

import (
	common2 "github.com/snowpal/pitch-project-management-sdk/lib/structs/common"
)

type UserKeys struct {
	Keys []UserKey `json:"keys"`
}

type UserKey struct {
	ID           string             `json:"id"`
	Name         string             `json:"keyName"`
	Type         string             `json:"keyType"`
	Projects     []UserProject      `json:"projects"`
	Cards        []common2.SlimCard `json:"cards"`
	LastModified string             `json:"lastModified"`
}

type UserProject struct {
	ID           string             `json:"id"`
	Name         string             `json:"keyName"`
	Cards        []common2.SlimCard `json:"cards"`
	LastModified string             `json:"lastModified"`
}

type FilteredKeys struct {
	Keys []FilteredKey `json:"keys"`
}

type ProjectsAndCards struct {
	Projects []common2.SlimProject `json:"projects"`
	Cards    []common2.SlimCard    `json:"cards"`
}

type FilteredKey struct {
	ID               string           `json:"id"`
	Name             string           `json:"keyName"`
	Type             string           `json:"keyType"`
	CreatedByMe      ProjectsAndCards `json:"createdByMe"`
	SharedWithMe     ProjectsAndCards `json:"sharedWithMe"`
	SharedWithOthers ProjectsAndCards `json:"sharedWithOthers"`
}

type ProjectTypesKeys struct {
	Keys []ProjectTypesKey `json:"keys"`
}

type ProjectTypesKey struct {
	ID           string        `json:"id"`
	Name         string        `json:"keyName"`
	Type         string        `json:"keyType"`
	ProjectTypes []ProjectType `json:"projectTypes"`
	LastModified string        `json:"lastModified"`
}

type CardTypesKeysCard struct {
	Card CardTypesKeys `json:"card"`
}

type CardTypesKeys struct {
	Keys *[]CardTypesKey `json:"keys"`
	Key  *CardTypesKey   `json:"key"`
}

type CardTypesKey struct {
	ID           string      `json:"id"`
	Name         string      `json:"keyName"`
	Type         string      `json:"keyType"`
	CardTypes    *[]CardType `json:"cardTypes"`
	LastModified string      `json:"lastModified"`
}

type ScalesKeysProject struct {
	Project ScalesKeys `json:"project"`
}

type ScalesKeysCard struct {
	Card ScalesKeys `json:"card"`
}

type ScalesKeys struct {
	Keys *[]ScalesKey `json:"keys"`
	Key  *ScalesKey   `json:"key"`
}

type ScalesKey struct {
	ID           string   `json:"id"`
	Name         string   `json:"keyName"`
	Type         string   `json:"keyType"`
	Scales       *[]Scale `json:"scales"`
	LastModified string   `json:"lastModified"`
}

type TaskStatus struct {
	Complete   int `json:"complete"`
	Incomplete int `json:"incomplete"`
}

type TasksStatusKeys struct {
	Keys []TasksStatusKey `json:"keys"`
}

type TasksStatusKey struct {
	ID           string     `json:"id"`
	Name         string     `json:"keyName"`
	TaskStatus   TaskStatus `json:"taskStatus"`
	LastModified string     `json:"lastModified"`
}

type TasksStatusProject struct {
	ID           string          `json:"id"`
	Name         string          `json:"projectName"`
	TaskStatus   TaskStatus      `json:"taskStatus"`
	Key          common2.SlimKey `json:"key"`
	LastModified string          `json:"lastModified"`
}

type LinkedResourcesKey struct {
	SlimKey common2.SlimKey `json:"key"`
}

type LinkedResources struct {
	CurrentKey LinkedResourcesKey `json:"currentKey"`
	SharedKey  LinkedResourcesKey `json:"sharedKey"`
	Keys       *[]UserKey         `json:"keys"`
	Projects   []UserProject      `json:"projects"`
}

type ProjectScaleValue struct {
	ID           string `json:"id"`
	Name         string `json:"projectName"`
	ScaleValue   string `json:"scaleValue"`
	NumericScale int    `json:"numericScale"`
}

type CardScaleValue struct {
	ID           string `json:"id"`
	Name         string `json:"cardName"`
	ScaleValue   string `json:"scaleValue"`
	NumericScale int    `json:"numericScale"`
}

type ScaleValues struct {
	Scale    Scale                `json:"scale"`
	Key      common2.SlimKey      `json:"key"`
	Project  *common2.SlimProject `json:"project"`
	Projects *[]ProjectScaleValue `json:"projects"`
	Cards    []CardScaleValue     `json:"cards"`
}
