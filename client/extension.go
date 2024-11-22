package client

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

func (client *Client) Extensions() ([]string, error) {
	apiUri := "/api/extensions"
	apiMethod := http.MethodGet

	apiUrl := url.URL{
		Scheme: "http",
		Host:   client.host,
		Path:   apiUri,
	}

	req, err := http.NewRequest(apiMethod, apiUrl.String(), nil)
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

	rs := make([]string, 0)
	err = json.Unmarshal(body, &rs)
	if err != nil {
		return nil, err
	}

	return rs, nil
}
