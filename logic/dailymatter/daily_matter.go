package dailymatter

import (
	"daily_matter/constant"
	"daily_matter/entity"
	mm "daily_matter/logic/mattermanager"
	"daily_matter/logic/mattertime"
	"daily_matter/logic/state"
	"fmt"
	"sort"
	"time"
)

var currInfo = entity.NewCurrMattersInfo()

func Init() {
	manager := mm.Manager
	// 初始化时，只拿其中一个用户的matter
	for _, matters := range manager.GetUserMatters() {
		currInfo.User = matters.GetUser()
		for _, matter := range matters.GetMatters() {
			currInfo.Matters = append(currInfo.Matters, matter)
		}
		break
	}
	if currInfo == nil || currInfo.User == nil || len(currInfo.Matters) == 0 {
		return
	}
	sort.Slice(currInfo.Matters, func(i, j int) bool {
		return state.LessState(currInfo.Matters[i].GetState(), currInfo.Matters[j].GetState())
	})
}

func InsertNewUser(userid string) error {
	manager := mm.Manager
	err := manager.RegisterUser(userid)
	if err != nil {
		return err
	}
	return FreshCurrInfo()
}

// 初始化时设定，指定开始时间和结束时间距今时长（日历接入前方案），时间可以为空
func InsertNewMatter(info entity.InsertedMatterInfo) error {
	manager := mm.Manager
	var timeNow = time.Now()
	timeGap, ok := mattertime.GapUnitMap[info.TimeGap]
	matter := &entity.Matter{
		Title: info.Title,
		Desc:  info.Desc,
		State: constant.StateUnplanned,
	}
	if ok && info.StartTimeFromNow != 0 && info.EndTimeFromNow != 0 {
		timeStart := timeNow.Add(timeGap * time.Duration(info.StartTimeFromNow))
		timeEnd := timeNow.Add(timeGap * time.Duration(info.EndTimeFromNow))
		matter.SetStartTime(timeStart)
		matter.SetEndTime(timeEnd)
	}
	if currInfo.User == nil {
		return fmt.Errorf("currInfo.User is nil")
	}
	userid := currInfo.User.GetName()
	err := manager.GetUserMatters()[userid].RegisterMatter(matter)
	if err != nil {
		return err
	}
	return FreshCurrInfo()
}

func ChangeMatterState(matterid string, state string) error {
	manager := mm.Manager
	if currInfo.User == nil {
		return fmt.Errorf("currInfo.User is nil")
	}
	userid := currInfo.User.GetName()
	manager.GetUserMatters()[userid].GetMatters()[matterid].SetState(state)
	return FreshCurrInfo()
}

func Save() error {
	c := mm.Controler
	c.MatterManager.DeleteDoneMatter()
	err := c.Save()
	if err != nil {
		return err
	}
	return nil
}

func FreshCurrInfo() error {
	if currInfo.User == nil {
		return fmt.Errorf("currInfo.User is nil")
	}
	userid := currInfo.User.GetName()
	manager := mm.Manager
	currInfo.Matters = nil
	for _, matter := range manager.GetUserMatters()[userid].GetMatters() {
		currInfo.Matters = append(currInfo.Matters, matter)
	}
	if len(currInfo.Matters) == 0 {
		return fmt.Errorf("currInfo.Matters have no matter")
	}
	sort.Slice(currInfo.Matters, func(i, j int) bool {
		return state.LessState(currInfo.Matters[i].GetState(), currInfo.Matters[j].GetState())
	})
	return nil
}

func ChangeCurrUser(userid string) error {
	manager := mm.Manager
	user, ok := manager.GetUserMatters()[userid]
	if !ok {
		return fmt.Errorf("user %s not exist", userid)
	}
	currInfo.User = user.GetUser()
	return FreshCurrInfo()
}
