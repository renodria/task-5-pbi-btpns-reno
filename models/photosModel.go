package models

type Photos struct {
	ID       int16  `gorm:"primaryKey" json:"id"`
	Title    string `gorm:"type:varchar(255)" json:"title" form:"title"`
	Caption  string `gorm:"type:text" json:"caption" form:"caption"`
	PhotoUrl string `gorm:"type:varchar" json:"photo_url" form:"photo_url"`
	UserID   int    `gorm:"index" json:"user_id" form:"user_id"`
}
