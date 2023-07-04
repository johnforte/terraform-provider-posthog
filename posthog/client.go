package posthog

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
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
		baseURL:            fmt.Sprintf("https://app.posthog.com/api/projects/%s/", project_id),
	}
}
func (c *Client) doRequest(method string, path string, body []byte) io.Reader {
	client := &http.Client{}
	req, _ := http.NewRequest(method, c.baseURL+path, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.personal_api_token))
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Errored when sending request to the posthog")
		fmt.Println(err)
		return nil
	}
	defer resp.Body.Close()
	responseBody, _ := ioutil.ReadAll(resp.Body)
	return bytes.NewReader(responseBody)
}
