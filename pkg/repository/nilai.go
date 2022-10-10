package repository

import (
	database "github.com/boombaw/go-ws-sia/pkg/database/mysql"
	"github.com/boombaw/go-ws-sia/pkg/model"
	"github.com/boombaw/go-ws-sia/pkg/query"
)

type nilaiRepository struct{}

type SyncNilaiParams struct {
	KdJadwal string `json:"kd_jadwal"`
}

type NilaiRepository interface {
	ListTransaksiNilai(arg SyncNilaiParams) ([]model.SyncNilai, error)
	ListJadwal(arg SyncNilaiParams) ([]model.Jadwal, error)
}

func NewNilaiRepository() *nilaiRepository {
	return &nilaiRepository{}
}

func (l *nilaiRepository) ListTransaksiNilai(arg SyncNilaiParams) ([]model.SyncNilai, error) {

	var db = database.Conn()
	var nilai []model.SyncNilai

	db.Raw(query.GetNilaiAkhir, arg.KdJadwal).Scan(&nilai)

	sql, _ := db.DB()
	defer sql.Close()

	return nilai, nil
}

func (l *nilaiRepository) ListJadwal(arg SyncNilaiParams) ([]model.Jadwal, error) {

	var db = database.Conn()
	var jadwal []model.Jadwal

	db.Raw(query.GetJadwal, arg.KdJadwal).Scan(&jadwal)

	sql, _ := db.DB()
	defer sql.Close()

	return jadwal, nil
}
