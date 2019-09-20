package main

import (
	"github.com/btcziggurat/btcsuite_to_btcziggurat/util_spawn"
	"github.com/jfixby/pin"
)

func main() {

	pin.D("Hello!")
	util_spawn.Build()
	pin.D("PFCD build done!")

}
