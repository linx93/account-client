package messager

import (
	"context"
	"fmt"
	"git.tuochexia.cn/tcxcn/common/global"
	"log"
)

type MsgClient struct{}

// CreateTemplate 创建模板
// tenant 租户ID      必传		例如：13
// title  模板的标题   必传		例如："自我介绍"
// link   链接模板内容 必传  		例如："www.baidu.com?id={id}&name={name}"
// tempStr消息模板内容 必传  		例如："我叫{name} ,今年{age}岁 ,我家住{address}"
func (MsgClient) CreateTemplate(tenant int64, title string, link string, tempStr string) (*Template, error) {
	request := &CreateTemplateRequest{
		Title:   title, //"linx-test",
		Link:    link,  //"www.baidu.com",
		Tenant:  tenant,
		Message: tempStr, //"我叫{name} ,今年{age}岁",
	}

	conn, err := global.GetAccountConn()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	template, err := NewMessageServiceClient(conn).CreateTemplate(context.Background(), request)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	if !template.Ok {
		log.Println(template.Message)
		return nil, fmt.Errorf(template.Message)
	}

	//log.Println(template)

	return template.Templage, nil
}

// GetTemplate 根绝tenant和模板id获取模板内容
// tenant 		  租户ID      必传		例如：13
// templateId     模板的id    必传		例如：10
func (MsgClient) GetTemplate(templateId, tenant int64) (*Template, error) {
	conn, err := global.GetAccountConn()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	template, err := NewMessageServiceClient(conn).GetTemplate(context.Background(), &GetTemplateRequest{Id: templateId, Tenant: tenant})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	if !template.Ok {
		log.Println(template.Message)
		return nil, fmt.Errorf(template.Message)
	}

	//log.Println(template)

	return template.Templage, nil
}

// GetTemplates 根绝tenant和模板id获取模板内容集
// tenant 		  租户ID      必传		例如：13
// templateId     模板的id    必传		例如：10
func (MsgClient) GetTemplates(tenant int64) ([]*Template, error) {
	conn, err := global.GetAccountConn()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	templates, err := NewMessageServiceClient(conn).GetTemplates(context.Background(), &GetTemplatesRequest{Tenant: tenant})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	if !templates.Ok {
		log.Println(templates.Message)
		return nil, fmt.Errorf(templates.Message)
	}

	//log.Println(template)

	return templates.Templages, nil
}

// CreateBinding 绑定模板消息可接受用户集合到消息模板，代表只有这个集合中的用户才可用查询到此模板生成的消息
// templateId 		消息模板
// tenant			租户id
func (MsgClient) CreateBinding(templateId, tenant int64, receivers []int64) (*Binding, error) {
	conn, err := global.GetAccountConn()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	bindRes, err := NewMessageServiceClient(conn).CreateBinding(context.Background(), &CreateBindingRequest{Template: templateId, Tenant: tenant, Receivers: receivers})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	if !bindRes.Ok {
		log.Println(bindRes.Message)
		return nil, fmt.Errorf(bindRes.Message)
	}

	//log.Println(bindRes)

	return bindRes.Binding, nil
}

// SetReceivers 重置绑定中的Receivers数组
// templateId 		消息模板
// tenant			租户id
func (MsgClient) SetReceivers(tenant, templateId int64, receivers []int64) error {
	conn, err := global.GetAccountConn()
	if err != nil {
		log.Println(err)
		return err
	}

	setRes, err := NewMessageServiceClient(conn).SetReceivers(context.Background(), &SetReceiversRequest{Template: templateId, Tenant: tenant, Receivers: receivers})
	if err != nil {
		log.Println(err)
		return err
	}
	if !setRes.Ok {
		log.Println(setRes.Message)
		return fmt.Errorf(setRes.Message)
	}

	//log.Println(bindRes)

	return nil
}

// SendMessage 发送消息
// templateId	模板ID
// tenant		租户ID
// msgArgs		消息模板中的参数map,key是模板中大括号中的字符串，value是对应需要填充的值
// linkArgs		链接模板中的参数map,key是模板中大括号中的字符串，value是对应需要填充的值
func (MsgClient) SendMessage(templateId, tenant int64, msgArgs, linkArgs map[string]string) (int64, error) {
	conn, err := global.GetAccountConn()
	if err != nil {
		log.Println(err)
		return 0, err
	}

	sendRes, err := NewMessageServiceClient(conn).SendMessage(context.Background(), &SendMessageRequest{
		Templage: templateId,
		Tenant:   tenant,
		MsgArgs:  msgArgs,
		LinkArgs: linkArgs,
	})
	if err != nil {
		log.Println(err)
		return 0, err
	}

	if !sendRes.Ok {
		log.Println(sendRes.Message)
		return 0, fmt.Errorf(sendRes.Message)
	}

	//log.Println(sendRes)

	return sendRes.Count, nil
}

// PageMessage 获取消息
// tenant	租户ID
// userId	用户ID
// page		页码
// size		每页的数量
func (MsgClient) PageMessage(userId, page, size, tenant int64) (*MessagePage, error) {
	conn, err := global.GetAccountConn()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	msgRes, err := NewMessageServiceClient(conn).PageMessage(context.Background(), &PageMessageRequest{
		User:   userId,
		Page:   page,
		Size:   size,
		Tenant: tenant,
	})
	if err != nil {
		log.Println(err)
		return nil, err
	}

	if !msgRes.Ok {
		log.Println(msgRes.Message)
		return nil, fmt.Errorf(msgRes.Message)
	}

	//log.Println(msgRes.Page)

	return msgRes.Page, nil
}

// MakeMessagesRead 设置为已读，可批量
// tenant		租户ID
// userId		用户ID
// messageIds	消息ID集
func (MsgClient) MakeMessagesRead(tenant, userId int64, messageIds []int64) error {
	conn, err := global.GetAccountConn()
	if err != nil {
		log.Println(err)
		return err
	}

	res, err := NewMessageServiceClient(conn).MakeMessagesRead(context.Background(), &MakeMessagesReadRequest{
		Messages: messageIds,
		Tenant:   tenant,
		User:     userId,
	})
	if err != nil {
		log.Println(err)
		return err
	}

	if !res.Ok {
		log.Println(res)
		return err
	}

	//log.Println(res)

	return nil
}

//// InitTemplates 初始化模板
//func InitTemplates(tenant int64) error {
//
//}
