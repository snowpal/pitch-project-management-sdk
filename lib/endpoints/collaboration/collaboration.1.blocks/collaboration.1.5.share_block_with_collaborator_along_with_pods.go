package collaboration

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

type ShareBlockWithPodsReqBody struct {
	Acl    string `json:"projectAcl"`
	PodIds string `json:"podIds"`
}

func ShareBlockWithUserWithPods(
	jwtToken string,
	reqBody ShareBlockWithPodsReqBody,
	projectAclParam common.AclParam,
) (response.Block, error) {
	resBlock := response.Block{}
	requestBody, err := helpers2.GetRequestBody(reqBody)
	if err != nil {
		fmt.Println(err)
		return resBlock, err
	}
	payload := strings.NewReader(requestBody)
	route, err := helpers2.GetRoute(
		lib.RouteCollaborationShareBlockWithCollaboratorAlongWithPods,
		projectAclParam.ResourceIds.BlockId,
		projectAclParam.UserId,
		projectAclParam.ResourceIds.KeyId,
	)
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
