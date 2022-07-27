package log

import (
	"github.com/sirupsen/logrus"
)

/**
* @Author: liujunren
* @Date: 2022/3/7 15:40
 */

var Logger *logrus.Logger

func InitLogger(log *logrus.Logger) {
	Logger = log
}
