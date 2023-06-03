package project_types

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/snowpal/pitch-project-management-sdk/lib"
	helpers2 "github.com/snowpal/pitch-project-management-sdk/lib/helpers"
	"github.com/snowpal/pitch-project-management-sdk/lib/structs/response"
)

func GetProjectTypes(jwtToken string, includeCounts bool) ([]response.ProjectType, error) {
	resProjectTypes := response.ProjectTypes{}
	route, err := helpers2.GetRoute(lib.RouteProjectTypesGetProjectTypes, strconv.FormatBool(includeCounts))
	if err != nil {
		fmt.Println(err)
		return resProjectTypes.ProjectTypes, err
	}
	req, err := http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		fmt.Println(err)
		return resProjectTypes.ProjectTypes, err
	}

	helpers2.AddUserHeaders(jwtToken, req)

	res, err := helpers2.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resProjectTypes.ProjectTypes, err
	}

	defer helpers2.CloseBody(res.Body)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return resProjectTypes.ProjectTypes, err
	}

	err = json.Unmarshal(body, &resProjectTypes)
	if err != nil {
		fmt.Println(err)
		return resProjectTypes.ProjectTypes, err
	}
	return resProjectTypes.ProjectTypes, nil
}
