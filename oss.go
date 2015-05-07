package ali_oss

import (
	"io"
)

type OSS interface {
	PutObject(location, bucketName, key string, file io.Reader) (err error)
	GetObjectURL(location, bucketName, objectName string) (url string)
	GetStaticWidthObjectURL(domain, bucketName, objectName string, width int64) (URL string)
}
