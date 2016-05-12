package main

import (
	"github.com/docker/machine/libmachine/drivers/plugin"
	"github.com/naymspace/docker-machine-driver-profitbricks"
)

func main() {
	plugin.RegisterDriver(profitbricks.NewDriver("", ""))
}
