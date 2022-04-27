package initialize

import (
	"fmt"
	"gorm.io/gorm"
	"server/global"
)

const sys = "system"

func DBList() {
	dbMap := make(map[string]*gorm.DB)
	fmt.Println("打印当前DBlist的长度", len(global.GVA_CONFIG.DBList))
	for _, info := range global.GVA_CONFIG.DBList { //循环多sql配置.这里是配置多数据库用的
		if info.Disable {
			continue
		}
		switch info.Type {
		case "mysql":
			dbMap[info.Dbname] = GormMysqlByConfig(info)
			break
		case "pgsql":
			dbMap[info.Dbname] = GormPgSqlByConfig(info)
			break
		default:
			continue
		}
	}
	// 做特殊判断,是否有迁移
	// 适配低版本迁移多数据库版本
	if sysDB, ok := dbMap[sys]; ok {
		global.GVA_DB = sysDB
	}
	global.GVA_DBList = dbMap
}
