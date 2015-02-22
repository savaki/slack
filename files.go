package slack

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"strings"

	"github.com/savaki/slack/form"
)

type FilesUploadReq struct {
	Content        io.Reader
	Filetype       string `form:"filetype"`
	Filename       string `form:"filename"`
	Title          string `form:"title"`
	InitialComment string `form:"initial_comment"`
	Channels       []string
}

type FilesUploadResp struct {
	Ok    bool   `json:"ok"`
	Error string `json:"error,omitempty"`
	File  *File  `json:"file,omitempty"`
}

func (s *Client) FilesUpload(req *FilesUploadReq) (*FilesUploadResp, error) {
	if req.Content != nil {
		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)
		part, err := writer.CreateFormFile("file", req.Filename)
		if err != nil {
			return nil, err
		}

		// copy the data elements over
		_, err = io.Copy(part, req.Content)
		if err != nil {
			return nil, err
		}

		// copy the parameters
		params := form.AsValues(req)
		for key, value := range params {
			writer.WriteField(key, strings.Join(value, ","))
		}
		writer.WriteField("channels", strings.Join(req.Channels, ","))

		// close the writer
		err = writer.Close()
		if err != nil {
			return nil, err
		}
		resp := &FilesUploadResp{}
		contentType := "multipart/form-data; boundary=" + writer.Boundary()
		err = s.post("files.upload", contentType, body, resp)
		return resp, err

	} else {
		return nil, fmt.Errorf("Content must be specified")
	}
}
