package bootstrap

import (
	"Goplus/pkg/config"
	"Goplus/pkg/logger"
	"Goplus/pkg/oss"
)

func SetupOss() {
	var oss oss.OssStruct
	oss.Endpoint = "oss-cn-beijing.aliyuncs.com"
	oss.AccessKeyId = config.Get("oss.accessKeyId")
	oss.AccessKeySecret = config.Get("oss.accessKeySecret")
	oss.Bucket = "go-land"
	oss.Region = "华北2(北京)"

	if err := oss.OssCennect(&oss); err != nil {
		logger.LogIf(err)
	}
}
