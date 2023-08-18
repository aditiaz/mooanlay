package models

type SubList struct {
	ID           int            `json:"id"`
	PostImageSub []PostImageSub `json:"post_image_sub" gorm:"foreignKey:ListID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	ListId       int            `json:"list_id" `
	SubListId    int            `json:"-" `
	Title        string         `json:"title"  validate:"required" gorm:"type: character varying(100);not null"`
	Deskripsi    string         `json:"deskripsi"  validate:"required" gorm:"type: character varying(1000);not null"`
}

func (SubList) TableName() string {
	return "sub_lists"
}
