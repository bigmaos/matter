package mattermanager

import (
	dc "daily_matter/util/diskcontrol"
	"encoding/json"
	"fmt"
)

type ManagerControler struct {
	MatterManager *MatterManager
}

const (
	matterManagerFile = "./MatterManager.json"
)

// 暂时使用json读写

func (c *ManagerControler) Load() error {
	jsonByte, err := dc.LoadJSON(matterManagerFile)
	if err != nil {
		return err
	}
	err = json.Unmarshal(jsonByte, &c.MatterManager)
	if err != nil {
		return err
	}
	fmt.Printf("load matter manager: %v\n", c.MatterManager)
	return nil
}

func (c *ManagerControler) Save() error {
	jsonBytes, err := json.Marshal(c.MatterManager)
	if err != nil {
		return err
	}
	if len(jsonBytes) == 0 {
		return fmt.Errorf("json bytes is empty")
	}
	err = dc.SaveJSON(matterManagerFile, jsonBytes)
	if err != nil {
		return err
	}
	fmt.Printf("save matter manager: %v\n", c.MatterManager)
	return nil
}
