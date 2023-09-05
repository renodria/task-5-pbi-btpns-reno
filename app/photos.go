package app

type CreatePhoto struct {
	Title   string `gorm:"type:varchar(255)" json:"title" form:"title"`
	Caption string `gorm:"type:text" json:"caption" form:"caption"`
}

type UpdatePhoto struct {
	Title   string `gorm:"type:varchar(255)" json:"title" form:"title"`
	Caption string `gorm:"type:text" json:"caption" form:"caption"`
}
