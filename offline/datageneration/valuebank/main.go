package valuebank

func SetupBank() *Bank {
	bank := NewBank()
	bank.AddVault("username", "female-en")
	bank.AddVault("slang", "internet-slang")
	bank.AddSpecialVault("datevault")

	return bank
}
