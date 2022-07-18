package model

type ExchangeBoardPortM map[*ExchangeBoard]int
type PortSenderM  map[int]Sender

type Sender interface {
	SendMessage(message *Message)
	NewMessage(ip,msg string)*Message
}
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
