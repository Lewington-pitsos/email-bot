package files

import (
	"io/ioutil"
	"encoding/json"
)

type Manager struct {
}

// Loads the database file matching that name into bytes.
// Marshals those bytes into Json data.
// Returns that JSON data.
func(m *Manager) LoadProfile(profileName string) map[string]string {
	profileBytes, err := ioutil.ReadFile(m.filePath(profileName))
	check(err)

	readableProfile := make(map[string]string)
	err2 := json.Unmarshal(profileBytes, readableProfile)
	check(err2)

	return readableProfile
}

// Constructs an absolote path to the relevent profile file name. 
func(m *Manager) filePath(fileName string) string {
	return profileBankPath + fileName + profileFilSuffix
}