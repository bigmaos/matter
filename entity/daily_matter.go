package entity

type CurrMattersInfo struct {
	Matters []*Matter
	User    *User
}

func NewCurrMattersInfo() *CurrMattersInfo {
	return &CurrMattersInfo{}
}

type InsertedMatterInfo struct {
	TimeGap          string
	StartTimeFromNow int64
	EndTimeFromNow   int64
	Title            string
	Desc             string
	State            string
}
