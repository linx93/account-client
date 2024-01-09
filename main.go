package main

/*func init() {
	target := fmt.Sprintf("%s:%s", "192.168.2.9", "18000")

	//无证书
	dial, err := grpc.Dial(target,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithPerRPCCredentials(new(global.TokenCredential)),
		//注入拦截器
		//grpc.WithUnaryInterceptor(preClientInterceptor(token)),
	)
	if err != nil {
		log.Fatal("连接服务端失败！", err)
	}

	initialize.InitAccountClientCoon(dial)
}

func main() {
	client := messager.MsgClient{}
	//获取消息
	msg, err := client.GetMsg(1, 1, 100, 13)
	if err != nil {
		log.Println(err)
	}


	log.Println(msg)
}*/
