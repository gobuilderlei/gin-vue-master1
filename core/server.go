package core

import (
	"server/global"
	"server/initialize"
	"server/service/system"
)

type server interface {
	ListenAndServe() error
}

func RunWindowsServer() {
	if global.GVA_CONFIG.System.UseMultipoint || global.GVA_CONFIG.System.UseRedis {
		initialize.Redis()
	}
	// 从db加载jwt数据
	if global.GVA_DB != nil {
		system.LoadAll()
	}
}
