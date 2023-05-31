package blocks

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/snowpal/pitch-building-blocks-sdk/lib"
	helpers2 "github.com/snowpal/pitch-building-blocks-sdk/lib/helpers"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/structs/common"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/structs/response"
)

func GetProjectsLinkedToPods(jwtToken string, blockParam common.ResourceIdParam) ([]response.Block, error) {
	resProjects := response.Projects{}
	route, err := helpers2.GetRoute(
		lib.RouteProjectsGetProjectsLinkedToPod,
		blockParam.PodId,
		blockParam.KeyId,
	)
	if err != nil {
		fmt.Println(err)
		return resProjects.Projects, err
	}

	req, _ := http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		fmt.Println(err)
		return resProjects.Projects, err
	}

	helpers2.AddUserHeaders(jwtToken, req)
	res, err := helpers2.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resProjects.Projects, err
	}

	defer helpers2.CloseBody(res.Body)

	body, _ := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return resProjects.Projects, err
	}

	err = json.Unmarshal(body, &resProjects)
	if err != nil {
		fmt.Println(err)
		return resProjects.Projects, err
	}
	return resProjects.Projects, nil
}
