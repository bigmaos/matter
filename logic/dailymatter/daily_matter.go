package dailymatter

import (
	"daily_matter/constant"
	"daily_matter/entity"
	mm "daily_matter/logic/mattermanager"
	"daily_matter/logic/mattertime"
	"daily_matter/logic/state"
	"sort"
	"time"
)

var currInfo *entity.CurrMattersInfo

func init() {
	manager := mm.Manager
	// 初始化时，只拿其中一个用户的matter
	for _, matters := range manager.GetUserMatters() {
		currInfo.User = matters.GetUser()
		for _, matter := range matters.GetMatters() {
			currInfo.Matters = append(currInfo.Matters, matter)
		}
		break
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
	return nil
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
	userid := currInfo.User.GetName()
	return manager.GetUserMatters()[userid].RegisterMatter(matter)
}

func ChangeMatterState(matterid string, state string) {
	manager := mm.Manager
	userid := currInfo.User.GetName()
	manager.GetUserMatters()[userid].GetMatters()[matterid].SetState(state)
}
