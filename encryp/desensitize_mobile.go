package encryp

import (
	"strings"
)

func MosaicMobile(mobile string) string {
	if mobile == "" || len(mobile) == 0 {
		return ""
	}
	if len(mobile) <= 3 {
		return "***"
	}
	space := strings.TrimSpace(mobile)
	rs := []rune(space)
	//suffixStr := rs
	start := string(rs[0:3])
	end := string(rs[len(mobile)-4 : len(mobile)])
	var build strings.Builder
	build.WriteString(start)
	build.WriteString("****")
	build.WriteString(end)
	return build.String()
}
