package main

import (
	"github.com/semiozbas/docker-machine-opennebula"
	"github.com/docker/machine/libmachine/drivers/plugin"
)

func main() {
	plugin.RegisterDriver(opennebula.NewDriver("", ""))
}
