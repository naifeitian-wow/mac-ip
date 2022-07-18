package main

import (
	"mac-ip/model"
	"time"
)
var exchangeBoard1,exchangeBoard2,exchangeBoard3 *model.ExchangeBoard
var computerM1=map[int]*model.Computer{
	0:{
		Ip:             "192.168.0.1",
		Mac:            "AAAA",
		SubnetMask:     "255.255.255.0",
		DefaultGateway: "192.168.0.254",
		ExchangeBoard:exchangeBoard1,
		Port: 0,
	},1:{
		Ip:             "192.168.0.2",
		Mac:            "BBBB",
		SubnetMask:     "255.255.255.0",
		DefaultGateway: "192.168.0.254",
		ExchangeBoard: exchangeBoard1,
		Port: 1,
	},
}
var computerM2=map[int]*model.Computer{
	0:{
		Ip:             "192.168.1.1",
		Mac:            "CCCC",
		SubnetMask:     "255.255.255.0",
		DefaultGateway: "192.168.1.254",
		ExchangeBoard: exchangeBoard2,
		Port: 0,
	},1:{
		Ip:             "192.168.1.2",
		Mac:            "DDDD",
		SubnetMask:     "255.255.255.0",
		DefaultGateway: "192.168.1.254",
		ExchangeBoard: exchangeBoard2,
		Port: 1,
	}}
var computerM3=map[int]*model.Computer{
	0:{
		Ip:             "192.168.2.1",
		Mac:            "EEEE",
		SubnetMask:     "255.255.255.0",
		DefaultGateway: "192.168.2.254",
		ExchangeBoard: exchangeBoard3,
		Port: 0,
	},1:{
		Ip:             "192.168.2.2",
		Mac:            "FFFF",
		SubnetMask:     "255.255.255.0",
		DefaultGateway: "192.168.2.254",
		ExchangeBoard: exchangeBoard3,
		Port: 1,
	}}
func init(){

	exchangeBoard1=model.NewExchangeBoard(computerM1, nil)
	exchangeBoard2=model.NewExchangeBoard(computerM2,nil)
	exchangeBoard3=model.NewExchangeBoard(computerM3,nil)
	initComputer(exchangeBoard1,exchangeBoard2,exchangeBoard3)
	exchangeBoard1.PortList[2]=exchangeBoard2
	exchangeBoard1.Port=2
	exchangeBoard2.PortList[2]=exchangeBoard1
	exchangeBoard2.Port=2
	wait(exchangeBoard1,exchangeBoard2,exchangeBoard3)//监听消息
}
func main(){
	computerM1[0].SendMessage("192.168.0.2","arp")
	computerM1[0].SendMessage("192.168.1.1","arp")
	time.Sleep(1*time.Second)
	computerM1[0].SendMessage("192.168.0.2","hello")
	computerM1[1].SendMessage("192.168.0.1","hello2")
	computerM1[1].SendMessage("192.168.0.1","hello3")
	computerM1[0].SendMessage("192.168.1.1","hello4")
	computerM2[0].SendMessage("192.168.0.1","hello5")
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
				computer.MsgCh=make(chan *model.Message,20)
				computer.Arp=map[string]string{}
			}
		}
	}
}