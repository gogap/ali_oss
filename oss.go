package ali_oss

import (
	"io"
)

type OSS interface {
	//PutObject
	//bucketName,Bucket名称
	//key,object的key
	//file,要上传的文件
	PutObject(location, bucketName, key string, file io.Reader) (err error)

	GetObjectURL(location, bucketName, objectName string) (url string)
}
