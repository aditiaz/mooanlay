package models

type List struct {
	ID        int         `json:"id"`
	PostImage []PostImage `json:"post_image" gorm:"foreignKey:ListID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	SubList   []SubList   `json:"sub_list" gorm:"foreignKey:SubListId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Title     string      `json:"title"  gorm:"type: character varying(100);not null"`
	Deskripsi string      `json:"deskripsi"  gorm:"type: character varying(1000);not null"`
}

func (List) TableName() string {
	return "lists"
}
