package collaboration

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/snowpal/pitch-building-projects-sdk/lib"
	helpers2 "github.com/snowpal/pitch-building-projects-sdk/lib/helpers"
	"github.com/snowpal/pitch-building-projects-sdk/lib/structs/common"
	"github.com/snowpal/pitch-building-projects-sdk/lib/structs/response"
)

func GetBlockCollaborators(jwtToken string, projectParam common.ResourceIdParam) (response.Block, error) {
	resBlock := response.Block{}
	route, err := helpers2.GetRoute(
		lib.RouteCollaborationGetBlockCollaborators,
		projectParam.BlockId,
		projectParam.KeyId,
	)
	if err != nil {
		fmt.Println(err)
		return resBlock, err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		fmt.Println(err)
		return resBlock, err
	}

	helpers2.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = helpers2.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resBlock, err
	}

	defer helpers2.CloseBody(res.Body)

	var body []byte
	body, err = io.ReadAll(res.Body)
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
