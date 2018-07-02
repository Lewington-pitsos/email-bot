package files

import (
	"encoding/json"
	"io/ioutil"
)

// +-------------------------------------------------------------------------------------+
// 									Vault STRUCT
// +-------------------------------------------------------------------------------------+

type Manager struct {
}

//
// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> HIDDEN METHODS <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
//

func (m *Manager) populate(filename string, emptyObject interface{}) {
	fileBytes, err := ioutil.ReadFile(filename)
	check(err)

	err2 := json.Unmarshal(fileBytes, emptyObject)
	check(err2)
}

//
// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> EXPOSED METHODS <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
//

func (m *Manager) BankeData(bankFile string) []string {
	bankData := make([]string, 0, 10000)
	m.populate(bankFile, &bankData)

	return bankData
}

// +-------------------------------------------------------------------------------------+
// 									EXPOSED FUNCTIONS
// +-------------------------------------------------------------------------------------+

func NewManager() *Manager {
	return &Manager{}
}
