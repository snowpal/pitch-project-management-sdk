package user

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/snowpal/pitch-project-management-sdk/lib"
	helpers2 "github.com/snowpal/pitch-project-management-sdk/lib/helpers"
	"github.com/snowpal/pitch-project-management-sdk/lib/structs/response"
)

func GetUserByEmail(jwtToken string, email string) (response.User, error) {
	resUser := response.User{}
	route, err := helpers2.GetRoute(lib.RouteUserGetUserByEmail, email)
	if err != nil {
		return resUser, err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		return resUser, err
	}

	helpers2.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = helpers2.MakeRequest(req)
	if err != nil {
		return resUser, err
	}

	defer helpers2.CloseBody(res.Body)

	var body []byte
	body, err = io.ReadAll(res.Body)
	if err != nil {
		return resUser, err
	}

	err = json.Unmarshal(body, &resUser)
	if err != nil {
		return resUser, err
	}
	return resUser, nil
}
