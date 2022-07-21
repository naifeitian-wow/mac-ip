package model

import (
	"strings"
)

//路由器
type Router struct {
	IpPort IpPort//ip网段与端口映射表  路由表 分为静态和动态
	Arp map[string]string//arp 用于缓存ip到mac的关系
	PortExchangeBoardM PortExchangeBoardM //路由器端口列表
	PortMacM PortMacM//保存端口和mac,ip的关系 初始化就有
}
type NextPort struct {
	Next *Router//下一跳
	Port int//端口
}
type MacAndIp struct {
	Mac string
	Ip string
}

func (r *Router)SendMessage(message Message){
	if message.Head.IsArpReq || message.Head.IsArpRes{
		r.Arp[message.Head.FromIp]=message.Head.FromMac
	}
	port:=r.GetPortFromIp(message.Head.ToIp)//根据路由表找到ip对应的端口
	message.Head.FromMac=r.PortMacM[port].Mac//来源mac置为路由器发送端口的mac
	message.Head.ToMac=r.Arp[message.Head.ToIp]//需要根据arp得到ip对应的mac
	message.Head.IsFromRouter=true
	message.Head.FromPort=-1//-1代表在交换机上 路由器没有插在正常的lan口上，而是wan口，这里用-1表示
	r.PortExchangeBoardM[port].SendMessage(message)//向指定端口发送消息
}

func (r *Router)GetPortFromIp(ip string)int{
	for k,v:=range r.IpPort{
		ipList:=strings.Split(k,"/")
		if len(ipList)==0{
			panic("error")
		}
		if isIpSubNet(ip,ipList[0],SubNetM[ipList[1]]){
			if v.Next!=nil{
				return v.Next.GetPortFromIp(ip)
			}
			return v.Port
		}
	}
	panic("数据不可达")
}