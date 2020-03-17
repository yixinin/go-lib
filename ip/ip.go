package ip

import (
	"fmt"
	"go-lib/utils"
)

var LocalIP string

func init() {
	LocalIP = utils.IPAddress()
}

func GrpcAddr(port string) string {
	return fmt.Sprintf("%s:%s", LocalIP, port)
}

func GetAddr(port string) string {
	return fmt.Sprintf("%s:%s", LocalIP, port)
}
