package slack

import "encoding/json"

type Attachment struct {
	Pretext string `json:"pretext"`
	Text    string `json:"text"`
}

type Properties map[string]interface{}

type User struct {
	Id             string     `json:"id"`
	Name           string     `json:"name"`
	Prefs          Properties `json:"prefs,omitempty"`
	Created        int        `json:"created"`
	ManualPresence string     `json:"manual_presence"`
}

type Team struct {
	Id                    string     `json:"id"`
	Name                  string     `json:"name"`
	EmailDomain           string     `json:"email_domain"`
	Domain                string     `json:"domain"`
	MessageEditWindowMins int        `json:"msg_edit_window_mins"`
	OverStorageLimit      bool       `json:"over_storage_limit"`
	Prefs                 Properties `json:"prefs,omitempty"`
	Icon                  Properties `json:"icon"`
}

type Edit struct {
	User string `json:"user"`
	Ts   string `json:"ts"`
}

type Message struct {
	Type      string `json:"type"`
	Inviter   string `json:"inviter"`
	Subtype   string `json:"subtype"`
	Hidden    bool   `json:"hidden"`
	Channel   string `json:"channel"`
	User      string `json:"user"`
	Text      string `json:"text"`
	Ts        string `json:"ts"`
	DeletedTs string `json:"deleted_ts,omitempty"`
	EventTs   string `json:"event_ts,omitempty"`
	Edited    *Edit  `json:"edited,omitempty"`
}

type Value struct {
	Value   string `json:"value"`
	Creator string `json:"creator"`
	LastSet int    `json:"last_set"`
}

type Group struct {
	Id                 string   `json:"id"`
	Name               string   `json:"name"`
	IsGroup            bool     `json:"is_group"`
	Created            int      `json:"created"`
	Creator            string   `json:"creator"`
	IsArchived         bool     `json:"is_archived"`
	IsOpen             bool     `json:"is_open"`
	LastRead           string   `json:"last_read"`
	Latest             *Message `json:"latest"`
	UnreadCount        int      `json:"unread_count"`
	UnreadCountDisplay int      `json:"unread_count_display"`
	Members            []string `json:"members"`
	Topic              Value    `json:"topic"`
	Purpose            Value    `json:"purpose"`
}

type Channel struct {
	Id                 string   `json:"id"`
	Name               string   `json:"name"`
	IsChannel          bool     `json:"is_channel"`
	Created            int      `json:"created"`
	Creator            string   `json:"creator"`
	IsArchived         bool     `json:"is_archived"`
	IsGeneral          bool     `json:"is_general"`
	IsMember           bool     `json:"is_member"`
	LastRead           string   `json:"last_read"`
	Latest             *Message `json:"latest"`
	UnreadCount        int      `json:"unread_count"`
	UnreadCountDisplay int      `json:"unread_count_display"`
	Members            []string `json:"members"`
	Topic              Value    `json:"topic"`
	Purpose            Value    `json:"purpose"`
}

type Bot struct {
	Id      string     `json:"id"`
	Name    string     `json:"name"`
	Deleted bool       `json:"deleted"`
	Icons   Properties `json:"icons"`
}

type File struct {
	Id                 string          `json:"id"`
	Created            int             `json:"created"`
	Timestamp          int             `json:"timestamp"`
	Name               string          `json:"name"`
	Title              string          `json:"title"`
	Mimetype           string          `json:"mimetype"`
	Filetype           string          `json:"filetype"`
	PrettyType         string          `json:"pretty_type"`
	User               string          `json:"user"`
	Mode               string          `json:"mode"`
	Editable           bool            `json:"editable"`
	IsExternal         bool            `json:"is_external"`
	ExternalType       string          `json:"external_type"`
	Size               int             `json:"size"`
	Url                string          `json:"url"`
	UrlDownload        string          `json:"url_download"`
	UrlPrivate         string          `json:"url_private"`
	UrlPrivateDownload string          `json:"url_private_download"`
	Thumb64            string          `json:"thumb_64"`
	Thumb80            string          `json:"thumb_80"`
	Thumb360           string          `json:"thumb_360"`
	Thumb360Gif        string          `json:"thumb_360_gif"`
	Thumb360W          string          `json:"thumb_360_w"`
	Thumb360H          string          `json:"thumb_360_h"`
	Permalink          string          `json:"permalink"`
	EditLink           string          `json:"edit_link"`
	Preview            string          `json:"preview"`
	PreviewHighlight   string          `json:"preview_highlight"`
	Lines              int             `json:"lines"`
	LinesMore          int             `json:"lines_more"`
	IsPublic           bool            `json:"is_public"`
	PublicUrlShared    bool            `json:"public_url_shared"`
	Channels           []string        `json:"channels"`
	Groups             []string        `json:"groups"`
	InitialComment     json.RawMessage `json:"initial_comment"`
	NumStars           int             `json:"num_stars"`
	IsStarred          bool            `json:"is_starred"`
}
