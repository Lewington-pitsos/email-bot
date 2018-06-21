package files

import "go/build"

func check(err error) {
	if err != nil {
		panic(err)
	}
}

var profileBankPath = build.Default.GOPATH + "/src/email-bot/data/hotmail/"
var profileFileSuffix = ".json"

var bankFilePath = build.Default.GOPATH + "/src/email-bot/data/bankvalues/"
var bankFileSuffix = ".json"

var profileLedgerPath = "/src/email-bot/data/hotmail/ledger.json"
