package initialize

import (
	"fmt"
	"git.tuochexia.cn/tcxcn/common/accountcenter/rpcclient/messager"
	"gopkg.in/yaml.v2"
	"io"
	"log"
	"os"
)

type template struct {
	Title   string `json:"title"`
	Link    string `json:"link"`
	TempStr string `json:"tempStr"`
}

// 初始化一个tenant的模板
func initTemplates(tenant int64) error {
	templatesYAML, err := readTemplatesYAML()
	if err != nil {
		return err
	}
	client := messager.MsgClient{}
	templates, err := client.GetTemplates(tenant)
	if err != nil {
		log.Printf("向账户中心获取[tenant=%d]的所有模板失败\n", tenant)
		return fmt.Errorf("向账户中心获取[tenant=%d]的所有模板失败:err=%s\n", tenant, err.Error())
	}

	needInit := make(map[string]*template, 0)
	for key, template := range templatesYAML {
		te := template
		k := key
		needInit[k] = &te
	}

	for key := range templatesYAML {
		for _, te := range templates {
			if key == te.Title {
				//剩下的就是需要新加入的初始模板
				delete(needInit, key)
			}
		}
	}

	if len(needInit) > 0 {
		//开始初始化新加入的模板
		for _, te := range needInit {
			_, err = client.CreateTemplate(tenant, te.Title, te.Link, te.TempStr)
			if err != nil {
				log.Printf("远程调用账户中心的创建模板接口失败:err=%s\n", err.Error())
				log.Printf("参数为:[tenant=%d,title=%s,link=%s,tempStr=%s]\n", tenant, te.Title, te.Link, te.TempStr)
				return fmt.Errorf("[tenant=%d]远程调用账户中心的创建模板接口失败:err=%s\n", tenant, err.Error())
			}
		}
	}

	return nil

}

// 初始化所有tenant的模板
func initTemplates_(tenants []int64) error {
	//初始化所有的tenant的模板
	for _, t := range tenants {
		err := initTemplates(t)
		if err != nil {
			return err
		}
	}
	return nil
}

func readTemplatesYAML() (map[string]template, error) {
	file, err := os.Open("templates.yaml")
	if err != nil {
		log.Printf("初始化失败:读取templates.yaml失败:%s\n", err.Error())
		return nil, err
	}

	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		log.Printf("初始化失败:读取templates.yaml失败:%s\n", err.Error())
		return nil, err
	}

	var config map[string]template

	// 解析 YAML 数据到 map[string]Model 变量中
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Printf("初始化失败:解析templates.yaml失败:%s\n", err.Error())
		return nil, err
	}

	// 打印解析后的数据
	for key, value := range config {
		log.Println(key, value)
	}
	return config, nil
}
