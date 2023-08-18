package listdto

type ListResponse struct {
	ID    int    `json:"id"`
	Title string `json:"title"  gorm:"type: character varying(100);not null"`
}
