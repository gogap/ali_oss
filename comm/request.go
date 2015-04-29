package comm

import (
	"encoding/xml"
	"errors"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/gogap/ali_oss/auth"
	"github.com/gogap/ali_oss/parser"
)

type Requester interface {
	Request(method, target string, header map[string]string, content io.Reader, creds auth.Credentials) (err error)
}

func DefaultRequester() Requester {
	return &request{}
}

type request struct{}

func (p *request) Request(method, target string, header map[string]string, content io.Reader, creds auth.Credentials) (err error) {
	client := &http.Client{}
	req, err := http.NewRequest(method, target, content)
	if err != nil {
		return
	}
	for k, v := range header {
		req.Header.Add(k, v)
	}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	if resp.StatusCode != 200 {
		var data []byte

		data, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			return
		}
		var result parser.ErrorMessageResponse
		err = xml.Unmarshal(data, &result)
		if err != nil {
			return
		}
		return errors.New(result.Code + ":" + result.Message)
	}
	return
}
