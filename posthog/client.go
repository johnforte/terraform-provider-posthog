package posthog

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type Client struct {
	personal_api_token string
	project_id         string
	baseURL            string
}

func New(personal_api_token string, project_id string) *Client {
	return &Client{
		personal_api_token: personal_api_token,
		project_id:         project_id,
		baseURL:            "https://app.posthog.com/api/",
	}
}
func (c *Client) doRequest(method string, path string) string {
	client := &http.Client{}
	req, _ := http.NewRequest(method, c.baseURL+path, nil)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Errored when sending request to the server")
		return ""
	}
	defer resp.Body.Close()
	responseBody, _ := ioutil.ReadAll(resp.Body)
	return string(responseBody)
}
