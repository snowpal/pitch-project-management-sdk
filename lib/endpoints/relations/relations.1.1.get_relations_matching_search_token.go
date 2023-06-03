package relations

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/snowpal/pitch-project-management-sdk/lib"
	helpers2 "github.com/snowpal/pitch-project-management-sdk/lib/helpers"
	"github.com/snowpal/pitch-project-management-sdk/lib/structs/response"
)

type SearchKeyRelationParam struct {
	Token        string
	CurrentKeyId string
}

type SearchProjectRelationParam struct {
	Token            string
	CurrentProjectId string
	KeyId            string
}

type SearchCardRelationParam struct {
	Token         string
	CurrentCardId string
	KeyId         string
	ProjectId     string
}

func searchRelationsMatchingSearchToken(jwtToken string, route string) ([]response.SearchResource, error) {
	resSearchResources := response.SearchResources{}
	req, err := http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		fmt.Println(err)
		return resSearchResources.Results, err
	}

	helpers2.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = helpers2.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resSearchResources.Results, err
	}

	defer helpers2.CloseBody(res.Body)

	var body []byte
	body, err = io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return resSearchResources.Results, err
	}

	err = json.Unmarshal(body, &resSearchResources)
	if err != nil {
		fmt.Println(err)
		return resSearchResources.Results, err
	}
	return resSearchResources.Results, nil
}

func SearchRelationsForKeyMatchingSearchToken(
	jwtToken string,
	relationParam SearchKeyRelationParam,
) ([]response.SearchResource, error) {
	var searchResults []response.SearchResource
	route, err := helpers2.GetRoute(
		lib.RouteRelationsGetRelationsForKeyMatchingSearchToken,
		relationParam.Token,
		relationParam.CurrentKeyId,
	)
	if err != nil {
		fmt.Println(err)
		return searchResults, err
	}
	searchResults, err = searchRelationsMatchingSearchToken(jwtToken, route)
	if err != nil {
		fmt.Println(err)
		return searchResults, err
	}
	return searchResults, nil
}

func SearchRelationsForProjectMatchingSearchToken(
	jwtToken string,
	relationParam SearchProjectRelationParam,
) ([]response.SearchResource, error) {
	var searchResults []response.SearchResource
	route, err := helpers2.GetRoute(
		lib.RouteRelationsGetRelationsForProjectMatchingSearchToken,
		relationParam.Token,
		relationParam.CurrentProjectId,
		relationParam.KeyId,
	)
	if err != nil {
		fmt.Println(err)
		return searchResults, err
	}
	searchResults, err = searchRelationsMatchingSearchToken(jwtToken, route)
	if err != nil {
		fmt.Println(err)
		return searchResults, err
	}
	return searchResults, nil
}

func SearchRelationsForCardMatchingSearchToken(
	jwtToken string,
	relationParam SearchCardRelationParam,
) ([]response.SearchResource, error) {
	var searchResults []response.SearchResource
	route, err := helpers2.GetRoute(
		lib.RouteRelationsGetRelationsForCardMatchingSearchToken,
		relationParam.Token,
		relationParam.CurrentCardId,
		relationParam.KeyId,
		relationParam.ProjectId,
	)
	if err != nil {
		fmt.Println(err)
		return searchResults, err
	}
	searchResults, err = searchRelationsMatchingSearchToken(jwtToken, route)
	if err != nil {
		fmt.Println(err)
		return searchResults, err
	}
	return searchResults, nil
}
