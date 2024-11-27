package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/iazkaban/comfy4go/client/websocket_message_model"
	"io"
	"math/rand"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path"
	"strings"
)

var (
	ErrorContentTypeIsUnknown = errors.New("content type is unknown")
)

type UpLoadImage struct {
	Name      string `json:"name"`
	SubFolder string `json:"subfolder"`
	Type      string `json:"type"`
}

func (client *Client) DownloadImage(info *websocket_message_model.Image) (string, []byte, error) {
	apiUri := "/api/view"
	apiMethod := http.MethodGet

	apiUrl := url.URL{
		Scheme: "http",
		Host:   client.host,
		Path:   apiUri,
	}

	apiParams := make(map[string]string)
	apiParams["filename"] = info.Filename
	apiParams["subfolder"] = info.SubFolder
	apiParams["type"] = info.Type
	apiParams["rand"] = fmt.Sprintf("%f", rand.Float64())

	paramsList := make([]string, 0, len(apiParams))
	for k, v := range apiParams {
		paramsList = append(paramsList, fmt.Sprintf("%s=%s", k, v))
	}
	apiUrl.RawQuery = strings.Join(paramsList, "&")

	req, err := http.NewRequest(apiMethod, apiUrl.String(), nil)
	if err != nil {
		return "", nil, err
	}

	req.Header.Add("Comfy-User", client.UserID)

	resp, err := client.client.Do(req)
	if err != nil {
		return "", nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", nil, err
	}

	if len(resp.Header["Content-Type"]) > 0 {
		return resp.Header["Content-Type"][0], body, nil
	}

	return "", body, ErrorContentTypeIsUnknown
}

func (client *Client) UploadImage(filename string, body []byte) (*UpLoadImage, error) {
	_, filename = path.Split(filename)
	apiUri := "/api/upload/image"
	apiMethod := http.MethodPost

	apiUrl := url.URL{
		Scheme: "http",
		Host:   client.host,
		Path:   apiUri,
	}

	newBody := &bytes.Buffer{}
	writer := multipart.NewWriter(newBody)
	part, err := writer.CreateFormFile("image", filename)
	if err != nil {
		return nil, err
	}
	_, err = part.Write(body)
	if err != nil {
		return nil, err
	}
	writer.Close()

	req, err := http.NewRequest(apiMethod, apiUrl.String(), newBody)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	resp, err := client.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	rs := &UpLoadImage{}
	err = json.Unmarshal(respBody, rs)
	if err != nil {
		return nil, err
	}

	return rs, nil
}

func (client *Client) UploadImageFile(filename string) (*UpLoadImage, error) {
	apiUri := "/api/upload/image"
	apiMethod := http.MethodPost

	apiUrl := url.URL{
		Scheme: "http",
		Host:   client.host,
		Path:   apiUri,
	}

	file, err := os.OpenFile(filename, 0644, os.ModePerm)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	_, filename = path.Split(filename)
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("image", filename)
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(part, file)
	if err != nil {
		return nil, err
	}
	writer.Close()

	req, err := http.NewRequest(apiMethod, apiUrl.String(), body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	resp, err := client.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	rs := &UpLoadImage{}
	err = json.Unmarshal(respBody, rs)
	if err != nil {
		return nil, err
	}

	return rs, nil
}
