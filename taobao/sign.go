package taobao

import (
	"crypto/md5"
	"fmt"
	"sort"
	"strings"
)

func makeSign(params map[string]string) string {
	keys := make([]string, 0, len(params))
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	tmp := ""
	for _, k := range keys {
		tmp += fmt.Sprintf("%s%s", k, params[k])
	}
	bs := []byte(mSecret + tmp + mSecret)
	sign := strings.ToUpper(fmt.Sprintf("%x", md5.Sum(bs)))
	return sign
}
