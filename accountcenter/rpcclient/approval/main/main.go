package main

import (
	"context"
	"github.com/tcxcn/vehicle_manager_server/core"
	"github.com/tcxcn/vehicle_manager_server/global"
	"github.com/tcxcn/vehicle_manager_server/initialize"
	"github.com/tcxcn/vehicle_manager_server/service/accountcenter/rpcclient/approval"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"log"
	"sync"
)

var wg = sync.WaitGroup{}

func init_() {
	global.GVA_VP = core.Viper() // 初始化Viper
	initialize.OtherInit()

	global.GVA_LOG = core.Zap() // 初始化zap日志库
	zap.ReplaceGlobals(global.GVA_LOG)

	global.GVA_DB = initialize.Gorm() // gorm连接数据库

	//初始化grpcServer
	go initialize.InitGrpcServer(&wg)
}

func main() {
	wg.Add(1)
	init_()
	wg.Wait()

	conn, err := global.GetAccountCenterConn()
	if err != nil {
		log.Println(err)
		return
	}

	//创建配置
	//CreateApprovalConfig(conn, "请假")

	//获取审批配置
	//GetApprovalConfig(conn, "请假")

	//获取应用租户的所有审批配置
	//GetApprovalConfigs(conn)

	//创建实例
	//CreateApproval(conn)

	//审核通过
	//ApproveApproval(conn, 2, "muyumei")

	//审核不通过
	//DenyApproval(conn, 2, "chenying")

	//获取我的待审批
	//GetMyApproval(conn, "zhoujiajia")

	//分页获取我的审批记录
	//PageMyApproval(conn, "wuhongzhen")

	//获取审核实例的审核序列
	GetApprovalSequences(conn, 2)
}

// CreateApprovalConfig 创建审批配置
func CreateApprovalConfig(conn *grpc.ClientConn, business string) {

	steps := []*approval.Step{
		{Step: 0, Name: "周佳佳", Username: "zhoujiajia"},
		{Step: 1, Name: "陈颖", Username: "chenying"},
		{Step: 2, Name: "吴宏珍", Username: "wuhongzhen"},
		{Step: 3, Name: "穆玉梅", Username: "muyumei"},
	}

	//周佳佳,zhoujiajia
	//袁华,yuanhua
	//杨富明,yangfuming
	//吴宏珍,wuhongzhen
	//四川拖车侠,scAdmin
	//穆玉梅,muyumei
	//贵州拖车侠,gzAdmin
	//陈颖,chenying

	ccs := []*approval.CC{
		{Name: "陈颖", Username: "chenying"},
		{Name: "吴宏珍", Username: "wuhongzhen"},
	}

	request := &approval.CreateConfigRequest{
		Tenant:   13,
		Business: business,
		Steps:    steps,
		CC:       ccs,
	}
	config, err := approval.NewApprovalServiceClient(conn).CreateApprovalConfig(context.Background(), request)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println(config)
}

// GetApprovalConfig 获取审批配置
func GetApprovalConfig(conn *grpc.ClientConn, business string) {
	request := &approval.GetConfigRequest{
		Tenant:   13,
		Business: business,
	}
	config, err := approval.NewApprovalServiceClient(conn).GetApprovalConfig(context.Background(), request)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println(config)
}

// GetApprovalConfigs 获取应用租户的所有审批配置
func GetApprovalConfigs(conn *grpc.ClientConn) {
	request := &approval.GetConfigsRequest{
		Tenant: 13,
	}
	config, err := approval.NewApprovalServiceClient(conn).GetApprovalConfigs(context.Background(), request)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println(config)
}

// CreateApproval 创建审批实例
func CreateApproval(conn *grpc.ClientConn) {

	res, err := approval.NewApprovalServiceClient(conn).CreateApproval(context.Background(),
		&approval.CreateApprovalRequest{
			Tenant:            13,
			Business:          "请假",
			ApplicantUsername: "chenying",
		})
	if err != nil {
		log.Println(err)
		return
	}

	log.Println(res)
}

// ApproveApproval 审批通过
func ApproveApproval(conn *grpc.ClientConn, approvalId int64, username string) {
	res, err := approval.NewApprovalServiceClient(conn).ApproveApproval(context.Background(), &approval.ApproveRequest{Tenant: 13, Approval: approvalId, Username: username})
	if err != nil {
		log.Println(err)
		return
	}

	log.Println(res)
}

// DenyApproval 审批不通过
func DenyApproval(conn *grpc.ClientConn, approvalId int64, username string) {
	res, err := approval.NewApprovalServiceClient(conn).DenyApproval(
		context.Background(),
		&approval.DenyRequest{
			Tenant:   13,
			Approval: approvalId,
			Username: username,
			Remark:   "测试测试不通过",
		})
	if err != nil {
		log.Println(err)
		return
	}

	log.Println(res)
}

// GetMyApproval 获取我的待审批
func GetMyApproval(conn *grpc.ClientConn, username string) {
	res, err := approval.NewApprovalServiceClient(conn).GetMyApproval(context.Background(), &approval.GetMyApprovalRequest{Tenant: 13, Username: username})
	if err != nil {
		log.Println(err)
		return
	}

	log.Println(res)
}

// PageMyApproval 分页获取我的审批记录
func PageMyApproval(conn *grpc.ClientConn, username string) {
	request := &approval.PageMyApprovalRequest{
		Tenant:   13,
		Username: username,
		Page:     1,
		Size:     100,
	}

	res, err := approval.NewApprovalServiceClient(conn).PageMyApproval(context.Background(), request)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println(res)
}

// GetApprovalSequences 获取审核实例的审核序列
func GetApprovalSequences(conn *grpc.ClientConn, approvalId int64) {
	res, err := approval.NewApprovalServiceClient(conn).GetApprovalSequences(context.Background(), &approval.GetSequencesRequest{Tenant: 13, Approval: approvalId})
	if err != nil {
		log.Println(err)
		return
	}

	log.Println(res)
}
