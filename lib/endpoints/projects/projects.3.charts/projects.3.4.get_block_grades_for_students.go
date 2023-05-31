package projects

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

type ProjectGradesForStudents struct {
	Project response.ProjectGrade `json:"project"`
	Cards   []response.CardGrade  `json:"pods"`
}

func GetProjectGradesForStudents(jwtToken string, gradeParam common.ResourceIdParam) (ProjectGradesForStudents, error) {
	resProjectGrades := ProjectGradesForStudents{}
	route, err := helpers2.GetRoute(
		lib.RouteProjectsGetProjectGradesForStudents,
		gradeParam.ProjectId,
		gradeParam.KeyId,
	)
	if err != nil {
		fmt.Println(err)
		return resProjectGrades, err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		fmt.Println(err)
		return resProjectGrades, err
	}

	helpers2.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = helpers2.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resProjectGrades, err
	}

	defer helpers2.CloseBody(res.Body)

	var body []byte
	body, err = io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return resProjectGrades, err
	}

	err = json.Unmarshal(body, &resProjectGrades)
	if err != nil {
		fmt.Println(err)
		return resProjectGrades, err
	}
	return resProjectGrades, nil
}
