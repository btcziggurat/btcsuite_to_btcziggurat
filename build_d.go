package main

import (
	"github.com/btcziggurat/btcsuite_to_btcziggurat/d_spawn"
	"github.com/jfixby/pin"
)

func main() {

	pin.D("Hello!")
	d_spawn.Build()
	pin.D("D build done!")

}
