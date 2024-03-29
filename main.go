package main

import (
	"mac-ip/model"
	"time"
)
var exchangeBoard1,exchangeBoard2,exchangeBoard3 *model.ExchangeBoard
var senderM1=model.PortSenderM{
	0:&model.Computer{
		Ip:             "192.168.0.1",
		Mac:            "AAAA",
		SubnetMask:     "255.255.255.0",
		DefaultGateway: "192.168.0.254",
		ExchangeBoard:exchangeBoard1,
		Port: 0,
	},1:&model.Computer{
		Ip:             "192.168.0.2",
		Mac:            "BBBB",
		SubnetMask:     "255.255.255.0",
		DefaultGateway: "192.168.0.254",
		ExchangeBoard: exchangeBoard1,
		Port: 1,
	},
}
var senderM2=model.PortSenderM{
	0:&model.Computer{
		Ip:             "192.168.1.1",
		Mac:            "CCCC",
		SubnetMask:     "255.255.255.0",
		DefaultGateway: "192.168.1.254",
		ExchangeBoard: exchangeBoard2,
		Port: 0,
	},1:&model.Computer{
		Ip:             "192.168.1.2",
		Mac:            "DDDD",
		SubnetMask:     "255.255.255.0",
		DefaultGateway: "192.168.1.254",
		ExchangeBoard: exchangeBoard2,
		Port: 1,
	}}
var senderM3=model.PortSenderM{
	0:&model.Computer{
		Ip:             "192.168.2.1",
		Mac:            "EEEE",
		SubnetMask:     "255.255.255.0",
		DefaultGateway: "192.168.2.254",
		ExchangeBoard: exchangeBoard3,
		Port: 0,
	},1:&model.Computer{
		Ip:             "192.168.2.2",
		Mac:            "FFFF",
		SubnetMask:     "255.255.255.0",
		DefaultGateway: "192.168.2.254",
		ExchangeBoard: exchangeBoard3,
		Port: 1,
	}}
var router=&model.Router{
	IpPort:             model.IpPort{
		"192.168.0.0/24":&model.NextPort{
			Next: nil,
			Port: 0,
		},
		"192.168.1.0/24":&model.NextPort{
			Next: nil,
			Port: 1,
		},
	},
	Arp:                map[string]string{},
	PortExchangeBoardM: nil,
	PortMacM:           model.PortMacM{
		0:model.MacAndIp{
			Mac:        "ABAB",
			Ip:         "192.168.0.254",
		},1:model.MacAndIp{
			Mac:        "CDCD",
			Ip:         "192.168.1.254",
		},
	},
}
func init(){
	exchangeBoard1,exchangeBoard2,exchangeBoard3=model.NewExchangeBoard(),model.NewExchangeBoard(),model.NewExchangeBoard()
	exchangeBoard1.SetSender(senderM1)
	exchangeBoard2.SetSender(senderM2)
	exchangeBoard3.SetSender(senderM3)
	router.PortExchangeBoardM=model.PortExchangeBoardM{
		0:exchangeBoard1,
		1:exchangeBoard2,
	}

	exchangeBoard1.Router=router
	exchangeBoard1.Port=0
	exchangeBoard2.Router=router
	exchangeBoard2.Port=1
	initComputer(exchangeBoard1,exchangeBoard2,exchangeBoard3)

	wait(exchangeBoard1,exchangeBoard2,exchangeBoard3)//监听消息
}
func main(){
	computer,_,computer3,_:=exchangeBoard1.GetSender(0),exchangeBoard1.GetSender(1),exchangeBoard2.GetSender(0),exchangeBoard3.GetSender(0)
	computer.SendMessage(computer.NewMessage("192.168.0.2","arp"))
	computer.SendMessage(computer.NewMessage("192.168.1.1","arp"))
	//computer.SendMessage(computer.NewMessage("192.168.2.1","arp"))
	time.Sleep(1*time.Second)
	computer.SendMessage(computer.NewMessage("192.168.0.2","hello"))
	//computer.SendMessage(computer.NewMessage("192.168.0.2","hello2"))
	//computer2.SendMessage(computer2.NewMessage("192.168.0.1","hello2"))
	computer.SendMessage(computer.NewMessage("192.168.1.1","hello3"))
	computer3.SendMessage(computer3.NewMessage("192.168.0.1","hello4"))
	for{

	}
}
func wait(ExchangeBoardList ...*model.ExchangeBoard){
	for i,_:=range ExchangeBoardList{
		go func(i int) {
			for _,v:=range ExchangeBoardList[i].PortList{
				if _,ok:=v.(*model.Computer);ok{
					v.(*model.Computer).Wait()
				}
			}
		}(i)
	}
}
func initComputer(exchangeBoardList ...*model.ExchangeBoard){
	for _,exchangeBoard:=range exchangeBoardList{
		for _,v:=range exchangeBoard.PortList{
			if computer,ok:=v.(*model.Computer);ok{
				computer.ExchangeBoard=exchangeBoard
				computer.MsgCh=make(chan model.Message,20)
				computer.Arp=map[string]string{}
			}
		}
	}
}