package util

import (
	"fmt"
	"math/rand"
	"net"
)

// 生成随机 mac 地址
func GenerateMac() string {
	mac := make([]byte, 6)
	rand.Read(mac)
	mac[0] |= 2 // 设置第一个字节的第二位为1，表示为Locally Administered Address

	// 将MAC地址转换为字符串
	macAddr := net.HardwareAddr(mac)
	macString := macAddr.String()
	fmt.Println("随机MAC地址为：", macString)
	return macString
}
