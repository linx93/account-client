package initialize

import (
	"fmt"
	"git.tuochexia.cn/tcxcn/common/accountcenter/rpcclient/account"
	"git.tuochexia.cn/tcxcn/common/accountcenter/rpcclient/approval"
	"git.tuochexia.cn/tcxcn/common/accountcenter/rpcclient/company"
	"git.tuochexia.cn/tcxcn/common/accountcenter/rpcclient/messager"
	"git.tuochexia.cn/tcxcn/common/accountcenter/rpcclient/todo"
	"git.tuochexia.cn/tcxcn/common/global"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"sync"
)

var once sync.Once

func initAccountGrpcConn(tenant []int64, clientConn *grpc.ClientConn) {

	once.Do(func() {
		global.SetAccountConn(clientConn)
	})

}

func initAccountGrpcConnByArgs(tenant []int64) error {

	once.Do(func() {
		global.SetAccountConn(clientConn)
	})

	//初始化所有的tenant的模板
	for _, t := range tenant {
		err := initTemplates(t)
		if err != nil {
			return err
		}
	}

	return nil
}

type tokenCredentials struct {
	appId  string
	secret string
}

func (c tokenCredentials) RequireTransportSecurity() bool {
	return false
}

func GetCentersService(cfg Config) (*CentersService, error) {
	var c = CentersService{}
	c.credentials = tokenCredentials{
		appId:  cfg.AppId,
		secret: cfg.Secret,
	}
	cc, err := grpc.Dial(fmt.Sprintf("%s:%d", cfg.Addr, cfg.Port), grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithPerRPCCredentials(c.credentials))
	if err != nil {
		return nil, err
	}
	c.conn = cc
	c.cfg = &cfg

	c.AccountClt = account.NewAccountServiceClient(cc)
	c.ApprovalClt = approval.NewApprovalServiceClient(cc)
	c.CompanyClt = company.NewCompanyServiceClient(cc)
	c.MessagerClt = messager.NewMessageServiceClient(cc)
	c.PaymentClt = payment.NewPaymentServiceClient(cc)
	c.TodoClt = todo.NewTodoServiceClient(cc)

	return &c, nil
}
