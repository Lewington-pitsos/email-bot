package files

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Manager struct {
}

// Loads the database file matching that name into bytes.
// Marshals those bytes into Json data.
// Returns that JSON data.
func (m *Manager) ProfileData(profileName string) map[string]string {
	filename := m.filePath(profileName)

	readableProfile := make(map[string]string)
	m.populate(filename, &readableProfile)

	return readableProfile
}

func (m *Manager) BankeData(bankName string) []string {
	filename := m.bankPath(bankName)
	bankData := make([]string, 0, 10000)
	m.populate(filename, &bankData)

	return bankData
}

func (m *Manager) populate(filename string, emptyObject interface{}) {
	fileBytes, err := ioutil.ReadFile(filename)
	check(err)

	err2 := json.Unmarshal(fileBytes, emptyObject)
	check(err2)
}

func (m *Manager) saveFile(fileName string, data interface{}) {
	fmt.Println(data)
	dataBytes, err := json.Marshal(data)
	check(err)

	ioutil.WriteFile(fileName, dataBytes, 0777)
}

func (m *Manager) SaveProfile(profileName string, profileData map[string]string) {
	m.saveFile(m.filePath(profileName), profileData)
}

func (m *Manager) RecordProfile(profileName string) {
	profileLedger := make([]string, 0, 1000)
	m.populate(profileLedgerPath, profileLedger)
	profileLedger = append(profileLedger, profileName)
	m.saveFile(profileLedgerPath, profileLedger)

}

// Constructs an absolote path to the relevent profile file name.
func (m *Manager) filePath(fileName string) string {
	return profileBankPath + fileName + profileFileSuffix
}

func (m *Manager) bankPath(fileName string) string {
	return bankFilePath + fileName + bankFileSuffix
}

func NewManager() *Manager {
	return &Manager{}
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
