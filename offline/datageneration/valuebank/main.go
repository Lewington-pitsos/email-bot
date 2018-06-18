package valuebank

func SetupBank() *Bank {
	bank := NewBank()
	bank.AddVault("username", "female-en")
	bank.AddVault("slang", "internet-slang")
	bank.AddSpecialVault("yearvault")
	bank.AddSpecialVault("monthvault")
	bank.AddSpecialVault("dayvault")
	bank.AddSpecialVault("passvault")

	return bank
}
