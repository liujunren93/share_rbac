module github.com/liujunren93/share_rbac

go 1.16

require (
	github.com/gin-gonic/gin v1.7.7
	github.com/liujunren93/share v0.0.0-20211026152618-a63a22d082cd
	github.com/liujunren93/share_utils v0.0.0-20220319164459-7b20e0b1a3f5
	github.com/sirupsen/logrus v1.8.1
	google.golang.org/grpc v1.41.0
	google.golang.org/protobuf v1.27.1
	gorm.io/gorm v1.23.1
)

replace github.com/liujunren93/share_utils => ../share_utils

//replace github.com/liujunren93/share => ../share
