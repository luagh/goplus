package bootstrap

import (
	"Goplus/pkg/config"
	"Goplus/pkg/logger"
)

// 存放初始化logger的逻辑
func SetupLogger() {
	logger.InitLogger(
		config.GetString("log.filename"),
		config.GetInt("log.max_size"),
		config.GetInt("log.max_backup"),
		config.GetInt("log.max_age"),
		config.GetBool("log.compress"),
		config.GetString("log.type"),
		config.GetString("log.level"),
	)
}
