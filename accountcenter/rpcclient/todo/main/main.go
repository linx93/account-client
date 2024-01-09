package main

import (
	"context"
	"github.com/tcxcn/vehicle_manager_server/core"
	"github.com/tcxcn/vehicle_manager_server/global"
	"github.com/tcxcn/vehicle_manager_server/initialize"
	"github.com/tcxcn/vehicle_manager_server/service/accountcenter/rpcclient/todo"
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

// 测试消息
func main() {

	wg.Add(1)
	init_()
	wg.Wait()

	conn, err := global.GetAccountCenterConn()
	if err != nil {
		log.Println(err)
		return
	}

	//创建代办
	//createTodo(conn, 1, 13, 1704964689, "linx-test-todo", "www.google.com")

	//获取代办
	//getTodos(conn, 1, 13)

	//分页获取代办
	//pageTodos(conn, 1, 100, 13)

	//提交备注
	//pushRemark(conn, 1, 13, "linx push a remark")

	//获取备注
	getRemarks(conn, 1, 13)

}

func createTodo(conn *grpc.ClientConn, user, tenant, deadLine int64, title, link string) {
	request := &todo.CreateTodoRequest{
		Title:    title,
		Link:     link,
		User:     user,
		Tenant:   tenant,
		DeadLine: deadLine,
	}

	todoRes, err := todo.NewTodoServiceClient(conn).CreateTodo(context.Background(), request)
	if err != nil {
		log.Println(err)
		return
	}
	if !todoRes.Ok {
		log.Println(todoRes.Message)
		return
	}

	log.Println(todoRes)
}

func getTodos(conn *grpc.ClientConn, user, tenant int64) {
	todosRes, err := todo.NewTodoServiceClient(conn).GetTodos(context.Background(), &todo.GetTodosRequest{User: user, Tenant: tenant})
	if err != nil {
		log.Println(err)
		return
	}
	if !todosRes.Ok {
		log.Println(todosRes.Message)
		return
	}

	log.Println(todosRes)
}

func pageTodos(conn *grpc.ClientConn, page, size, tenant int64) {
	pageRes, err := todo.NewTodoServiceClient(conn).PageTodos(context.Background(), &todo.PageTodosRequest{Page: page, Size: size, Tenant: tenant})
	if err != nil {
		log.Println(err)
		return
	}
	if !pageRes.Ok {
		log.Println(pageRes.Message)
		return
	}

	log.Println(pageRes)
}

func pushRemark(conn *grpc.ClientConn, todoId, tenant int64, remark string) {
	pushRemarkRes, err := todo.NewTodoServiceClient(conn).PushRemark(context.Background(), &todo.PushRemarkRequest{Todo: todoId, Tenant: tenant, Remark: remark})
	if err != nil {
		log.Println(err)
		return
	}
	if !pushRemarkRes.Ok {
		log.Println(pushRemarkRes.Message)
		return
	}

	log.Println(pushRemarkRes)
}

func getRemarks(conn *grpc.ClientConn, todoId, tenant int64) {
	remarkRes, err := todo.NewTodoServiceClient(conn).GetRemarks(context.Background(), &todo.GetRemarksRequest{Todo: todoId, Tenant: tenant})
	if err != nil {
		log.Println(err)
		return
	}
	if !remarkRes.Ok {
		log.Println(remarkRes.Message)
		return
	}

	log.Println(remarkRes)
}
