package repository

import (
	database "github.com/boombaw/go-ws-sia/pkg/database/mysql"
	"github.com/boombaw/go-ws-sia/pkg/model"
	"github.com/boombaw/go-ws-sia/pkg/query"
)

type doRepository struct{}

type DORepository interface {
	List(arg DOParams) ([]model.Akm, error)
}

type DOParams struct {
	KdProdi  string `json:"kd_prodi"`
	Semester string `json:"semester"`
}

func NewDORepository() *doRepository {
	return &doRepository{}
}

func (l *doRepository) List(arg DOParams) ([]model.DO, error) {

	var db = database.Conn()
	var do []model.DO

	db.Debug().Raw(query.SelectListDO, arg.KdProdi, arg.Semester).Scan(&do)

	return do, nil
}

// type UpdateSyncParams struct {
// 	ID int `json:"id"`
// }

// func (a *akmRepository) UpdateHasSync(arg UpdateSyncParams) error {

// 	var db = database.Conn()

// 	err := db.Exec(query.UpdateHasSync, arg.ID)

// 	if err != nil {
// 		return err.Error
// 	}

// 	return nil
// }
