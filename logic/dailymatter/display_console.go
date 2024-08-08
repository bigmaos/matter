package dailymatter

import (
	"daily_matter/entity"
	"fmt"

	"github.com/liushuochen/gotable"
)

type DisplayConsolePacker struct {
}

// 暂时在控制台Display
func (d DisplayConsolePacker) Display() {
	if currInfo == nil || currInfo.User == nil {
		fmt.Printf("empty manager\n")
		return
	}
	fmt.Printf("user: %s\n", currInfo.User.GetName())
	d.ShowTable()
}

func (d DisplayConsolePacker) ShowTable() {
	matter := &entity.Matter{}
	table, err := gotable.CreateSafeTable(matter.GetLabel()...)
	if err != nil {
		fmt.Printf("create table failed: %v\n", err)
		return
	}
	for _, matter := range currInfo.Matters {
		table.AddRow(matter.Print())
	}
	fmt.Println(table)
}
