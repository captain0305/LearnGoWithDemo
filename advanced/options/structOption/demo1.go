package main

import "fmt"

// ServerConfig 定义服务器配置
type ServerConfig struct {
	Port    int
	Timeout int
}

// ServerOption 定义结构体选项类型
type ServerOption struct {
	Port    int
	Timeout int
}

// NewServer 创建一个新的服务器实例
func NewServer(opts ...ServerOption) *ServerConfig {
	cfg := &ServerConfig{
		Port:    8080,
		Timeout: 30,
	}
	for _, opt := range opts {
		cfg.Port = opt.Port
		cfg.Timeout = opt.Timeout
	}
	return cfg
}

func main() {
	// 创建服务器实例并指定选项
	server := NewServer(
		ServerOption{Port: 9090, Timeout: 60},
	)

	fmt.Printf("Server Port: %d, Timeout: %d\n", server.Port, server.Timeout)
}
