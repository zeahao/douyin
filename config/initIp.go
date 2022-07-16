package config

import (
	"net"
	"strings"
)

var URL string //储存视频链接的前缀url

// InitGetURL 获取服务端URL
func InitGetURL() {
	conn, err := net.Dial("udp", "8.8.8.8:53")
	if err != nil {
		panic(err)
	}
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	URL = "http://" + strings.Split(localAddr.String(), ":")[0]
}
