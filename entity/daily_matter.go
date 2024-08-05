package entity

type CurrMattersInfo struct {
	Matters []*Matter
	User    *User
}

type InsertedMatterInfo struct {
	TimeGap          string
	StartTimeFromNow int64
	EndTimeFromNow   int64
	Title            string
	Desc             string
	State            string
}
