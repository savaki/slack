package slack

import (
	"log"
	"os"
	"strings"

	. "github.com/smartystreets/goconvey/convey"

	"testing"
)

func TestFilesUpload(t *testing.T) {
	Convey("Given file contents to upload", t, func() {
		content := strings.NewReader("<div><strong>hello</strong> world</div>")

		Convey("When I upload the contents to slack", func() {
			token := os.Getenv("SLACK_TOKEN")
			if token == "" {
				log.Println("SLACK_TOKEN not set.  skipping test")
				t.Skipped()
				return
			}

			channel := os.Getenv("SLACK_TEST_CHANNEL")
			if channel == "" {
				log.Println("SLACK_TEST_CHANNEL not set.  skipping test")
				t.Skipped()
				return
			}

			s := New(token)
			req := &FilesUploadReq{
				Content:  content,
				Filetype: "text/html",
				Filename: "sample.html",
				Title:    "Sample",
				Channels: []string{channel},
			}
			resp, err := s.FilesUpload(req)

			So(err, ShouldBeNil)
			So(resp, ShouldNotBeNil)
			So(resp.Ok, ShouldBeTrue)
		})
	})
}
