package state

import (
	"daily_matter/constant"
	"slices"
)

// state 排序逻辑
var stateOrder = []string{
	constant.StateUnknown,
	constant.StateUnplanned,
	constant.StatePlanned,
	constant.StateDoing,
	constant.StateCheck,
	constant.StateDone,
}

func LessState(i, j string) bool {
	return slices.Index(stateOrder, i) < slices.Index(stateOrder, j)
}
