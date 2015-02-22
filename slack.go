package slack

import (
	"net/url"

	"github.com/savaki/slack/form"
)

type Client struct {
	get  GetFunc
	post PostFunc
}

func New(token string) *Client {
	get, post := newApiFunc(token)

	client := &Client{
		get:  get,
		post: post,
	}

	return client
}

type ApiTestReq struct {
	Foo string `form:"foo"`
	Err string `form:"err"`
}

type ApiTestResp struct {
	Ok    bool              `json:"ok"`
	Error string            `json:"error,omitempty"`
	Args  map[string]string `json:"args,omitempty"`
}

func (c *Client) ApiTest(input ApiTestReq) (*ApiTestResp, error) {
	values := form.AsValues(input)
	resp := &ApiTestResp{}
	err := c.get("api.test", values, resp)
	return resp, err
}

type AuthTestResponse struct {
	Ok     bool   `json:"ok"`
	Error  string `json:"error,omitempty"`
	Url    string `json:"url"`
	Team   string `json:"team"`
	User   string `json:"user"`
	TeamId string `json:"team_id"`
	UserId string `json:"user_id"`
}

func (c *Client) AuthTest() (*AuthTestResponse, error) {
	return authTest(c.get)
}

func authTest(get GetFunc) (*AuthTestResponse, error) {
	resp := &AuthTestResponse{}
	err := get("auth.test", url.Values{}, resp)
	return resp, err
}
