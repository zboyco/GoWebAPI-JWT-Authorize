package database

// User 用户信息
type User struct {
	Model        `xorm:"extends"`
	UserName     string
	UserPwd      string
	TokenVersion int64
}
