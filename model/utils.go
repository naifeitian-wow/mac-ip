package model

import (
	"strconv"
	"strings"
)
var SubNetM=map[string]string{
	"0":"0.0.0.0",
	"8":"255.0.0.0",
	"16":"255.255.0.0",
	"24":"255.255.255.0",
	"32":"255.255.255.255",
}
func ipAddrToInt(ipAddr string) int64 {//ip转int
	bits := strings.Split(ipAddr, ".")
	b0, _ := strconv.Atoi(bits[0])
	b1, _ := strconv.Atoi(bits[1])
	b2, _ := strconv.Atoi(bits[2])
	b3, _ := strconv.Atoi(bits[3])
	var sum int64
	sum += int64(b0) << 24
	sum += int64(b1) << 16
	sum += int64(b2) << 8
	sum += int64(b3)
	return sum
}
func isIpSubNet(ip1,ip2,subnet string)bool{
	myIp:=ipAddrToInt(ip1)
	ipInt:=ipAddrToInt(ip2)
	subnetMaskInt:=ipAddrToInt(subnet)
	return myIp&subnetMaskInt==ipInt&subnetMaskInt
}