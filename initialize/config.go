package initialize

type Config struct {
	GrpcConfig GrpcConfig `json:"grpcConfig" yaml:"grpcConfig"`
}

type GrpcConfig struct {
	Addr   string `yaml:"address"`
	Port   uint16 `yaml:"port"`
	AppId  string `yaml:"appId"`
	Secret string `yaml:"secret"`
}
