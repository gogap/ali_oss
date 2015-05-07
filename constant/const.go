package constant

const (
	AUTHORIZATION = "Authorization"
	CONTENT_TYPE  = "Content-Type"
	CONTENT_MD5   = "Content-MD5"
	OSS_VERSION   = "x-oss-version"
	HOST          = "Host"
	DATE          = "Date"
	KEEP_ALIVE    = "Keep-Alive"
	PUT           = "PUT"
	GET           = "GET"
	OSS           = "OSS"
	OSS_PREFIX    = "x-oss-"
)

const (
	EXPIRES = 86400 * 3 // 3 day
)

const (
	MAX_IMAGE_WIDTH  = 4096
	MAX_IMAGE_HEIGHT = 4096
)

const (
	TPL_OBJECT_URL               = "http://%s.oss-cn-%s.aliyuncs.com/%s?Expires=%d&OSSAccessKeyId=%s&Signature=%s"
	TPL_STATIC_WIDTH_OBJECT_URL  = "http://%s/%s%s%dw.jpg?Expires=%d&OSSAccessKeyId=%s&Signature=%s"
	TPL_STATIC_WIDTH_OBJECT      = "%s@%dw.jpg"
	TPL_STATIC_HEIGHT_OBJECT_URL = "http://%s.oss-cn-%s.aliyuncs.com/%s@%dh.jpg?Expires=%d&OSSAccessKeyId=%s&Signature=%s"
	TPL_DYNAMIC_OBJECT_URL       = "http://%s.oss-cn-%s.aliyuncs.com/%s@%dw_%dh.jpg?Expires=%d&OSSAccessKeyId=%s&Signature=%s"
	TPL_PROPORTION_OBJECT_URL    = "http://%s.oss-cn-%s.aliyuncs.com/%s@%dp.jpg?Expires=%d&OSSAccessKeyId=%s&Signature=%s"
)
