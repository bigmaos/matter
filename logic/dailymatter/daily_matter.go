package dailymatter

import (
	"daily_matter/entity"
	mm "daily_matter/logic/mattermanager"
	"fmt"
	"sort"
)

var currInfo *entity.CurrMattersInfo

func init() {
	manager := mm.Manager
	// 初始化时，只拿其中一个用户的matter
	for _, matters := range manager.GetMatters() {
		currInfo.User = matters.User
		for _, matter := range matters.Matters {
			currInfo.Matters = append(currInfo.Matters, matter)
		}
		break
	}
	sort.Slice(currInfo.Matters, func(i, j int) bool {
		if currInfo.Matters[i].State < currInfo.Matters[j].State {
			return true
		}
		return false
	})
}

// 暂时在控制台Display
func display() {
	if currInfo == nil {
		fmt.Printf("empty manager\n")
		return
	}
	fmt.Printf("user: %s\n", currInfo.User.Name)
	for _, matter := range currInfo.Matters {
		fmt.Printf("")
	}
}
