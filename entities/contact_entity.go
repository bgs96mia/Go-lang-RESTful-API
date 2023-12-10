package entities

type Contact struct {
	ID        string    `gorm:"column:id;primaryKey"`
	FirstName string    `gorm:"column:first_name"`
	LastName  string    `gorm:"column:last_name"`
	Email     string    `gorm:"column:email"`
	Phone     string    `gorm:"column:phone"`
	Username  string    `gorm:"column:username"`
	CreatedAt int64     `gorm:"column:created_at;autoCreateTime:milli"`
	UpdatedAt int64     `gorm:"column:updated_at;autoCreateTime:milli;autoUpdateTime:milli"`
	User      User      `gorm:"foreignKey:username;references:username"`
	Addresses []Address `gorm:"foreignKey:contact_id;references:id"`
}

func (c *Contact) TableName() string {
	return "contacts"
}
