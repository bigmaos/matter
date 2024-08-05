package mattermanager

import (
	"daily_matter/constant"
	"daily_matter/entity"
	"daily_matter/util/diskcontrol"
	"fmt"
)

var Manager *MatterManager

type MatterManager struct {
	UserMatters map[string]*MatterSingleUser `json:"user_matters"`
}

type MatterSingleUser struct {
	User    *entity.User              `json:"user"`
	Matters map[string]*entity.Matter `json:"matters"`
}

func NewMatterManager() *MatterManager {
	return &MatterManager{
		UserMatters: make(map[string]*MatterSingleUser),
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
	defer func() {
		err := c.Save()
		if err != nil {
			fmt.Printf("save manager error:%v", err)
		}
	}()
}

func (m *MatterManager) GetUserMatters() map[string]*MatterSingleUser {
	return m.UserMatters
}

func (m *MatterManager) RegisterUser(userid string) error {
	if _, ok := m.GetUserMatters()[userid]; ok {
		return fmt.Errorf("user %s already exists", userid)
	}
	m.GetUserMatters()[userid] = &MatterSingleUser{
		User:    &entity.User{Name: userid},
		Matters: make(map[string]*entity.Matter),
	}
	return nil
}

func (m *MatterSingleUser) RegisterMatter(matter *entity.Matter) error {
	if _, ok := m.GetMatters()[matter.GetTitle()]; ok {
		return fmt.Errorf("matter %s already exists", matter.GetTitle())
	}
	m.GetMatters()[matter.GetTitle()] = matter
	return nil
}

func (m *MatterSingleUser) GetMatter(title string) *entity.Matter {
	return m.GetMatters()[title]
}

func (m *MatterSingleUser) GetUser() *entity.User {
	return m.User
}

func (m *MatterSingleUser) GetMatters() map[string]*entity.Matter {
	return m.Matters
}

// DeleteDoneMatter 统一在落盘时删除已完成事项
func (m *MatterManager) DeleteDoneMatter() {
	for _, single := range m.GetUserMatters() {
		for _, matter := range single.GetMatters() {
			if matter.State == constant.StateDone {
				delete(single.GetMatters(), matter.GetTitle())
			}
		}
	}
}
