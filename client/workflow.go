package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type Workflow struct {
	Path     string  `json:"path"`
	Size     int     `json:"size"`
	Modified float64 `json:"modified"`
}

type WorkflowDetail struct {
	Config interface{} `json:"config"`
	Extra  struct {
		Ds struct {
			Scale  float64   `json:"scale"`
			Offset []float64 `json:"offset"`
		} `json:"ds"`
		GroupNodes map[string]struct {
			//----------------
		} `json:"groupNodes"`
	} `json:"extra"`
	Groups     []interface{}   `json:"groups"`
	LastLinkID int             `json:"last_link_id"`
	LastNodeID int             `json:"last_node_id"`
	Links      [][]interface{} `json:"links"`
	Nodes      []struct {
		ID      int        `json:"id"`
		Type    string     `json:"type"`
		Flags   struct{}   `json:"flags"`
		Inputs  []struct{} `json:"inputs"`
		Mode    int        `json:"mode"`
		Order   int        `json:"order"`
		Outputs []struct {
			Name      string `json:"name"`
			Label     string `json:"label"`
			Links     []int  `json:"links"`
			SlotIndex int    `json:"slot_index"`
			Type      string `json:"type"`
		} `json:"outputs"`
		Pos           []float64              `json:"pos"`
		Properties    map[string]interface{} `json:"properties"`
		Size          []float64              `json:"size"`
		WidgetsValues []interface{}          `json:"widgets_values"`
	} `json:"nodes"`
	Version float64 `json:"version"`
}

func (client *Client) WorkflowList() ([]*Workflow, error) {
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

	rs := make([]*Workflow, 0)
	err = json.Unmarshal(body, &rs)
	if err != nil {
		return nil, err
	}

	return rs, nil
}

func (client *Client) Workflow(path string) (*WorkflowDetail, error) {
	apiUri := "/api/userdata/workflows%2F" + path
	apiMethod := http.MethodGet

	apiUrl := url.URL{
		Scheme: "http",
		Host:   client.host,
		Path:   apiUri,
	}

	req, err := http.NewRequest(apiMethod, apiUrl.String(), nil)
	if err != nil {
		panic(err)
		return nil, err
	}

	req.Header.Add("Comfy-User", client.UserID)

	resp, err := client.client.Do(req)
	if err != nil {
		panic(err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
		return nil, err
	}

	rs := &WorkflowDetail{}
	err = json.Unmarshal(body, rs)
	if err != nil {
		fmt.Println(req.URL)
		fmt.Println(resp.StatusCode)
		fmt.Println(string(body))
		panic(err)
		return nil, err
	}

	return rs, nil
}
