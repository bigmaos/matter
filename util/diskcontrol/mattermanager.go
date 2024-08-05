package diskcontrol

import (
	mm "daily_matter/logic/mattermanager"
	"encoding/json"
)

type ManagerControler struct {
	MatterManager *mm.MatterManager
}

const (
	matterManagerFile = "./MatterManager.json"
)

// 暂时使用json读写

func (c *ManagerControler) Load() error {
	jsonByte, err := LoadJSON(matterManagerFile)
	if err != nil {
		return err
	}
	err = json.Unmarshal(jsonByte, c.MatterManager)
	if err != nil {
		return err
	}
	return nil
}

func (c *ManagerControler) Save() error {
	jsonBytes, err := json.Marshal(c.MatterManager)
	if err != nil {
		return err
	}
	err = SaveJSON(matterManagerFile, jsonBytes)
	if err != nil {
		return err
	}
	return nil
}
