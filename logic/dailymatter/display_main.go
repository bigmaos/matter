package dailymatter

import (
	"daily_matter/entity"
	"fmt"

	"github.com/liushuochen/gotable"
)

// 暂时在控制台Display
func Display() {
	if currInfo == nil {
		fmt.Printf("empty manager\n")
		return
	}
	fmt.Printf("user: %s\n", currInfo.User.Name)
	ShowTable()
}

func ShowTable() {
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
