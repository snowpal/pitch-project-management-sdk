package cards

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/snowpal/pitch-project-management-sdk/lib"
	helpers2 "github.com/snowpal/pitch-project-management-sdk/lib/helpers"
	"github.com/snowpal/pitch-project-management-sdk/lib/structs/request"
	"github.com/snowpal/pitch-project-management-sdk/lib/structs/response"
)

func GetCardNotes(jwtToken string, noteParam request.NoteIdParam) ([]response.Note, error) {
	resNotes := response.Notes{}
	route, err := helpers2.GetRoute(
		lib.RouteCardsGetCardNotes,
		*noteParam.CardId,
		noteParam.KeyId,
		*noteParam.ProjectId,
	)
	if err != nil {
		return resNotes.Notes, err
	}

	req, err := http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		return resNotes.Notes, err
	}

	helpers2.AddUserHeaders(jwtToken, req)

	res, err := helpers2.MakeRequest(req)
	if err != nil {
		return resNotes.Notes, err
	}

	defer helpers2.CloseBody(res.Body)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return resNotes.Notes, err
	}

	err = json.Unmarshal(body, &resNotes)
	if err != nil {
		return resNotes.Notes, err
	}
	return resNotes.Notes, nil
}
