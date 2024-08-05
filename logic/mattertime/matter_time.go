package mattertime

import (
	"daily_matter/constant"
	"time"
)

var GapUnitMap map[string]time.Duration

// 默认30天一个月
func init() {
	GapUnitMap[constant.TimeGapUnitHour] = time.Hour
	GapUnitMap[constant.TimeGapUnitDay] = time.Hour * 24
	GapUnitMap[constant.TimeGapUnitWeek] = time.Hour * 24 * 7
	GapUnitMap[constant.TimeGapUnitMonth] = time.Hour * 24 * 30
}
