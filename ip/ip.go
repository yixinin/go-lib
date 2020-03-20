package ip

import (
	"fmt"
	"go-lib/utils"
	"strings"
)

var LocalIP string

func init() {
	LocalIP = utils.IPAddress()
}

func GetAddr(port string) string {
	if strings.Contains(port, ":") {
		return port
	}
	return fmt.Sprintf("%s:%s", LocalIP, port)
}
