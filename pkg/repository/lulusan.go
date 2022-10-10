package repository

import (
	database "github.com/boombaw/go-ws-sia/pkg/database/mysql"
	"github.com/boombaw/go-ws-sia/pkg/model"
	"github.com/boombaw/go-ws-sia/pkg/query"
)

type synclulusanRepository struct{}

type SyncLulusanParams struct {
	KdProdi  string `json:"kd_prodi"`
	Semester string `json:"semester"`
}

type SyncLulusanRepository interface {
	ListLulusan(arg SyncLulusanParams) ([]model.SyncLulusan, error)
}

func NewSyncLulusanRepository() *synclulusanRepository {
	return &synclulusanRepository{}
}

func (s *synclulusanRepository) ListLulusan(arg SyncLulusanParams) ([]model.SyncLulusan, error) {

	var db = database.Conn()
	var lulusan []model.SyncLulusan

	db.Raw(query.GetLulusan, arg.KdProdi, arg.Semester).Scan(&lulusan)

	sql, _ := db.DB()
	defer sql.Close()

	return lulusan, nil
}
