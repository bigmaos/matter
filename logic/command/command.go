package command

import (
	"daily_matter/entity"
	"daily_matter/logic/dailymatter"
	"daily_matter/logic/state"
	"fmt"

	"github.com/spf13/cast"
)

var argsTooShort = fmt.Errorf("args too short")
var argsNotMatch = fmt.Errorf("args not match")

func CommandManager(args ...string) {
	if len(args) == 0 {
		return
	}

	switch args[0] {
	case "display":
		DisplayCommand()
		break
	case "adduser":
		{
			if len(args) <= 2 {
				fmt.Printf("addUser err: %v", argsTooShort)
			}
			NewUserCommand(args[1])
			break
		}
	case "addmatter":
		{
			// 不含时间
			if len(args) == 5 {
				info := []string{args[2], args[3], args[4]}
				NewMatterCommand(args[1], nil, info)
			} else if len(args) == 7 {
				// 含时间
				time := []int64{cast.ToInt64(args[2]), cast.ToInt64(args[3])}
				info := []string{args[4], args[5], args[6]}
				NewMatterCommand(args[1], time, info)
			} else {
				fmt.Printf("addMatter err: %v", argsNotMatch)
			}

		}

	}

}

func DisplayCommand() {
	dailymatter.Display()
}

func NewUserCommand(userid string) {
	err := dailymatter.InsertNewUser(userid)
	if err != nil {
		fmt.Printf("NewUserCommand err: %v", err)
	}
}

/*
type InsertedMatterInfo struct {
	TimeGap          string
	StartTimeFromNow int64
	EndTimeFromNow   int64
	Title            string
	Desc             string
	State            string
}

*/

func NewMatterCommand(gapUnit string, timeFromNow []int64, info []string) {
	if len(timeFromNow) != 2 || len(info) != 3 {
		fmt.Printf("NewMatterCommand err: %v", fmt.Errorf("format illegal"))
	}
	var matterInfo = entity.InsertedMatterInfo{
		TimeGap:          gapUnit,
		StartTimeFromNow: timeFromNow[0],
		EndTimeFromNow:   timeFromNow[1],
		Title:            info[0],
		Desc:             info[1],
		State:            info[2],
	}
	err := dailymatter.InsertNewMatter(matterInfo)
	if err != nil {
		fmt.Printf("NewMatterCommand err: %v", err)
	}
}

func GetStateCommand() {
	states := state.GetAllState()
	fmt.Printf("All State hash: \nindex:\t\tstate\n")
	for i, v := range states {
		fmt.Printf("%-*d %-*s", 3, i, 20, v)
	}
}

func ChangeStateCommand(matterTitle string, stateIdx int) {
	states := state.GetAllState()
	if stateIdx >= len(states) {
		fmt.Printf("ChangeStateCommand err: %v", fmt.Errorf("idx out of range"))
	}
	dailymatter.ChangeMatterState(matterTitle, states[stateIdx])
}
