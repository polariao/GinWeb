package helper

import (
	"github.com/polariao/polarLog/polarLog"
)

var PolarLog *polarLog.PolarLog

func InitPolarLog(path string)  {
	//初始化日志
	log := polarLog.PolarLog{Path:path}
	PolarLog = log.Default()
}
