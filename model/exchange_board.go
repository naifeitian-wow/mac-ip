package model

import "fmt"

//交换机
type ExchangeBoard struct {
	MacPort map[string]int//mac与端口映射表
	PortList map[int]interface{}//所有的端口 端口后可能是电脑，也可能是交换机
	Port int//交换机也可能有插到别的交换机的端口
}
func NewExchangeBoard(computerM map[int]*Computer,exchangeBoardM map[int]*ExchangeBoard)*ExchangeBoard{
	m:=map[int]interface{}{}
	for k,v:=range computerM{
		m[k]=v
	}
	for k,v:=range exchangeBoardM{
		m[k]=v
	}
	exchangeBoard:=&ExchangeBoard{
		MacPort:  map[string]int{},
		PortList: m,
	}
	return exchangeBoard
}
func (e *ExchangeBoard)SendMessage(message *Message){
	e.MacPort[message.Head.FromMac]=message.Head.FromPort//更新mac与端口映射表，每次都更新，防止机器换端口
	if message.Head.ToMac==""{//目标mac为空，发送广播
		for k,v:=range e.PortList{
			if k==message.Head.FromPort{//发送方端口不对其发送消息
				continue
			}
			message.Head.IsArpReq=true
			e.commonSendMessage(v,message)
		}
	}else{//有mac
		if port,ok:=e.MacPort[message.Head.ToMac];ok{
			e.commonSendMessage(e.PortList[port],message)
		}else{//广播，更新mac地址表，不存在了
			fmt.Println("?")
		}
	}
}
//给指定端口机器发送消息
func (e *ExchangeBoard)commonSendMessage(v interface{},message *Message){
	switch v.(type) {
	case *Computer:
		v.(*Computer).MsgCh<-message
	case *ExchangeBoard:
		message.Head.FromPort=e.Port//交换机的的fromport需要替换
		v.(*ExchangeBoard).SendMessage(message)
	}
}
