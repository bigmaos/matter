package command

import (
	"daily_matter/entity"
	"daily_matter/logic/dailymatter"
	"daily_matter/logic/state"
	"fmt"

	"github.com/spf13/cast"
)

var argsNotMatch = fmt.Errorf("args not match")

type CommandManager struct {
	dp dailymatter.DisplayPacker
}

func NewCommandManager(packer dailymatter.DisplayPacker) *CommandManager {
	// 无传入：默认控制台Display
	if packer == nil {
		return &CommandManager{
			dp: dailymatter.DisplayConsolePacker{},
		}
	}
	return &CommandManager{
		dp: packer,
	}
}

func (c *CommandManager) Manager(args ...string) {
	if len(args) == 0 {
		return
	}

	switch args[0] {
	case "display":
		c.DisplayCommand()
		break
	case "adduser":
		{
			if len(args) != 2 {
				fmt.Printf("addUser err: %v\n", argsNotMatch)
				break
			}
			c.NewUserCommand(args[1])
			break
		}
	case "addmatter":
		{
			// 不含时间
			if len(args) == 3 {
				info := []string{args[1], args[2]}
				c.NewMatterCommand(nil, info)
			} else if len(args) == 6 {
				// 含时间
				time := []string{args[1], args[2], args[3]}
				info := []string{args[4], args[5]}
				c.NewMatterCommand(time, info)
			} else {
				fmt.Printf("addMatter err: %v\n", argsNotMatch)
			}
			break
		}
	case "getstate":
		{
			c.GetStateCommand()
			break
		}
	case "changestate":
		{
			if len(args) != 3 {
				fmt.Printf("changeState err: %v\n", argsNotMatch)
				break
			}
			c.ChangeStateCommand(args[1], cast.ToInt(args[2]))
			break
		}
	case "changeuser":
		{
			if len(args) != 2 {
				fmt.Printf("changeUser err: %v\n", argsNotMatch)
				break
			}
			c.ChangeUserCommand(args[1])
			break
		}

	case "save":
		{
			c.Save()
			break
		}
	case "fresh":
		{
			c.Fresh()
			break
		}

	case "exit":
		{
			c.Exit()
			break
		}

	default:
		Help()
	}

}

func Help() {
	fmt.Printf("Manager help:\n")
	fmt.Printf("[display]:                                                     	display curr user info\n")
	fmt.Printf("[adduser $usrname]:                                            	add user\n")
	fmt.Printf("[addmatter option{$gapUnit $startTime $endTime} $title $desc]: 	add matter\n")
	fmt.Printf("[getstate]:                                                    	get all state\n")
	fmt.Printf("[changestate $matterTitle $stateNumber]:                       	change matter state\n")
	fmt.Printf("[changeuser $usrname]:              								change curr user\n")
	fmt.Printf("[save]:                                                       	save\n")
	fmt.Printf("[fresh]:                                                      	fresh currInfo\n")
	fmt.Printf("[exit]:                                                         	exit\n")
}

func (c *CommandManager) DisplayCommand() {
	c.dp.Display()
}

func (c *CommandManager) NewUserCommand(userid string) {
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

func (c *CommandManager) NewMatterCommand(time []string, info []string) {
	if len(info) != 2 {
		fmt.Printf("NewMatterCommand err: %v\n", fmt.Errorf("format illegal"))
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
		fmt.Printf("NewMatterCommand err: %v\n", err)
	}
}

func (c *CommandManager) GetStateCommand() {
	states := state.GetAllState()
	fmt.Printf("All State hash: \nindex:\t\tstate\n")
	for i, v := range states {
		fmt.Printf("%-*d %-*s \n", 15, i, 20, v)
	}
}

func (c *CommandManager) ChangeStateCommand(matterTitle string, stateIdx int) {
	states := state.GetAllState()
	if stateIdx >= len(states) {
		fmt.Printf("ChangeStateCommand err: %v\n", fmt.Errorf("idx out of range"))
	}
	err := dailymatter.ChangeMatterState(matterTitle, states[stateIdx])
	if err != nil {
		fmt.Printf("ChangeStateCommand err: %v\n", err)
	}
}

func (c *CommandManager) ChangeUserCommand(userid string) {
	err := dailymatter.ChangeCurrUser(userid)
	if err != nil {
		fmt.Printf("ChangeUserCommand err: %v\n", err)
	}
}

func (c *CommandManager) Save() {
	err := dailymatter.Save()
	if err != nil {
		fmt.Printf("Save err: %v\n", err)
	}
}

func (c *CommandManager) Fresh() {
	err := dailymatter.FreshCurrInfo()
	if err != nil {
		fmt.Printf("Fresh err: %v\n", err)
	}
}

func (c *CommandManager) Exit() {
	c.Save()
	fmt.Printf("Exiting...\n")
}
