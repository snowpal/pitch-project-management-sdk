package response

import (
	common2 "github.com/snowpal/pitch-building-projects-sdk/lib/structs/common"
)

type UserKeys struct {
	Keys []UserKey `json:"keys"`
}

type UserKey struct {
	ID           string            `json:"id"`
	Name         string            `json:"keyName"`
	Type         string            `json:"keyType"`
	Projects     []UserBlock       `json:"projects"`
	Pods         []common2.SlimPod `json:"pods"`
	LastModified string            `json:"lastModified"`
}

type UserBlock struct {
	ID           string            `json:"id"`
	Name         string            `json:"keyName"`
	Pods         []common2.SlimPod `json:"pods"`
	LastModified string            `json:"lastModified"`
}

type FilteredKeys struct {
	Keys []FilteredKey `json:"keys"`
}

type ProjectsAndPods struct {
	Projects []common2.SlimBlock `json:"projects"`
	Pods     []common2.SlimPod   `json:"pods"`
}

type FilteredKey struct {
	ID               string          `json:"id"`
	Name             string          `json:"keyName"`
	Type             string          `json:"keyType"`
	CreatedByMe      ProjectsAndPods `json:"createdByMe"`
	SharedWithMe     ProjectsAndPods `json:"sharedWithMe"`
	SharedWithOthers ProjectsAndPods `json:"sharedWithOthers"`
}

type BlockTypesKeys struct {
	Keys []BlockTypesKey `json:"keys"`
}

type BlockTypesKey struct {
	ID           string      `json:"id"`
	Name         string      `json:"keyName"`
	Type         string      `json:"keyType"`
	BlockTypes   []BlockType `json:"projectTypes"`
	LastModified string      `json:"lastModified"`
}

type PodTypesKeysKeyPod struct {
	KeyPod PodTypesKeys `json:"keyPod"`
}

type PodTypesKeysBlockPod struct {
	BlockPod PodTypesKeys `json:"card"`
}

type PodTypesKeysOtherPod struct {
	OtherPod PodTypesKeys `json:"otherPod"`
}

type PodTypesKeys struct {
	Keys *[]PodTypesKey `json:"keys"`
	Key  *PodTypesKey   `json:"key"`
}

type PodTypesKey struct {
	ID           string     `json:"id"`
	Name         string     `json:"keyName"`
	Type         string     `json:"keyType"`
	PodTypes     *[]PodType `json:"podTypes"`
	LastModified string     `json:"lastModified"`
}

type ScalesKeysBlock struct {
	Block ScalesKeys `json:"project"`
}

type ScalesKeysKeyPod struct {
	KeyPod ScalesKeys `json:"keyPod"`
}

type ScalesKeysBlockPod struct {
	BlockPod ScalesKeys `json:"card"`
}

type ScalesKeysOtherPod struct {
	OtherPod ScalesKeys `json:"otherPod"`
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

type TasksStatusBlock struct {
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
	Projects   []UserBlock        `json:"projects"`
}

type BlockScaleValue struct {
	ID           string `json:"id"`
	Name         string `json:"projectName"`
	ScaleValue   string `json:"scaleValue"`
	NumericScale int    `json:"numericScale"`
}

type PodScaleValue struct {
	ID           string `json:"id"`
	Name         string `json:"podName"`
	ScaleValue   string `json:"scaleValue"`
	NumericScale int    `json:"numericScale"`
}

type ScaleValues struct {
	Scale    Scale              `json:"scale"`
	Key      common2.SlimKey    `json:"key"`
	Block    *common2.SlimBlock `json:"project"`
	Projects *[]BlockScaleValue `json:"projects"`
	Pods     []PodScaleValue    `json:"pods"`
}

type BlockGrade struct {
	ID       string          `json:"id"`
	Name     string          `json:"projectName"`
	Key      common2.SlimKey `json:"key"`
	Students []Student       `json:"students"`
}

type PodGrade struct {
	ID       string            `json:"id"`
	Name     string            `json:"podName"`
	Key      common2.SlimKey   `json:"key"`
	Block    common2.SlimBlock `json:"project"`
	Students []Student         `json:"students"`
}
