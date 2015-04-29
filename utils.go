package ali_oss

import (
	"encoding/json"
	"fmt"
)

func pretty(v interface{}) {
	b, err := json.MarshalIndent(v, " ", "  ")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))

}
