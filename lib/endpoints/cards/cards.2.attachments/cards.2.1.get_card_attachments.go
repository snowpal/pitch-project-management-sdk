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

func GetCardAttachments(jwtToken string, attachmentParam request.AttachmentParam) ([]response.Attachment, error) {
	resAttachments := response.Attachments{}
	route, err := helpers2.GetRoute(
		lib.RouteCardsGetCardAttachments,
		*attachmentParam.CardId,
		attachmentParam.KeyId,
		*attachmentParam.ProjectId,
	)
	if err != nil {
		return resAttachments.Attachments, err
	}

	req, err := http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		return resAttachments.Attachments, err
	}
	helpers2.AddUserHeaders(jwtToken, req)

	res, err := helpers2.MakeRequest(req)
	if err != nil {
		return resAttachments.Attachments, err
	}

	defer helpers2.CloseBody(res.Body)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return resAttachments.Attachments, err
	}

	err = json.Unmarshal(body, &resAttachments)
	if err != nil {
		return resAttachments.Attachments, err
	}
	return resAttachments.Attachments, nil
}
