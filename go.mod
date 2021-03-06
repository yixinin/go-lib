module go-lib

go 1.13

require (
	github.com/coreos/etcd v3.3.18+incompatible
	github.com/docker/docker v1.13.1 // indirect
	github.com/go-redis/redis/v7 v7.2.0
	github.com/go-sql-driver/mysql v1.5.0
	github.com/go-xorm/xorm v0.7.9
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
	github.com/golang/protobuf v1.3.2
	github.com/konsorten/go-windows-terminal-sequences v1.0.2
	github.com/micro/go-micro v1.18.0
	github.com/micro/go-micro/v2 v2.2.0
	github.com/mitchellh/hashstructure v1.0.0
	github.com/satori/go.uuid v1.2.0
	github.com/sirupsen/logrus v1.4.2
	github.com/stretchr/testify v1.4.0
	go.mongodb.org/mongo-driver v1.3.1
	go.uber.org/zap v1.14.1
	google.golang.org/grpc v1.26.0
	gopkg.in/yaml.v2 v2.2.8
	xorm.io/core v0.7.2-0.20190928055935-90aeac8d08eb
)

replace github.com/coreos/go-systemd => github.com/coreos/go-systemd/v22 v22.0.0

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
