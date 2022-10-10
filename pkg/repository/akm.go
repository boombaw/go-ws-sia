package repository

import (
	"strconv"

	database "github.com/boombaw/go-ws-sia/pkg/database/mysql"
	"github.com/boombaw/go-ws-sia/pkg/model"
	"github.com/boombaw/go-ws-sia/pkg/query"
	"github.com/boombaw/go-ws-sia/pkg/util"
)

type akmRepository struct{}

type AkmRepository interface {
	List(arg AkmParams) ([]model.Akm, error)
	ListNA(arg AkmParams) ([]model.Akm, error)
	LastAKM(arg LastAkmParams) (model.Akm, error)
	UpdateHasSync(arg UpdateSyncParams) error
}

type AkmParams struct {
	KdProdi  string `json:"kd_prodi"`
	Semester string `json:"semester"`
}

func NewAKMRepository() *akmRepository {
	return &akmRepository{}
}

func (l *akmRepository) List(arg AkmParams) ([]model.Akm, error) {

	var db = database.Conn()
	var akm []model.Akm

	db.Debug().Raw(query.SelectAKM, arg.KdProdi, arg.Semester).Scan(&akm)

	return akm, nil
}

type AkmNAParams struct {
	Npm string `json:"npm"`
}

func (l *akmRepository) ListNA(arg AkmParams) ([]model.MhsNA, error) {

	var db = database.Conn()
	var akm []model.MhsNA

	// var actyear model.Actyear
	// db.Raw(query.GetActYear).Scan(&actyear)

	s, _ := strconv.Atoi(arg.Semester)

	studyStart := util.StudyStart(s)

	/**
	param 1 = kd_prodi
	param 2 = actyear
	param 3 = kd_prodi
	param 4 = actyear
	param 5 = study start
	param 6 = actyear
	*/

	db.Raw(query.SelectNA, arg.KdProdi, arg.Semester, arg.KdProdi, arg.Semester, studyStart, arg.Semester).Scan(&akm)

	return akm, nil
}

func (l *akmRepository) ListDO(arg AkmParams) ([]model.MhsDO, error) {

	var db = database.Conn()
	var akm []model.MhsDO

	// var actyear model.Actyear
	// db.Raw(query.GetActYear).Scan(&actyear)

	s, _ := strconv.Atoi(arg.Semester)

	studyStart := util.StudyStart(s)

	/**
	param 1 = kd_prodi
	param 2 = actyear
	param 3 = kd_prodi
	param 4 = actyear
	param 5 = study start
	param 6 = actyear
	*/

	sql, _ := db.DB()
	defer sql.Close()

	db.Raw(query.SelectNA, arg.KdProdi, arg.Semester, arg.KdProdi, arg.Semester, studyStart, arg.Semester).Scan(&akm)

	return akm, nil
}

type LastAkmParams struct {
	Npm string `json:"npm"`
}

func (l *akmRepository) LastAKM(arg LastAkmParams) (model.Akm, error) {

	var db = database.Conn()
	var akm model.Akm

	db.Raw(query.SelectLastAKM, arg.Npm).Scan(&akm)

	sql, _ := db.DB()
	defer sql.Close()

	return akm, nil
}

type UpdateSyncParams struct {
	ID int `json:"id"`
}

func (a *akmRepository) UpdateHasSync(arg UpdateSyncParams) error {

	var db = database.Conn()

	err := db.Exec(query.UpdateHasSync, arg.ID)

	if err != nil {
		return err.Error
	}

	sql, _ := db.DB()
	defer sql.Close()

	return nil
}
