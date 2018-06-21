package files

import 	"build"

func Check(err error) {
	if	err != nil {
		panic(err)
	}
}

var profileBankPath = build.Default.GOPATH + "src/email-bot/data/"
var profileFileSuffix = ".json"