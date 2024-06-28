package test

import (
	"fmt"
	"sync"
	"time"
)

// Environment 模拟环境对象
type Environment struct {
	ID       int
	Name     string
	CloudEnv int
}

func TestGrountine() {
	envList := []Environment{
		{ID: 1, Name: "Env1", CloudEnv: 100},
		{ID: 2, Name: "Env2", CloudEnv: 200},
		{ID: 3, Name: "Env3", CloudEnv: 300},
		{ID: 4, Name: "Env4", CloudEnv: 0}, // CloudEnv 为 0 的环境跳过处理
		{ID: 5, Name: "Env1", CloudEnv: 100},
		{ID: 6, Name: "Env2", CloudEnv: 200},
		{ID: 7, Name: "Env3", CloudEnv: 300},
		{ID: 8, Name: "Env4", CloudEnv: 0}, // CloudEnv 为 0 的环境跳过处理
		{ID: 9, Name: "Env1", CloudEnv: 100},
		{ID: 10, Name: "Env2", CloudEnv: 200},
		{ID: 11, Name: "Env3", CloudEnv: 300},
		{ID: 12, Name: "Env4", CloudEnv: 0}, // CloudEnv 为 0 的环境跳过处理
		{ID: 13, Name: "Env1", CloudEnv: 100},
		{ID: 14, Name: "Env2", CloudEnv: 200},
		{ID: 15, Name: "Env3", CloudEnv: 300},
		{ID: 16, Name: "Env4", CloudEnv: 0}, // CloudEnv 为 0 的环境跳过处理		{ID: 1, Name: "Env1", CloudEnv: 100},
		{ID: 17, Name: "Env2", CloudEnv: 200},
		{ID: 18, Name: "Env3", CloudEnv: 300},
		{ID: 19, Name: "Env4", CloudEnv: 0}, // CloudEnv 为 0 的环境跳过处理		{ID: 1, Name: "Env1", CloudEnv: 100},
		{ID: 20, Name: "Env2", CloudEnv: 200},
		{ID: 21, Name: "Env3", CloudEnv: 300},
		{ID: 22, Name: "Env4", CloudEnv: 0}, // CloudEnv 为 0 的环境跳过处理
		{ID: 1, Name: "Env1", CloudEnv: 100},
		{ID: 2, Name: "Env2", CloudEnv: 200},
		{ID: 3, Name: "Env3", CloudEnv: 300},
		{ID: 4, Name: "Env4", CloudEnv: 0}, // CloudEnv 为 0 的环境跳过处理
		{ID: 5, Name: "Env1", CloudEnv: 100},
		{ID: 6, Name: "Env2", CloudEnv: 200},
		{ID: 7, Name: "Env3", CloudEnv: 300},
		{ID: 8, Name: "Env4", CloudEnv: 0}, // CloudEnv 为 0 的环境跳过处理
		{ID: 9, Name: "Env1", CloudEnv: 100},
		{ID: 10, Name: "Env2", CloudEnv: 200},
		{ID: 11, Name: "Env3", CloudEnv: 300},
		{ID: 12, Name: "Env4", CloudEnv: 0}, // CloudEnv 为 0 的环境跳过处理
		{ID: 13, Name: "Env1", CloudEnv: 100},
		{ID: 14, Name: "Env2", CloudEnv: 200},
		{ID: 15, Name: "Env3", CloudEnv: 300},
		{ID: 16, Name: "Env4", CloudEnv: 0}, // CloudEnv 为 0 的环境跳过处理		{ID: 1, Name: "Env1", CloudEnv: 100},
		{ID: 17, Name: "Env2", CloudEnv: 200},
		{ID: 18, Name: "Env3", CloudEnv: 300},
		{ID: 19, Name: "Env4", CloudEnv: 0}, // CloudEnv 为 0 的环境跳过处理		{ID: 1, Name: "Env1", CloudEnv: 100},
		{ID: 20, Name: "Env2", CloudEnv: 200},
		{ID: 21, Name: "Env3", CloudEnv: 300},
		{ID: 22, Name: "Env4", CloudEnv: 0}, // CloudEnv 为 0 的环境跳过处理
	}

	var wg sync.WaitGroup
	concurrency := 10
	semaphore := make(chan struct{}, concurrency)

	for _, env := range envList {
		if env.CloudEnv == 0 {
			continue
		}

		semaphore <- struct{}{} // 占用一个并发信号量

		wg.Add(1)
		go func(env *Environment) {
			defer func() {
				<-semaphore // 释放并发信号量
				wg.Done()
			}()

			fmt.Printf("处理环境：%s，CloudEnv：%d, id is:%d\n", env.Name, env.CloudEnv, env.ID)
			time.Sleep(1 * time.Second) // 模拟耗时操作
		}(&env) // 传递环境的指针
	}

	wg.Wait()
}
