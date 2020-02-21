package dao

// UserProfile 用户
type UserProfile struct {
	ID           int64  `column:"id"`
	AccName      string `column:"acc_name"`
	ChineseName  string `column:"chinese_name"`
	ScholarField string `column:"scholar_field"`
	Introduction string `column:"introduction"`
}

// func QueryUserProfile(tableName string) []
