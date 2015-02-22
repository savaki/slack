package slack

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

type GetFunc func(string, url.Values, interface{}) error

type PostFunc func(method string, contentType string, payload io.Reader, v interface{}) error

func newApiFunc(token string) (GetFunc, PostFunc) {
	toUrl := func(method string, params url.Values) string {
		return "https://slack.com/api/" + method + "?" + params.Encode()
	}

	getFunc := func(method string, params url.Values, v interface{}) error {
		params.Set("token", token)

		resp, err := http.Get(toUrl(method, params))
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		return json.NewDecoder(resp.Body).Decode(v)
	}

	postFunc := func(method string, contentType string, payload io.Reader, v interface{}) error {
		params := url.Values{}
		params.Set("token", token)

		resp, err := http.Post(toUrl(method, params), contentType, payload)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		return json.NewDecoder(resp.Body).Decode(v)
	}

	return getFunc, postFunc
}
