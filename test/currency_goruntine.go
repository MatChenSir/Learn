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
	start := time.Now() // 获取当前时间
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
		{ID: 23, Name: "Env1", CloudEnv: 100},
		{ID: 24, Name: "Env2", CloudEnv: 200},
		{ID: 25, Name: "Env3", CloudEnv: 300},
		{ID: 26, Name: "Env4", CloudEnv: 0}, // CloudEnv 为 0 的环境跳过处理
		{ID: 27, Name: "Env1", CloudEnv: 100},
		{ID: 28, Name: "Env2", CloudEnv: 200},
		{ID: 29, Name: "Env3", CloudEnv: 300},
		{ID: 30, Name: "Env4", CloudEnv: 0}, // CloudEnv 为 0 的环境跳过处理
		{ID: 31, Name: "Env1", CloudEnv: 100},
		{ID: 32, Name: "Env2", CloudEnv: 200},
		{ID: 33, Name: "Env3", CloudEnv: 300},
		{ID: 34, Name: "Env4", CloudEnv: 0}, // CloudEnv 为 0 的环境跳过处理
		{ID: 35, Name: "Env1", CloudEnv: 100},
		{ID: 36, Name: "Env2", CloudEnv: 200},
		{ID: 37, Name: "Env3", CloudEnv: 300},
		{ID: 38, Name: "Env4", CloudEnv: 0}, // CloudEnv 为 0 的环境跳过处理		{ID: 1, Name: "Env1", CloudEnv: 100},
		{ID: 39, Name: "Env2", CloudEnv: 200},
		{ID: 40, Name: "Env3", CloudEnv: 300},
		{ID: 41, Name: "Env4", CloudEnv: 0}, // CloudEnv 为 0 的环境跳过处理		{ID: 1, Name: "Env1", CloudEnv: 100},
		{ID: 42, Name: "Env2", CloudEnv: 200},
		{ID: 43, Name: "Env3", CloudEnv: 300},
		{ID: 44, Name: "Env4", CloudEnv: 0}, // CloudEnv 为 0 的环境跳过处理
	}
	var wg sync.WaitGroup
	semaphore := make(chan struct{}, 10) //并发控制: 使用一个带有容量的通道 semaphore 来控制并发度，容量为 10
	//确保最多同时有 10 个 goroutine 在运行。每个 goroutine 进入时，会占用一个通道位置，执行完毕后释放，这样就限制了最大并发数量为 10。
	wg.Add(len(envList)) //等待总个数

	for _, env := range envList {
		semaphore <- struct{}{} // 占用一个并发信号量，控制并发度
		go func(env Environment) {
			defer func() {
				<-semaphore // 释放并发信号量
				if r := recover(); r != nil {
					fmt.Printf("Recovered from panic in goroutine: %v\n", r)
				}
				wg.Done()
			}()
			if env.CloudEnv == 0 {
				return
			}
			fmt.Printf("处理环境：%d，Cloud id：%v\n", env.ID, env.CloudEnv)
			singleStart := time.Now()
			time.Sleep(1 * time.Second)
			// 模拟某些操作
			// if env.ID == 3 {
			// 	panic("模拟发生 panic")
			// }
			singleElapsed := time.Since(singleStart)
			fmt.Printf("single check kess liveness cost %v ms, env id is %v, service KessName is %v\n",
				singleElapsed.Milliseconds(), env.ID, env.Name)
			//time.Sleep(1 * time.Second) // 模拟耗时操作
		}(env)
	}

	wg.Wait() //WaitGroup: 使用 sync.WaitGroup 等待所有 goroutine 完成任务。每个 goroutine 在启动时增加计数器，执行完毕后调用 wg.Done() 减少计数器。
	elapsed := time.Since(start)
	fmt.Printf("check kess liveness cost %v ms, envCount %v",
		elapsed.Milliseconds(), len(envList))
	fmt.Println("所有环境处理完成")
}
