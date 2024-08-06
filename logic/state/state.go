package state

import (
	"daily_matter/constant"
	"slices"
)

// state 排序逻辑
var allState = []string{
	constant.StateUnknown,
	constant.StateUnplanned,
	constant.StatePlanned,
	constant.StateDoing,
	constant.StateCheck,
	constant.StateDone,
}

func LessState(i, j string) bool {
	return slices.Index(allState, i) < slices.Index(allState, j)
}

func GetAllState() []string {
	return allState
}
