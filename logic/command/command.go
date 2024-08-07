package command

import (
	"daily_matter/entity"
	"daily_matter/logic/dailymatter"
	"daily_matter/logic/state"
	"fmt"

	"github.com/spf13/cast"
)

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
			if len(args) != 2 {
				fmt.Printf("addUser err: %v\n", argsNotMatch)
				break
			}
			NewUserCommand(args[1])
			break
		}
	case "addmatter":
		{
			// 不含时间
			if len(args) == 3 {
				info := []string{args[1], args[2]}
				NewMatterCommand(nil, info)
			} else if len(args) == 6 {
				// 含时间
				time := []string{args[1], args[2], args[3]}
				info := []string{args[4], args[5]}
				NewMatterCommand(time, info)
			} else {
				fmt.Printf("addMatter err: %v\n", argsNotMatch)
			}
			break
		}
	case "getstate":
		{
			GetStateCommand()
			break
		}
	case "changestate":
		{
			if len(args) != 3 {
				fmt.Printf("changeState err: %v\n", argsNotMatch)
				break
			}
			ChangeStateCommand(args[1], cast.ToInt(args[2]))
			break
		}
	case "save":
		{
			Save()
			break
		}
	case "exit":
		{
			Exit()
			break
		}

	default:
		Help()
	}

}

func Help() {
	fmt.Printf("CommandManager help:\n")
	fmt.Printf("[display]:                                                     	display all matter\n")
	fmt.Printf("[adduser $usrname]:                                            	add user\n")
	fmt.Printf("[addmatter option{$gapUnit $startTime $endTime} $title $desc]: 	add matter\n")
	fmt.Printf("[getstate]:                                                    	get all state\n")
	fmt.Printf("[changestate $matterTitle $stateNumber]:                       	change matter state\n")
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

func NewMatterCommand(time []string, info []string) {
	if len(info) != 2 {
		fmt.Printf("NewMatterCommand err: %v", fmt.Errorf("format illegal"))
		return
	}
	var matterInfo = entity.InsertedMatterInfo{
		Title: info[0],
		Desc:  info[1],
	}
	if len(time) == 3 {
		matterInfo.TimeGap = time[0]
		matterInfo.StartTimeFromNow = cast.ToInt64(time[1])
		matterInfo.EndTimeFromNow = cast.ToInt64(time[2])
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
		fmt.Printf("%-*d %-*s \n", 15, i, 20, v)
	}
}

func ChangeStateCommand(matterTitle string, stateIdx int) {
	states := state.GetAllState()
	if stateIdx >= len(states) {
		fmt.Printf("ChangeStateCommand err: %v", fmt.Errorf("idx out of range"))
	}
	dailymatter.ChangeMatterState(matterTitle, states[stateIdx])
}

func Save() {
	err := dailymatter.Save()
	if err != nil {
		fmt.Printf("Save err: %v", err)
	}
}

func Exit() {
	Save()
	fmt.Printf("Exiting...")
}
