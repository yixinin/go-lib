package main

import (
	"fmt"

	"go-lib/registry"

	"go-lib/registry/etcd"
)

func main() {
	var reg registry.Registry
	reg = etcd.NewRegistry()
	fmt.Println("", reg.String())
}
