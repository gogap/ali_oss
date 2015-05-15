package auth

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"net/http"
	"sort"
	"strings"
	"time"

	"github.com/gogap/ali_oss/constant"
)

type RequestSigner interface {
	Sign(method string, headers map[string]string, resource string, creds Credentials) (signature string, err error)
	HeaderSign(method string, expires int64, resource string, creds Credentials) string
}

func DefaultRequestSigner() RequestSigner {
	return &requestSigner{}
}

type requestSigner struct{}

func (p *requestSigner) Sign(method string, headers map[string]string, resource string, creds Credentials) (signature string, err error) {
	contentMD5 := ""
	contentType := ""
	date := time.Now().UTC().Format(http.TimeFormat)

	if v, exist := headers[constant.CONTENT_TYPE]; exist {
		contentMD5 = v
	}

	if v, exist := headers[constant.CONTENT_TYPE]; exist {
		contentType = v
	}

	if v, exist := headers[constant.DATE]; exist {
		date = v
	}

	ossHeaders := []string{}

	for k, v := range headers {
		if strings.HasPrefix(k, constant.OSS_PREFIX) {
			ossHeaders = append(ossHeaders, k+":"+strings.TrimSpace(v))
		}
	}
	xossHeader := strings.Join(ossHeaders, "\n")
	if len(ossHeaders) > 0 {
		xossHeader = xossHeader + "\n"
	}
	sort.Sort(sort.StringSlice(ossHeaders))

	stringToSign := method + "\n" +
		contentMD5 + "\n" +
		contentType + "\n" +
		date + "\n" +
		xossHeader +
		resource
	sha1Hash := hmac.New(sha1.New, []byte(creds.GetSecretAccessKey()))
	if _, err = sha1Hash.Write([]byte(stringToSign)); err != nil {
		return
	}
	signature = base64.StdEncoding.EncodeToString(sha1Hash.Sum(nil))
	return
}

func (p *requestSigner) HeaderSign(method string, expires int64, resource string, creds Credentials) string {
	stringToSign := fmt.Sprintf("%s\n\n\n%d\n%s", method, expires, resource)
	sha1Hash := hmac.New(sha1.New, []byte(creds.GetSecretAccessKey()))
	sha1Hash.Write([]byte(stringToSign))
	return base64.StdEncoding.EncodeToString(sha1Hash.Sum(nil))
}
