package proc

import "github.com/picfight/coin_knife/fileproc"

func PicfightCoinFileNameGenerator(data string) string {
	data = fileproc.Replace(data, "btcsuite", "picfight")
	data = fileproc.Replace(data, "bitcoind", "picfightcoind")
	data = fileproc.Replace(data, "btc", "pfc")
	return data
}

func PicfightCoinFileGenerator(data string) string {
	data = fileproc.Replace(data, "AmountBTC", "AmountPFC")
	data = fileproc.Replace(data, "MicroBTC", "MicroPFC")
	data = fileproc.Replace(data, "MilliBTC", "MilliPFC")
	data = fileproc.Replace(data, "MegaBTC", "MegaPFC")
	data = fileproc.Replace(data, "KiloBTC", "KiloPFC")
	data = fileproc.Replace(data, "BTC(", "PFC(")
	data = fileproc.Replace(data, "BTC ", "PFC ")
	data = fileproc.Replace(data, " BTC", " PFC")
	data = fileproc.Replace(data, `"BTC"`, `"PFC"`)
	data = fileproc.Replace(data, `kBTC`, `kPFC`)
	data = fileproc.Replace(data, `mBTC`, `mPFC`)
	data = fileproc.Replace(data, `MBTC`, `MPFC`)

	data = fileproc.Replace(data, `BtcAddr`, `PfcAddr`)
	data = fileproc.Replace(data, `btcAddr`, `pfcAddr`)
	data = fileproc.Replace(data, `btcaddr`, `pfcaddr`)

	data = fileproc.Replace(data, `btcsuite\btcwallet`, `picfight\pfcwallet`)
	data = fileproc.Replace(data, `btcsuite/btcwallet`, `picfight/pfcwallet`)

	data = fileproc.Replace(data, `GOPATH\src\github.com\btcsuite\btcd`, `GOPATH\src\github.com\picfight\pfcd`)
	data = fileproc.Replace(data, `GOPATH/src/github.com/btcsuite`, `GOPATH/src/github.com/picfight`)
	data = fileproc.Replace(data, `github.com/lightninglabs/neutrino`, `github.com/picfight/pfcneutrino`)
	data = fileproc.Replace(data, `lightningnetwork/lnd`, `picfight/pfclnd`)

	data = fileproc.Replace(data, "btcsuite/btcutil", "picfight/pfcutil")
	data = fileproc.Replace(data, "btcutil", "pfcutil")

	data = fileproc.Replace(data, "btcsuite/btcwallet", "picfight/pfcwallet")
	data = fileproc.Replace(data, "btcd/btcwallet", "pfcd/pfcwallet")
	data = fileproc.Replace(data, `btcwallet`, `pfcwallet`)

	data = fileproc.Replace(data, "btcsuite/btcd/btcec", "picfight/pfcd/pfcec")
	data = fileproc.Replace(data, "btcec", "pfcec")

	data = fileproc.Replace(data, `btcsuite/btcd/btcjson`, `picfight/pfcd/pfcjson`)
	data = fileproc.Replace(data, `btcjson`, `pfcjson`)

	data = fileproc.Replace(data, `"github.com/lightninglabs/neutrino"`, `"github.com/picfight/pfcneutrino"`)

	data = fileproc.Replace(data, `btcsuite/btcd/btcjson`, `picfight/pfcd/pfcjson`)
	data = fileproc.Replace(data, `btcjson`, `pfcjson`)

	data = fileproc.Replace(data, `jfixby/btcharness`, `picfight/pfcharness`)
	data = fileproc.Replace(data, `btcharness`, `pfcharness`)
	data = fileproc.Replace(data, `jfixby/btcregtest`, `picfight/pfcregtest`)
	data = fileproc.Replace(data, `btcregtest`, `pfcregtest`)

	data = fileproc.Replace(data, `btcsuite/btcd`, `picfight/pfcd`)

	data = fileproc.Replace(data, `btcd,`, `pfcd,`)
	data = fileproc.Replace(data, `btcd `, `pfcd `)
	data = fileproc.Replace(data, `btcd.`, `pfcd.`)

	data = fileproc.Replace(data, `./btcd`, `./pfcd`)
	data = fileproc.Replace(data, `"BTCD"`, `"PFCD"`)
	data = fileproc.Replace(data, "`BTCD`", "`PFCD`")

	data = fileproc.Replace(data, `btcd.conf`, `pfcd.conf`)

	data = fileproc.Replace(data, "'btcd'", "'pfcd'")
	data = fileproc.Replace(data, "`btcd`", "`pfcd`")
	data = fileproc.Replace(data, `"btcd"`, `"pfcd"`)
	data = fileproc.Replace(data, ` btcd `, ` pfcd `)
	data = fileproc.Replace(data, ` Btcd `, ` Pfcd `)
	data = fileproc.Replace(data, `">Btcd `, `">Pfcd `)

	data = fileproc.Replace(data, `BtcDecode`, `PfcDecode`)
	data = fileproc.Replace(data, `BtcEncode`, `PfcEncode`)
	data = fileproc.Replace(data, `BtcPerK`, `PfcPerK`)
	data = fileproc.Replace(data, `BtcdConnected`, `PfcdConnected`)
	data = fileproc.Replace(data, `BtcctlConfig`, `PfcctlConfig`)
	data = fileproc.Replace(data, `TestBtc`, `TestPfc`)
	data = fileproc.Replace(data, `findBTCDFolder`, `findPFCDFolder`)

	data = fileproc.Replace(data, `oldBtcdHomeDir`, `oldPfcdHomeDir`)
	data = fileproc.Replace(data, `Btcd Service`, `Pfcd Service`)
	data = fileproc.Replace(data, `Btcd Suite`, `Pfcd Suite`)
	data = fileproc.Replace(data, `Btcsuite`, `Picfight`)
	data = fileproc.Replace(data, `Support/Btcd/data`, `Support/Pfcd/data`)
	data = fileproc.Replace(data, `$LOCALAPPDATA/Btcd`, `$LOCALAPPDATA/Pfcd`)
	data = fileproc.Replace(data, `Btcd\addblock`, `Pfcd\addblock`)
	data = fileproc.Replace(data, `BTCKey%2BTests.m#L23-L49`, `PFCKey%2BTests.m#L23-L49`)

	//data = fileproc.Replace(data, `Btc`, `Pfc`)

	data = fileproc.Replace(data, `btcnet`, `pfcnet`)
	data = fileproc.Replace(data, `BitcoinNet`, `PicfightcoinNet`)
	data = fileproc.Replace(data, ` Bitcoin`, ` Picfightcoin`)
	data = fileproc.Replace(data, `bitcoinTip`, `picfightcoinTip`)
	data = fileproc.Replace(data, `bitcoinHeader`, `picfightcoinHeader`)

	data = fileproc.Replace(data, `(AmountBitcoin)`, `(AmountPicfightcoin)`)
	data = fileproc.Replace(data, `SatoshiPerBitcoin`, `SatoshiPerPicfightcoin`)
	data = fileproc.Replace(data, `Bitcoin`, `Picfightcoin`)
	data = fileproc.Replace(data, ` bitcoin`, ` picfightcoin`)

	data = fileproc.Replace(data, `btcharness`, `pfcharness`)
	data = fileproc.Replace(data, `btcregtest`, `pfcregtest`)

	data = fileproc.Replace(data, `btcAddress`, `pfcAddress`)

	data = fileproc.Replace(data, `Btcwallet`, `Pfcwallet`)

	data = fileproc.Replace(data, `newBtcAddress`, `newPfcAddress`)

	//---------------------------------------------------------
	data = fileproc.Replace(data, `Btcd`, `Pfcd`)
	data = fileproc.Replace(data, `btcd`, `pfcd`)

	return data
}
