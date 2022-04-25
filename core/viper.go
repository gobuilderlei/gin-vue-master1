package core

import (
	"flag"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/songzhibin97/gkit/cache/local_cache"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
	"server/global"
	"server/utils"
	"time"
)

//可变参数    ...string
//viper go语言配置管理神器,主要处理 yaml配置的的
func Viper(path ...string) *viper.Viper {
	var config string //主要申明为了存储config配置文件的路劲
	if len(path) == 0 {
		flag.StringVar(&config, "c", "", "choose config file.[选择配置文件]")
		flag.Parse()
		if config == "" {
			configEnv := os.Getenv(utils.ConfigEnv)
			if configEnv == "" {
				config = utils.ConfigFile
				fmt.Printf("您正在使用config默认值,config的路劲为%v\n", utils.ConfigFile)
			} else {
				config = configEnv
				fmt.Printf("您正在使用GVA_CONFIG环境变量,config的路径为%v\n", config)
			}
		} else {
			fmt.Printf("您正在使用命令行的-c参数传递的值,config的路径为%v\n", config)
		}
	} else {
		config = path[0]
		fmt.Printf("您正在使用func Viper()传递的值,config的路径为%v\n", config)
	}
	v := viper.New()
	v.SetConfigFile(config)
	v.SetConfigType("yaml") //设置配置文件后缀类型
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) { //判断配置文件是否变更了
		fmt.Println("config file changed:", e.Name)
		if erro := v.Unmarshal(&global.GVA_CONFIG); erro != nil {
			fmt.Println("配置文件config.yaml出现错误,错误代码为:", erro)
		}
	})
	if err := v.Unmarshal(&global.GVA_CONFIG); err != nil {
		fmt.Println(err)
	}
	global.GVA_CONFIG.AutoCode.Root, _ = filepath.Abs("..")
	global.BlackCache = local_cache.NewCache(local_cache.SetDefaultExpire(time.Second * time.Duration(global.GVA_CONFIG.JWT.ExpiresTime)))
	return v
}
