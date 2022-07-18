package model

import "fmt"

//主机
type Computer struct {
	Ip string//ip
	Mac string//mac地址
	SubnetMask string//子网掩码
	DefaultGateway string//默认网关
	Arp map[string]string//arp 用于缓存ip到mac的关系
	ExchangeBoard *ExchangeBoard//所属的交换机
	Port int//所属的交换机的端口
	MsgCh chan *Message //监听消息
	MsgList []*Message//消息存储下来
}
func (c *Computer)NewMessage(ip,msg string)*Message{
	mac:=c.Arp[ip]
	return &Message{
		Head: &MessageHead{
			FromMac:  c.Mac,
			ToMac: mac,
			FromIp:   c.Ip,
			ToIp:     ip,
			FromPort: c.Port,
		},
		Body: msg,
	}
}
func (c *Computer)SendMessage(message *Message){//只知道目标ip
	c.ExchangeBoard.SendMessage(message)//通过交换机发送消息
}
func (c *Computer)Wait(){
	go func() {
		for{
			select {
			case v:=<-c.MsgCh:
				if v.Head.ToMac==""{//第一次，不知道我的mac，需要通过ip验证
					if v.Head.ToIp==c.Ip{
						if v.Head.IsArpReq{//arp请求
							c.Arp[v.Head.FromIp]=v.Head.FromMac//更新自己的arp表
							message:=&Message{
								Head: &MessageHead{
									FromMac: c.Mac,
									ToMac:   v.Head.FromMac,
									FromIp:  c.Ip,
									ToIp:    v.Head.FromIp,
									IsArpRes: true,
									FromPort: c.Port,
								},
							}
							c.MsgList=append(c.MsgList,v)//不确定要不要发送
							c.ExchangeBoard.SendMessage(message)//通过交换机响应
						}
					}
				} else if v.Head.ToMac==c.Mac{//收到消息，确实是自己的mac
					if v.Head.IsArpRes{
						c.Arp[v.Head.FromIp]=v.Head.FromMac//只更新自己的arp表即可
					}else{
						fmt.Println(fmt.Sprintf("[%s]电脑收到[%s]的消息啦：%s",c.Mac,v.Head.FromMac,v.Body))
						//正常请求，接收即可
						c.MsgList=append(c.MsgList,v)
					}
				}
			}
		}
	}()
}