package client

import (
	"bytes"
	"encoding/json"
	"github.com/iazkaban/comfy4go/model"
	"io"
	"net/http"
	"net/url"
)

func (client *Client) Prompt(requestBody *model.PromptRequest) (*model.PromptResponse, error) {
	apiUri := "/api/prompt"
	apiMethod := http.MethodPost

	apiUrl := url.URL{
		Scheme: "http",
		Host:   client.host,
		Path:   apiUri,
	}

	requestBody.ClientID = client.ClientID
	reqBody, err := json.Marshal(requestBody)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(apiMethod, apiUrl.String(), bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Comfy-User", client.UserID)

	resp, err := client.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	rs := &model.PromptResponse{}
	err = json.Unmarshal(body, rs)
	if err != nil {
		return nil, err
	}

	return rs, nil
}
