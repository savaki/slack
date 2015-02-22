package slack

import "net/url"

type ChannelsListResp struct {
	Ok       bool      `json:"ok"`
	Error    string    `json:"error,omitempty"`
	Channels []Channel `json:"channels,omitempty"`
}

func (s *Client) ChannelsList() (*ChannelsListResp, error) {
	resp := &ChannelsListResp{}
	err := s.get("channels.list", url.Values{}, resp)
	return resp, err
}
