package ali_oss

import (
	"fmt"
	"io"
	"net/http"
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

func (p *client) getSignature(bucketName, objectName string) string {
	resource := fmt.Sprintf("/%s/%s", bucketName, objectName)
	return p.signer.HeaderSign(constant.GET, defaultExpires(), resource, p.creds)
}

func (p *client) GetObjectURL(location, bucketName, objectName string) (URL string) {
	signature := p.getSignature(bucketName, objectName)
	return fmt.Sprintf(constant.TPL_OBJECT_URL, bucketName, location, objectName, defaultExpires(), p.creds.GetAccessKeyId(), urlEncode(signature))
}

func (p *client) GetObjectURLWithWatermark(domain, bucketName, objectName, watermark string) (URL string) {
	//watermarke：不能以中文开头，避免使用负担，默认在前面加一个空格
	watermark = " " + watermark
	signature := p.getSignature(bucketName, fmt.Sprintf("%s@watermark=2&text=%s", objectName, base64String(watermark)))
	return fmt.Sprintf(constant.TPL_OBJECT_WITH_WATERMARK_URL, trimDomain(domain), fmt.Sprintf("%s@watermark=2&text=%s", objectName, base64String(watermark)), defaultExpires(), p.creds.GetAccessKeyId(), urlEncode(signature))
}

func (p *client) GetStaticWidthObjectURL(domain, bucketName, objectName string, width int64) (URL string) {
	signature := p.getSignature(bucketName, fmt.Sprintf(constant.TPL_STATIC_WIDTH_OBJECT, objectName, width))
	return fmt.Sprintf(constant.TPL_STATIC_WIDTH_OBJECT_URL, trimDomain(domain), objectName, urlEncode("@"), width, defaultExpires(), p.creds.GetAccessKeyId(), urlEncode(signature))
}

func (p *client) GetStaticHeightObjectURL(domain, bucketName, objectName string, height int64) (URL string) {
	signature := p.getSignature(bucketName, fmt.Sprintf(constant.TPL_STATIC_HEIGHT_OBJECT, objectName, height))
	return fmt.Sprintf(constant.TPL_STATIC_HEIGHT_OBJECT_URL, trimDomain(domain), objectName, urlEncode("@"), height, defaultExpires(), p.creds.GetAccessKeyId(), urlEncode(signature))
}

func (p *client) GetDynamicObjectURL(domain, bucketName, objectName string, width, height int64) (URL string) {
	signature := p.getSignature(bucketName, fmt.Sprintf(constant.TPL_DYNAMIC_OBJEC, objectName, width, height))
	return fmt.Sprintf(constant.TPL_DYNAMIC_OBJECT_URL, trimDomain(domain), objectName, urlEncode("@"), width, height, defaultExpires(), p.creds.GetAccessKeyId(), urlEncode(signature))
}

func (p *client) GetProportionObjectURL(domain, bucketName, objectName string, proportion int64) (URL string) {
	signature := p.getSignature(bucketName, fmt.Sprintf(constant.TPL_PROPORTION_OBJECT, objectName, proportion))
	return fmt.Sprintf(constant.TPL_PROPORTION_OBJECT_URL, trimDomain(domain), objectName, urlEncode("@"), proportion, defaultExpires(), p.creds.GetAccessKeyId(), urlEncode(signature))
}
