package ali_oss

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/gogap/ali_oss/auth"
	"github.com/gogap/ali_oss/comm"
	"github.com/gogap/ali_oss/constant"
)

type client struct {
	creds     auth.Credentials
	signer    auth.RequestSigner
	requester comm.Requester
}

func NewClient(accessKeyId, secretAccessKey string) OSS {
	cli := client{}
	cli.creds = auth.DefaultCredentialsProvider(accessKeyId, secretAccessKey).GetCredentials()
	cli.signer = auth.DefaultRequestSigner()
	cli.requester = comm.DefaultRequester()
	return &cli
}

func (p *client) PutObject(location, bucketName, objectName string, file io.Reader) (err error) {
	header := make(map[string]string)
	target := fmt.Sprintf("http://%s.oss-cn-%s.aliyuncs.com/%s", bucketName, location, objectName)

	date := time.Now().UTC().Format(http.TimeFormat)
	header[constant.DATE] = date

	resource := fmt.Sprintf("/%s/%s", bucketName, objectName)
	signature, err := p.signer.Sign(constant.PUT, header, resource, p.creds)
	if err != nil {
		return
	}
	header[constant.AUTHORIZATION] = fmt.Sprintf("%s %s:%s", constant.OSS, p.creds.GetAccessKeyId(), signature)
	err = p.requester.Request(constant.PUT, target, header, file, p.creds)
	return
}

func (p *client) GetObjectURL(location, bucketName, objectName string) (URL string) {
	resource := fmt.Sprintf("/%s/%s", bucketName, objectName)
	signature := p.signer.HeaderSign(constant.GET, constant.EXPIRES+time.Now().Unix(), resource, p.creds)
	rawURL := fmt.Sprintf("http://%s.oss-cn-%s.aliyuncs.com/%s?Expires=%d&OSSAccessKeyId=%s&Signature=%s",
		bucketName, location, objectName, constant.EXPIRES+time.Now().Unix(), p.creds.GetAccessKeyId(), signature)
	u, _ := url.Parse(rawURL)
	return u.String()
}

func (p *client) GetStaticWidthObjectURL(location, bucketName, objectName string, width int64) (URL string) {
	// resource := fmt.Sprintf("/%s/%s", bucketName, objectName)
	// signature := p.signer.HeaderSign(constant.GET, constant.EXPIRES+time.Now().Unix(), resource, p.creds)
	// rawURL := fmt.Sprintf("http://%s.oss-cn-%s.aliyuncs.com/%s?Expires=%d&OSSAccessKeyId=%s&Signature=%s",
	// 	bucketName, location, objectName, constant.EXPIRES+time.Now().Unix(), p.creds.GetAccessKeyId(), signature)

	return
}

func (p *client) GetStaticHeightObjectURL(location, bucketName, objectName string, height int64) (URL string) {
	return
}

func (p *client) GetDynamicObjectURL(location, bucketName, objectName string, width, height int64) (URL string) {
	return
}

func (p *client) GetProportionObjectURL(location, bucketName, objectName string, proportion int64) (URL string) {
	return
}
