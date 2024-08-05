package mattermanager

import (
	"daily_matter/entity"
	"daily_matter/util/diskcontrol"
	"fmt"
)

var Manager *MatterManager

type MatterManager struct {
	UserMatters map[string]MatterSingleUser `json:"user_matters"`
}

type MatterSingleUser struct {
	User    *entity.User              `json:"user"`
	Matters map[string]*entity.Matter `json:"matters"`
}

func NewMatterManager() *MatterManager {
	return &MatterManager{
		UserMatters: make(map[string]MatterSingleUser),
	}
}

func Init() {
	c := diskcontrol.ManagerControler{}
	err := c.Load()
	if err != nil {
		fmt.Printf("load manager error:%v", err)
		Manager = NewMatterManager()
		return
	}
	Manager = c.MatterManager
}

func (m *MatterManager) GetMatters() map[string]MatterSingleUser {
	return m.UserMatters
}
