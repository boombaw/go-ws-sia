package repository

import (
	database "github.com/boombaw/go-ws-sia/pkg/database/mysql"
	"github.com/boombaw/go-ws-sia/pkg/query"
)

type smsProdiRepository struct{}

type SmsProdiRepository interface {
	SmsProdi(arg SmsParams) string
}

type SmsParams struct {
	KdProdi string `json:"kd_prodi"`
}

func NewSmsProdiRepository() *smsProdiRepository {
	return &smsProdiRepository{}
}

func (l *smsProdiRepository) SMSProdi(arg SmsParams) string {

	var db = database.Conn()
	var data []uint8

	row := db.Raw(query.SmsProdi, arg.KdProdi).Row()
	row.Scan(&data)

	sql, _ := db.DB()
	defer sql.Close()

	return string(data)
}
