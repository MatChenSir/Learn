package test

// import (
// 	"encoding/json"
// 	"fmt"
// 	"log"

// 	v1 "k8s.io/api/core/v1"
// )

// // 模拟的 Env 结构体
// type Env struct {
// 	ID int `json:"id"`
// }

// // 模拟的 EnvVersion 结构体
// type EnvVersion struct {
// 	EnvVol string `json:"env_vol"`
// }

// // 模拟的 EnvOl 结构体
// type EnvOl struct {
// 	HostPaths []struct {
// 		Name     string `json:"name"`
// 		HostPath string `json:"host_path"`
// 		Type     string `json:"type,omitempty"`
// 	} `json:"host_paths"`

// 	EmptyDirs []struct {
// 		Name      string           `json:"name"`
// 		Medium    v1.StorageMedium `json:"medium"`
// 		SizeLimit *int64           `json:"size_limit"`
// 	} `json:"empty_dirs"`
// }

// // 模拟的 addVolumeWhetherExists 方法
// func addVolumeWhetherExists(v *[]v1.Volume, volume v1.Volume) {
// 	*v = append(*v, volume)
// }

// // addPaasCustomDefinitionVolume 函数实现
// func addPaasCustomDefinitionVolume(v *[]v1.Volume, env *Env, version *EnvVersion) error {
// 	var envVolsIn EnvOl
// 	if err := json.Unmarshal([]byte(version.EnvVol), &envVolsIn); err != nil {
// 		return fmt.Errorf("json unmarshal env vol failed. err: %v", err)
// 	}
// 	envVolsInBtye, _ := json.Marshal(envVolsIn)
// 	fmt.Printf("env %d, env vols %s\n", env.ID, string(envVolsInBtye))

// 	// 主机目录挂载
// 	for index, hostPath := range envVolsIn.HostPaths {
// 		volume := v1.Volume{
// 			Name: hostPath.Name,
// 			VolumeSource: v1.VolumeSource{
// 				HostPath: &v1.HostPathVolumeSource{
// 					Path: hostPath.HostPath,
// 				},
// 			},
// 		}
// 		// 添加判断，hostPath.Type json字段为空时，显示赋值会指向同一段内存地址
// 		if hostPath.Type != "" {
// 			volume.HostPath.Type = &hostPath.Type
// 			// 输出 volume.HostPath.Type 的值
// 			if volume.HostPath.Type != nil {
// 				fmt.Printf("env %d, env HostPath type is %v\n", env.ID, *volume.HostPath.Type)
// 			} else {
// 				fmt.Printf("env %d, env HostPath type is nil\n", env.ID)
// 			}
// 		}

// 		// 确保 volume.HostPath.Type 不是指向空字符串的指针
// 		if volume.HostPath.Type != nil && *volume.HostPath.Type == "" {
// 			volume.HostPath.Type = nil
// 			fmt.Printf("env %d, hostPath.Typ is null\n", env.ID)
// 		}
// 		volumeByte1, _ := json.Marshal(v)
// 		fmt.Printf("env %d, test1 the vol is %s\n", env.ID, string(volumeByte1))
// 		addVolumeWhetherExists(v, volume)
// 		vListByte, _ := json.Marshal(&v)
// 		volumeByte, _ := json.Marshal(volume)
// 		fmt.Printf("idx %d, env %d, after add host path, volume list is %s, volume is %s\n", index, env.ID, string(vListByte), string(volumeByte))
// 	}

// 	// 临时目录挂载
// 	for _, emptyDir := range envVolsIn.EmptyDirs {
// 		addVolumeWhetherExists(v, v1.Volume{
// 			Name: emptyDir.Name,
// 			VolumeSource: v1.VolumeSource{
// 				EmptyDir: &v1.EmptyDirVolumeSource{
// 					Medium:    emptyDir.Medium,
// 					SizeLimit: emptyDir.SizeLimit,
// 				},
// 			},
// 		})
// 	}

// 	return nil
// }

// func PersonalTestPointer() {
// 	// 创建模拟的数据
// 	v := make([]v1.Volume, 0)
// 	env := Env{ID: 1573621}
// 	version := EnvVersion{
// 		EnvVol: `[{"name":"api-log","host_path":"/data/logs/1573621-hb1-yz-prod-phhb1az4"},{"name":"kcs-log","host_path":"/data/logs/1573621-hb1-yz-prod-phhb1az4"},{"name":"model","host_path":"/media/disk1/cloud_local_storage","type":"DirectoryOrCreate"},{"name":"shm-pod","host_path":"/dev/shm","type":"DirectoryOrCreate"},{"name":"logs","host_path":"/home/web_server/kuaishou-worker/project/ad_twin_towers_retrieval.gpu/log"},{"name":"empty-dir","medium":"Memory","size_limit":1073741824}]`,
// 	}

// 	// 调用测试方法
// 	err := addPaasCustomDefinitionVolume(&v, &env, &version)
// 	if err != nil {
// 		log.Fatalf("Error occurred: %v", err)
// 	}

// 	// 输出最终的 Volumes
// 	vBytes, _ := json.MarshalIndent(v, "", "  ")
// 	fmt.Printf("Final Volumes: %s\n", vBytes)
// }
