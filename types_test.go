package slack

import (
	"encoding/json"
	"strings"

	. "github.com/smartystreets/goconvey/convey"

	"testing"
)

func TestTypes(t *testing.T) {
	Convey("Given a json file type message", t, func() {
		text := `{
    "id" : "F2147483862",
    "created" : 1356032811,
    "timestamp" : 1356032811,

    "name" : "file.htm",
    "title" : "My HTML file",
    "mimetype" : "text/plain",
    "filetype" : "text",
    "pretty_type": "Text",
    "user" : "U2147483697",

    "mode" : "hosted",
    "editable" : true,
    "is_external": false,
    "external_type": "",

    "size" : 12345,

    "url": "https://slack-files.com/files-pub/T024BE7LD-F024BERPE-09acb6/1.png",
    "url_download": "https://slack-files.com/files-pub/T024BE7LD-F024BERPE-09acb6/download/1.png",
    "url_private": "https://slack.com/files-pri/T024BE7LD-F024BERPE/1.png",
    "url_private_download": "https://slack.com/files-pri/T024BE7LD-F024BERPE/download/1.png",

    "thumb_64": "https://slack-files.com/files-tmb/T024BE7LD-F024BERPE-c66246/1_64.png",
    "thumb_80": "https://slack-files.com/files-tmb/T024BE7LD-F024BERPE-c66246/1_80.png",
    "thumb_360": "https://slack-files.com/files-tmb/T024BE7LD-F024BERPE-c66246/1_360.png",
    "thumb_360_gif": "https://slack-files.com/files-tmb/T024BE7LD-F024BERPE-c66246/1_360.gif",
    "thumb_360_w": "100",
    "thumb_360_h": "125",

    "permalink" : "https://tinyspeck.slack.com/files/cal/F024BERPE/1.png",
    "edit_link" : "https://tinyspeck.slack.com/files/cal/F024BERPE/1.png/edit",
    "preview" : "<!DOCTYPE html>\n<html>\n<meta charset='utf-8'>",
    "preview_highlight" : "<div class=\"sssh-code\"><div class=\"sssh-line\"><pre><!DOCTYPE html...",
    "lines" : 123,
    "lines_more": 118,

    "is_public": true,
    "public_url_shared": false,
    "channels": ["C024BE7LT"],
    "groups": ["G12345"],
    "num_stars": 7,
    "is_starred": true
}`

		Convey("When I json.Unmarshal into a File instance", func() {
			file := &File{}
			err := json.NewDecoder(strings.NewReader(text)).Decode(file)

			Convey("Then I expect no errors", func() {
				So(err, ShouldBeNil)
			})

			Convey("And I expect the values to be set correctly", func() {
				So(file.Id, ShouldEqual, "F2147483862")
				So(file.Created, ShouldEqual, 1356032811)
				So(file.Timestamp, ShouldEqual, 1356032811)
				So(file.Name, ShouldEqual, "file.htm")
				So(file.Title, ShouldEqual, "My HTML file")
				So(file.Mimetype, ShouldEqual, "text/plain")
				So(file.Filetype, ShouldEqual, "text")
				So(file.PrettyType, ShouldEqual, "Text")
				So(file.User, ShouldEqual, "U2147483697")
				So(file.Mode, ShouldEqual, "hosted")
				So(file.Editable, ShouldEqual, true)
				So(file.IsExternal, ShouldEqual, false)
				So(file.ExternalType, ShouldEqual, "")
				So(file.Size, ShouldEqual, 12345)

				So(file.Url, ShouldEqual, "https://slack-files.com/files-pub/T024BE7LD-F024BERPE-09acb6/1.png")
				So(file.UrlDownload, ShouldEqual, "https://slack-files.com/files-pub/T024BE7LD-F024BERPE-09acb6/download/1.png")
				So(file.UrlPrivate, ShouldEqual, "https://slack.com/files-pri/T024BE7LD-F024BERPE/1.png")
				So(file.UrlPrivateDownload, ShouldEqual, "https://slack.com/files-pri/T024BE7LD-F024BERPE/download/1.png")

				So(file.Thumb64, ShouldEqual, "https://slack-files.com/files-tmb/T024BE7LD-F024BERPE-c66246/1_64.png")
				So(file.Thumb80, ShouldEqual, "https://slack-files.com/files-tmb/T024BE7LD-F024BERPE-c66246/1_80.png")
				So(file.Thumb360, ShouldEqual, "https://slack-files.com/files-tmb/T024BE7LD-F024BERPE-c66246/1_360.png")
				So(file.Thumb360Gif, ShouldEqual, "https://slack-files.com/files-tmb/T024BE7LD-F024BERPE-c66246/1_360.gif")
				So(file.Thumb360W, ShouldEqual, "100")
				So(file.Thumb360H, ShouldEqual, "125")

				So(file.Permalink, ShouldEqual, "https://tinyspeck.slack.com/files/cal/F024BERPE/1.png")
				So(file.EditLink, ShouldEqual, "https://tinyspeck.slack.com/files/cal/F024BERPE/1.png/edit")
				So(file.Preview, ShouldEqual, "<!DOCTYPE html>\n<html>\n<meta charset='utf-8'>")
				So(file.PreviewHighlight, ShouldEqual, "<div class=\"sssh-code\"><div class=\"sssh-line\"><pre><!DOCTYPE html...")
				So(file.Lines, ShouldEqual, 123)
				So(file.LinesMore, ShouldEqual, 118)

				So(file.IsPublic, ShouldEqual, true)
				So(file.PublicUrlShared, ShouldEqual, false)
				So(file.Channels, ShouldResemble, []string{"C024BE7LT"})
				So(file.Groups, ShouldResemble, []string{"G12345"})
				So(file.NumStars, ShouldEqual, 7)
				So(file.IsStarred, ShouldEqual, true)
			})
		})
	})
}
