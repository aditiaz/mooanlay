package listdto

type SubListRequest struct {
	ID        int    `json:"id"`
	Title     string `json:"title"  gorm:"type: character varying(100);not null"`
	ListId    string `json:"list_id"`
	Deskripsi string `json:"deskripsi"  gorm:"type: character varying(1000);not null"`
}
