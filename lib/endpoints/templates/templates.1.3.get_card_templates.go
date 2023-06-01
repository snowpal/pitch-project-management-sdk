package templates

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/snowpal/pitch-building-projects-sdk/lib"
	helpers2 "github.com/snowpal/pitch-building-projects-sdk/lib/helpers"
	"github.com/snowpal/pitch-building-projects-sdk/lib/structs/response"
)

func GetCardTemplates(jwtToken string) ([]response.CardTemplate, error) {
	resCardTemplates := response.CardTemplates{}
	route, err := helpers2.GetRoute(lib.RouteTemplatesGetCardTemplates)
	if err != nil {
		fmt.Println(err)
		return resCardTemplates.Templates, err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		fmt.Println(err)
		return resCardTemplates.Templates, err
	}

	helpers2.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = helpers2.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resCardTemplates.Templates, err
	}

	defer helpers2.CloseBody(res.Body)

	var body []byte
	body, err = io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return resCardTemplates.Templates, err
	}

	err = json.Unmarshal(body, &resCardTemplates)
	if err != nil {
		fmt.Println(err)
		return resCardTemplates.Templates, err
	}
	return resCardTemplates.Templates, nil
}