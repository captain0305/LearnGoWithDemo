package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

type AuthConfigCache struct {
	authConfigsPointer *atomic.Pointer[AuthConfigMap]
}

type AuthConfigMap map[string][]AuthConfig

type AuthConfigs []AuthConfig

type AuthConfig struct {
	APIType int32 `json:"api_type"` // api 类型

	TargetFieldName string `json:"target_field_name"` // 目标字段名
	AuthPath        string `json:"auth_path"`         // 请求路径

}

// NewFunctionCache 创建一个新的 FunctionCache
func NewAuthConfigCache() *AuthConfigCache {
	return &AuthConfigCache{
		authConfigsPointer: &atomic.Pointer[AuthConfigMap]{},
	}
}

func main() {
	cache := NewAuthConfigCache()

	// 初始化 authConfigMap
	authConfigMap := AuthConfigMap{
		"key1": []AuthConfig{{APIType: 1, TargetFieldName: "field1", AuthPath: "/path1"}},
		"key2": []AuthConfig{{APIType: 2, TargetFieldName: "field2", AuthPath: "/path2"}},
	}

	// 设置 authConfigMap 的值
	cache.authConfigsPointer.Store(&authConfigMap)
	authConfigs := *cache.authConfigsPointer.Load()
	fmt.Println("Read from main:", authConfigs)
	// 并发读取 authConfigMap
	go func() {
		for {
			authConfigs := *cache.authConfigsPointer.Load()
			fmt.Println("Read from Goroutine 1:", authConfigs)
			time.Sleep(1 * time.Second)
		}
	}()

	// 并发读取 authConfigMap
	go func() {
		for {
			authConfigs := *cache.authConfigsPointer.Load()
			fmt.Println("Read from Goroutine 2:", authConfigs)
			time.Sleep(1 * time.Second)
		}
	}()

	// 并发更新 authConfigMap
	go func() {
		i := 0
		for {
			newAuthConfigMap := AuthConfigMap{
				"key1": []AuthConfig{{APIType: int32(i), TargetFieldName: "field1_updated", AuthPath: "/path1_updated"}},
				"key2": []AuthConfig{{APIType: int32(i + 1), TargetFieldName: "field2_updated", AuthPath: "/path2_updated"}},
			}
			cache.authConfigsPointer.Store(&newAuthConfigMap)
			fmt.Println("Updated authConfigMap")
			time.Sleep(5 * time.Second)
			i++
		}
	}()

	// 防止主 Goroutine 退出
	select {}
}
