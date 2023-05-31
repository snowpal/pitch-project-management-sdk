package favorites

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

func AddCardAsFavorite(jwtToken string, favoriteParam common.ResourceIdParam) (response.AddFavorite, error) {
	resFavorite := response.AddFavorite{}
	route, err := helpers2.GetRoute(
		lib.RouteFavoritesAddCardAsFavorite,
		favoriteParam.CardId,
		favoriteParam.KeyId,
		favoriteParam.ProjectId,
	)
	if err != nil {
		fmt.Println(err)
		return resFavorite, err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodPost, route, nil)
	if err != nil {
		fmt.Println(err)
		return resFavorite, err
	}

	helpers2.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = helpers2.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resFavorite, err
	}

	defer helpers2.CloseBody(res.Body)

	var body []byte
	body, err = io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return resFavorite, err
	}

	err = json.Unmarshal(body, &resFavorite)
	if err != nil {
		fmt.Println(err)
		return resFavorite, err
	}
	return resFavorite, nil
}