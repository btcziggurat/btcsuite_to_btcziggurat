package main

import (
	"github.com/btcziggurat/btcsuite_to_btcziggurat/regtest_spawn"
	"github.com/jfixby/pin"
)

func main() {

	pin.D("Hello!")
	pfcregtest_spawn.Build()
	pin.D("PFCD build done!")

}
