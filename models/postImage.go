package models

type PostImage struct {
	ID     string `json:"id"`
	ListID int    `json:"-"`
	Image  string `json:"image" gorm:"type: varchar(255);`
}
