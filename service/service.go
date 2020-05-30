package service
//服务接口
type Service interface {
	Name()string
	Init()
	Start()error
	Stop()error
}