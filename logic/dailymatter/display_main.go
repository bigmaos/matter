package dailymatter

import (
	"fmt"

	"github.com/liushuochen/gotable"
)

/*
type Matter struct {
	Title       string
	Desc        string
	TimeStart   time.Time
	TimeEnd     time.Time
	State       constant.State
	MatterClock *Clock
}
*/
// 暂时在控制台Display
func display() {
	if currInfo == nil {
		fmt.Printf("empty manager\n")
		return
	}
	fmt.Printf("user: %s\n", currInfo.User.Name)
	showTable()
}

func showTable() {
	header := []string{"Title", "Desc", "TimeStart", "TimeEnd", "State"}
	table, err := gotable.CreateSafeTable(header...)
	if err != nil {
		fmt.Printf("create table failed: %v\n", err)
		return
	}
	for _, matter := range currInfo.Matters {
		table.AddRow(matter.Print())
	}
	fmt.Println(table)
}
