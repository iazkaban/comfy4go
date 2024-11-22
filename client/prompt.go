package client

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

type PromptResponse struct {
	NodeErrors struct{} `json:"node_errors"`
	Number     int      `json:"number"`
	PromptID   string   `json:"prompt_id"`
}

func (client *Client) Prompt(requestBody interface{}) (*PromptResponse, error) {
	apiUri := "/api/prompt"
	apiMethod := http.MethodPost

	apiUrl := url.URL{
		Scheme: "http",
		Host:   client.host,
		Path:   apiUri,
	}

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

	rs := &PromptResponse{}
	err = json.Unmarshal(body, rs)
	if err != nil {
		return nil, err
	}

	return rs, nil
}
