package model


type MessageHead struct {
	FromMac string
	ToMac string
	FromIp string
	ToIp string
	FromPort int//发送方的端口
	IsArpReq bool//是否是arp请求
	IsArpRes bool//是否是arp响应
}
type Message struct {
	Head *MessageHead
	Body string
}
