package initialize

import (
	"git.tuochexia.cn/tcxcn/common/global"
	"google.golang.org/grpc"
	"sync"
)

var once sync.Once

func InitAccountClientCoon(clientConn *grpc.ClientConn) {
	once.Do(func() {
		global.SetAccountConn(clientConn)
	})
}
