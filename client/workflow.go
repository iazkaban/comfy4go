package client

import (
	"encoding/json"
	"fmt"
	"github.com/iazkaban/comfy4go/model"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func (client *Client) WorkflowList() ([]*model.Workflow, error) {
	apiUri := "/api/userdata"
	apiMethod := http.MethodGet

	apiUrl := url.URL{
		Scheme: "http",
		Host:   client.host,
		Path:   apiUri,
	}

	apiParams := make(map[string]string)
	apiParams["dir"] = "workflows"
	apiParams["recurse"] = "true"
	apiParams["split"] = "false"
	apiParams["full_info"] = "true"

	paramsList := make([]string, 0, len(apiParams))
	for k, v := range apiParams {
		paramsList = append(paramsList, fmt.Sprintf("%s=%s", k, v))
	}
	apiUrl.RawQuery = strings.Join(paramsList, "&")

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

	rs := make([]*model.Workflow, 0)
	err = json.Unmarshal(body, &rs)
	if err != nil {
		return nil, err
	}

	return rs, nil
}

func (client *Client) Workflow(path string) (*model.WorkflowDetail, error) {
	apiUri := "/api/userdata/workflows%2F" + path
	apiMethod := http.MethodGet

	apiUrl := url.URL{
		Scheme: "http",
		Host:   client.host,
		Path:   apiUri,
	}

	req, err := http.NewRequest(apiMethod, apiUrl.String(), nil)
	if err != nil {
		client.log.Error(err)
		return nil, err
	}

	req.Header.Add("Comfy-User", client.UserID)

	resp, err := client.client.Do(req)
	if err != nil {
		client.log.Error(err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, ErrorsServerError
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		client.log.Error(err)
		return nil, err
	}

	rs := &model.WorkflowDetail{}
	err = json.Unmarshal(body, rs)
	if err != nil {
		return nil, err
	}

	return rs, nil
}
