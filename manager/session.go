package manager

import (
	"fmt"

	svcpb "github.com/fregie/simple/proto/simple-interface"

	tool "github.com/fregie/gotool"
)

const (
	sessionIDLength = 32
)

type Session struct {
	ID            string `gorm:"primary_key"`
	Name          string `gorm:"type:varchar(1024);unique;index"`
	Proto         string
	Index         string
	ConfigType    int32
	Config        string
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
	return fmt.Sprintf("%s-%s-%s", str1, str2, string(tool.RandomString(sessionIDLength-len(str1)-len(str2))))
}

func (s *Session) convertOption() *svcpb.Option {
	return &svcpb.Option{
		SendRateLimit: s.SendRateLimit,
		RecvRateLimit: s.RecvRateLimit,
	}
}
