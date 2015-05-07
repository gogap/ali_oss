package ali_oss

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/gogap/ali_oss/constant"
)

func pretty(v interface{}) {
	b, err := json.MarshalIndent(v, " ", "  ")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))

}

func defaultExpires() int64 {
	return time.Now().Unix() + constant.EXPIRES
}

func urlEncode(s string) string {
	return url.QueryEscape(s)
}

func trimDomain(domain string) string {
	return strings.TrimSuffix(strings.TrimSuffix(domain, " "), "/")
}
