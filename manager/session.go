package manager

import (
	"fmt"

	tool "github.com/fregie/gotool"
)

const (
	sessionIDLength = 32
)

type Session struct {
	ID            string
	Proto         string
	Index         string
	ConfigType    int32
	SendRateLimit uint64
	RecvRateLimit uint64
	CustomOption  string
}

func genSessionID(proto, index string) string {
	str1 := proto
	str2 := index
	if len(str1) > 8 {
		str1 = str1[:8]
	}
	if len(str2) > 8 {
		str2 = str2[:8]
	}
	return fmt.Sprintf("%s-%s-%s", str1, str2, string(tool.RandomString(sessionIDLength-len(proto)-len(index))))
}
