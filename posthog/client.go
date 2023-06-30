package posthog

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Client struct {
	personal_api_token string
	project_id         string
	baseURL            string
}

func newClient(personal_api_token string, project_id string) *Client {
	return &Client{
		personal_api_token: personal_api_token,
		project_id:         project_id,
		baseURL:            "https://app.posthog.com/api/",
	}
}
func (c *Client) doRequest(method string, path string, body ...string) string {
	var b string
	if len(body) > 0 {
		b = body[0]
	}
	client := &http.Client{}
	req, _ := http.NewRequest(method, c.baseURL+path, strings.NewReader(b))
	req.Header.Set("Authorization", "Bearer "+c.personal_api_token)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Errored when sending request to the posthog")
		return ""
	}
	defer resp.Body.Close()
	responseBody, _ := ioutil.ReadAll(resp.Body)
	return string(responseBody)
}
