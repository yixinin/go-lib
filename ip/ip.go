package ip

import (
	"go-lib/utils"
)

var LocalIP string

func init() {
	LocalIP = utils.IPAddress()
}
