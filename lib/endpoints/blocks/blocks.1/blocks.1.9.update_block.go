package projects

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/snowpal/pitch-building-projects-sdk/lib"
	helpers2 "github.com/snowpal/pitch-building-projects-sdk/lib/helpers"
	"github.com/snowpal/pitch-building-projects-sdk/lib/structs/common"
	"github.com/snowpal/pitch-building-projects-sdk/lib/structs/response"
)

type UpdateBlockReqBody struct {
	Name              *string `json:"projectName"`
	BlockId           *string `json:"projectId"`
	SimpleDescription *string `json:"simpleDescription"`
	DueDate           *string `json:"projectDueDate"`
	StartTime         *string `json:"projectStartTime"`
	EndTime           *string `json:"projectEndTime"`
	Color             *string `json:"projectColor"`
	Tags              *string `json:"projectTags"`
	KanbanMode        *bool   `json:"kanbanMode"`
	Completed         bool    `json:"projectCompleted"`
}

func UpdateBlock(jwtToken string, reqBody UpdateBlockReqBody, projectParam common.ResourceIdParam) (response.Block, error) {
	resBlock := response.Block{}
	requestBody, err := helpers2.GetRequestBody(reqBody)
	if err != nil {
		fmt.Println(err)
		return resBlock, err
	}
	payload := strings.NewReader(requestBody)
	route, err := helpers2.GetRoute(lib.RouteProjectsUpdateBlock, projectParam.BlockId, projectParam.KeyId)
	if err != nil {
		fmt.Println(err)
		return resBlock, err
	}

	req, err := http.NewRequest(http.MethodPatch, route, payload)
	if err != nil {
		fmt.Println(err)
		return resBlock, err
	}

	helpers2.AddUserHeaders(jwtToken, req)

	res, err := helpers2.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resBlock, err
	}

	defer helpers2.CloseBody(res.Body)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return resBlock, err
	}

	err = json.Unmarshal(body, &resBlock)
	if err != nil {
		fmt.Println(err)
		return resBlock, err
	}
	return resBlock, nil
}
