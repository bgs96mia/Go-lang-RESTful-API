package entities

type User struct {
	Username  string    `gorm:"column:username;primaryKey"`
	Password  string    `gorm:"column:password"`
	Name      string    `gorm:"column:name"`
	Token     string    `gorm:"column:token"`
	CreatedAt int64     `gorm:"column:created_at;autoCreateTime:milli"`
	UpdatedAt int64     `gorm:"column:updated_at;autoCreateTime:milli;autoUpdateTime:milli"`
	Contacts  []Contact `gorm:"foreignKey:username;references:username"`
}

func (u *User) TableName() string {
	return "users"
}
