package main

import (
	"github.com/btcziggurat/btcsuite_to_btcziggurat/wallet_spawn"
	"github.com/jfixby/pin"
)

func main() {

	pin.D("Hello!")
	pfcwallet_spawn.Build()
	pin.D("PFCD build done!")

}
