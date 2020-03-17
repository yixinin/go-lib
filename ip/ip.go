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
	return fmt.Sprintf(LocalIP, port)
}
