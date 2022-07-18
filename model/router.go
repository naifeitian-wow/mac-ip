package model

//路由器
type Router struct {
	IpPort map[string]*NextPort//ip网段与端口映射表
	Arp map[string]string//arp 用于缓存ip到mac的关系
}
type NextPort struct {
	Next *Router//下一跳
	Port int//端口
}
