package ali_oss

import (
	"io"

	"github.com/gogap/ali_oss/parser"
)

type OSS interface {
	//上传图片
	PutObject(location, bucketName, key string, file io.Reader) (err error)
	//获取授权访问的对象地址
	GetObjectURL(location, bucketName, objectName string) (url string)
	//将对象按指定宽度进行等比例缩放
	GetStaticWidthObjectURL(domain, bucketName, objectName string, width int64) (URL string)
	//将对象按指定高度进行等比例缩放
	GetStaticHeightObjectURL(domain, bucketName, objectName string, height int64) (URL string)
	//指定对象的高度和宽度进行缩放
	GetDynamicObjectURL(domain, bucketName, objectName string, width, height int64) (URL string)
	//按比例缩放图片，proportion取值1~1000,100代表原图（即最大支持10倍放大）
	GetProportionObjectURL(domain, bucketName, objectName string, proportion int64) (URL string)
	//获取含有水印的授权对象访问地址
	GetObjectURLWithWatermark(domain, bucketName, objectName, watermark string) (URL string)
	//获取object list
	GetBucket(filter map[string]string, bucketName, location string) (result parser.ListBucketResult, err error)
}
