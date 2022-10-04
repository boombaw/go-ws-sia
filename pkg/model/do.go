package model

type DO struct {
	ID          string `json:"id" gorm:"id"`
	Skep        string `json:"skep" gorm:"skep"`
	TglSkep     string `json:"tgl_skep" gorm:"tgl_skep"`
	Npm         string `json:"npm" gorm:"npm"`
	Alasan      string `json:"alasan" gorm:"alasan"`
	Tahunajaran string `json:"tahunajaran" gorm:"tahunajaran"`
}
